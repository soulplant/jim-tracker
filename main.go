package main

import (
	"fmt"
	"log"
	"net/http"

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

//go:generate ./gen-protos.sh

func main() {
	apiMux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	api.RegisterApiServiceHandlerFromEndpoint(ctx, apiMux, grpcPort, []grpc.DialOption{grpc.WithInsecure()})
	http.Handle("/api/", http.StripPrefix("/api", apiMux))
	http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		http.StripPrefix("/files/", http.FileServer(http.Dir("."))).ServeHTTP(w, r)
	})
	go func() {
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatal("Failed to listen", err)
		}
		rpcServer := grpc.NewServer()
		reflection.Register(rpcServer)
		api.RegisterApiServiceServer(rpcServer, NewApiService())
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
