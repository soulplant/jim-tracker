package main

import (
	"context"
	"testing"

	"github.com/soulplant/talk-tracker/api"
)

func Test_reorder(t *testing.T) {
	for _, tc := range []struct {
		move, anchor int
		before       bool
		expected     []int
	}{
		{0, 1, true, []int{0, 1, 2}},
		{0, 1, false, []int{1, 0, 2}},
		{0, 2, true, []int{1, 0, 2}},
		{2, 0, true, []int{2, 0, 1}},
		{2, 0, false, []int{0, 2, 1}},
		{2, 2, true, []int{0, 1, 2}},
		{2, 1, true, []int{0, 2, 1}},
	} {
		a := NewApiService()
		names := []string{"james", "anu", "phuz"}
		var userIds []string
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
		for i, u := range res.GetUser() {
			if u.GetId() != userIds[tc.expected[i]] {
				t.Errorf("after moving %d to %d (before: %v), users: %v, but expected %v", tc.move, tc.anchor, tc.before, res.GetUser(), tc.expected)
			}
		}
	}
}
func Test_reorder__invalid_users(t *testing.T) {
	a := NewApiService()
	_, err := a.Reorder(context.Background(), &api.ReorderRequest{
		AnchorUserId: "adsf",
		MoveUserId:   "def",
		Before:       true,
	})
	if err == nil {
		t.Fail()
	}
}

func addUser(t *testing.T, a *apiService, name string) string {
	u, err := a.AddUser(context.Background(), &api.AddUserRequest{Name: name})
	if err != nil {
		t.Errorf("failed to add User '%s'", name)
	}
	return u.GetUser().GetId()
}
