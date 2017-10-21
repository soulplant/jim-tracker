package main

import (
	"reflect"
	"testing"

	context "golang.org/x/net/context"

	"github.com/soulplant/talk-tracker/api"
)

// MakeNewTestService creates and inits a datastore-backed api service.
func MakeNewTestService(t *testing.T) *dsApiService {
	ds := InitTestDatastore(t)
	a := NewDsApiService(ds)
	err := a.Init()
	if err != nil {
		t.Fatalf("Failed to init %v", err)
	}
	return a
}

func TestUpdateUser(t *testing.T) {
	a := MakeNewTestService(t)
	defer a.Close()

	id := addUser(t, a, "steve")
	_, err := a.UpdateUser(context.Background(), &api.UpdateUserRequest{
		UserId:      id,
		HasName:     true,
		Name:        "steve2",
		HasNextTalk: true,
		NextTalk:    "next-talk",
	})
	if err != nil {
		t.Errorf("failed to update user: %v", err)
	}

	resp, err := a.GetUsers(context.Background(), &api.GetUsersRequest{})
	if err != nil {
		t.Errorf("failed to retrieve users: %v", err)
	}

	if "steve2" != resp.User[0].GetName() {
		t.Errorf("Failed to update the name, expected 'steve2', but got '%s'", resp.User[0].GetName())
	}
}

func TestRemoveUser(t *testing.T) {
	a := MakeNewTestService(t)
	defer a.Close()
	userId := addUser(t, a, "steve")

	_, err := a.RemoveUser(context.Background(), &api.RemoveUserRequest{
		UserId: userId,
	})
	if err != nil {
		t.Errorf("Failed to remove user %d: %s", userId, err)
	}
	resp, err := a.GetUsers(context.Background(), &api.GetUsersRequest{})
	if err != nil {
		t.Errorf("Failed to retrieve users")
	}
	if len(resp.User) != 0 {
		t.Errorf("There are more than zero users (%d user(s)), but we expected the only user to be deleted", len(resp.User))
	}
}

func TestCompleteTalk(t *testing.T) {
	a := MakeNewTestService(t)
	defer a.Close()
	steveId := addUser(t, a, "steve")
	addUser(t, a, "joe")

	_, err := a.UpdateUser(context.Background(), &api.UpdateUserRequest{
		UserId:      steveId,
		HasNextTalk: true,
		NextTalk:    "next talk",
	})
	if err != nil {
		t.Fatalf("Failed to update the user %d: %v", steveId, err)
	}

	_, err = a.CompleteTalk(context.Background(), &api.CompleteTalkRequest{
		UserId: steveId,
	})
	if err != nil {
		t.Fatalf("Failed to complete the talk: %s", err)
	}

	resp, err := a.GetUsers(context.Background(), &api.GetUsersRequest{})
	if err != nil {
		t.Fatalf("Failed to GetUsers(): %s", err)
	}
	if len(resp.GetUser()) != 2 {
		t.Errorf("Expected 2 users, but actually %d", len(resp.GetUser()))
	}
	next := resp.GetUser()[0].Name
	if "joe" != next {
		t.Errorf("Expected joe to be next in line for a talk, but it was %s", next)
	}
	if resp.GetUser()[1].GetNextTalk() != "" {
		t.Errorf("Expected steve's next talk field to be cleared")
	}
}

func TestReorder(t *testing.T) {
	for _, tc := range []struct {
		move, anchor int
		before       bool
		expected     []int64
	}{
		{0, 1, true, []int64{0, 1, 2}},
		{0, 1, false, []int64{1, 0, 2}},
		{0, 2, true, []int64{1, 0, 2}},
		{2, 0, true, []int64{2, 0, 1}},
		{2, 0, false, []int64{0, 2, 1}},
		{2, 2, true, []int64{0, 1, 2}},
		{2, 1, true, []int64{0, 2, 1}},
	} {
		a := MakeNewTestService(t)
		defer a.Close()
		names := []string{"james", "anu", "phuz"}
		var userIds []int64
		for _, name := range names {
			userIds = append(userIds, addUser(t, a, name))
		}

		a.Reorder(context.Background(), &api.ReorderRequest{
			AnchorUserId: userIds[tc.anchor],
			MoveUserId:   userIds[tc.move],
			Before:       tc.before,
		})

		res, err := a.GetUsers(context.Background(), &api.GetUsersRequest{})
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		resUserIds := make([]int64, len(res.GetUser()))
		for i, v := range res.GetUser() {
			resUserIds[i] = v.GetId()
		}
		expUserIds := make([]int64, len(res.GetUser()))
		for i, v := range tc.expected {
			expUserIds[i] = userIds[v]
		}
		if !reflect.DeepEqual(resUserIds, expUserIds) {
			t.Errorf("after moving %d to %d (before: %v), users: %v, but expected %v", tc.move, tc.anchor, tc.before, resUserIds, expUserIds)
		}
	}
}

func TestReorder_InvalidUsers(t *testing.T) {
	a := MakeNewTestService(t)
	defer a.Close()
	_, err := a.Reorder(context.Background(), &api.ReorderRequest{
		AnchorUserId: 100,
		MoveUserId:   200,
		Before:       true,
	})
	if err == nil {
		t.Fail()
	}
}

func TestAddUser(t *testing.T) {
	s := MakeNewTestService(t)
	defer s.Close()
	_, err := s.AddUser(context.Background(), &api.AddUserRequest{
		Name: "james",
	})
	if err != nil {
		t.Fatal(err)
	}

	resp, err := s.FetchAll(context.Background(), &api.FetchAllRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.User) == 0 {
		t.Fatal("Expected at least one user")
	}
	if resp.User[0].Name != "james" {
		t.Fatal("Expected retrieved user to be james")
	}
}

func TestInit(t *testing.T) {
	s := MakeNewTestService(t)
	defer s.Close()

	steveId := addUser(t, s, "steve")
	if err := s.Init(); err != nil {
		t.Fatalf("Failed to do idempotent Init: %v", err)
	}
	resp, err := s.GetUsers(context.Background(), &api.GetUsersRequest{})
	if err != nil {
		t.Fatalf("Failed to get users")
	}
	if len(resp.GetUser()) == 0 {
		t.Fatalf("Expected users to be preserved, but there are none")
	}
	if resp.GetUser()[0].Id != steveId {
		t.Fatalf("Expected user to be persisted, but id is changed (from %d to %d)", steveId, resp.GetUser()[0].Id)
	}
}

func TestRootData(t *testing.T) {
	input := []int64{1, 2, 3}
	expected := []int64{2, 3, 1}
	data := rootData{Order: input}
	data.MoveUserToEnd(1)
	if !reflect.DeepEqual(data.Order, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

type UserAdder interface {
	AddUser(context.Context, *api.AddUserRequest) (*api.AddUserResponse, error)
}

func addUser(t *testing.T, a UserAdder, name string) int64 {
	u, err := a.AddUser(context.Background(), &api.AddUserRequest{Name: name})
	if err != nil {
		t.Fatalf("failed to add User '%s': %v", name, err)
	}
	return u.GetUserId()
}
