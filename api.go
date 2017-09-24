package main

import (
	"github.com/soulplant/talk-tracker/api"
	context "golang.org/x/net/context"
)

type apiService struct {
	user  []*api.User
	talks []*api.Talk
}

func (s *apiService) FetchAll(context.Context, *api.FetchAllRequest) (*api.FetchAllResponse, error) {
	// TODO(james): Make real.
	return &api.FetchAllResponse{
		Talk: []*api.Talk{&api.Talk{
			Id:        "1",
			Done:      true,
			Name:      "KK",
			SpeakerId: "1",
		}},
		User: []*api.User{&api.User{
			Id:   "1",
			Name: "James",
		}},
	}, nil
}
