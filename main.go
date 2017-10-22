package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"

	"context"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/soulplant/talk-tracker/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":1234"
const grpcPort = "127.0.0.1:1235"

var projectId = flag.String("projectId", "dev", "The GCP project to connect to")
var basicAuthUser = os.Getenv("BASIC_AUTH_USER")
var basicAuthPass = os.Getenv("BASIC_AUTH_PASS")

//go:generate ./gen-protos.sh

func filesHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	http.StripPrefix("/files/", http.FileServer(http.Dir("."))).ServeHTTP(w, r)
}

// authCheck returns true if the request has a valid authentication header.
func authCheck(r *http.Request) bool {
	// If the environment variables aren't set, then don't do the auth check.
	if basicAuthUser == "" {
		return true
	}
	user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return user == basicAuthUser && pass == basicAuthPass
}

func RequireAuth(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Helix Talks"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func main() {
	apiMux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	api.RegisterApiServiceHandlerFromEndpoint(ctx, apiMux, grpcPort, []grpc.DialOption{grpc.WithInsecure()})
	http.Handle("/api/", RequireAuth(http.StripPrefix("/api", apiMux)))
	http.Handle("/files/", RequireAuth(http.HandlerFunc(filesHandler)))

	go func() {
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatal("Failed to listen", err)
		}
		rpcServer := grpc.NewServer()
		reflection.Register(rpcServer)

		client, err := datastore.NewClient(ctx, *projectId)
		if err != nil {
			log.Fatalf("Couldn't create datastore client %v", err)
		}
		service := NewDsApiService(client)
		if err = service.Init(); err != nil {
			log.Fatalf("Couldn't init datastore %s", err)
		}
		service.FetchAll(context.Background(), &api.FetchAllRequest{})
		api.RegisterApiServiceServer(rpcServer, service)
		fmt.Printf("Listening for gRPC on %s\n", grpcPort)
		log.Fatal(rpcServer.Serve(lis))
	}()

	fmt.Printf("Listening for HTTP on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Failed to listen", err)
	}
}

// func OpenTestDb() *gorm.DB {
// 	os.Remove("test.db")
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	// Migrate the schema
// 	db.AutoMigrate(&Project{}, &User{}, &Task{}, &Stretch{}, &Category{})
// 	return db
// }

// func test(db *gorm.DB) {
// 	// Read
// 	var project Project
// 	if e := db.First(&project, 1000); e.Error != nil {
// 		fmt.Println("Couldn't find 1000")
// 	}
// 	if e := db.First(&project, 1); e.Error != nil {
// 		fmt.Println("Couldn't find 1")
// 	}
// 	db.First(&project, "Name = ?", "Dreamer")

// 	// Delete - delete project
// 	// db.Delete(&project)
// }
