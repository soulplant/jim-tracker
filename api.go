package main

import (
	"fmt"

	"github.com/soulplant/talk-tracker/api"
	context "golang.org/x/net/context"
)

type apiService struct {
	user []*api.User
	talk []*api.Talk

	nextId int
}

func NewApiService() *apiService {
	return &apiService{
		talk: []*api.Talk{&api.Talk{
			Id:        "1",
			Done:      true,
			Name:      "Kazam! It's Kubernetes!",
			SpeakerId: "0",
		}},
		user: []*api.User{&api.User{
			Id:   "0",
			Name: "James",
		}},
		nextId: 2,
	}
}

func (s *apiService) FetchAll(context.Context, *api.FetchAllRequest) (*api.FetchAllResponse, error) {
	// TODO(james): Make real.
	return &api.FetchAllResponse{
		Talk: s.talk,
		User: s.user,
	}, nil
}

func (s *apiService) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error) {
	user := &api.User{
		Id:   fmt.Sprintf("%d", s.nextId),
		Name: req.GetName(),
	}
	s.nextId += 1

	s.user = append(s.user, user)

	return &api.AddUserResponse{
		User: user,
	}, nil
}

func (s *apiService) AddTalk(ctx context.Context, req *api.AddTalkRequest) (*api.AddTalkResponse, error) {
	talk := &api.Talk{
		Id:        s.allocateId(),
		SpeakerId: req.GetUserId(),
	}
	s.talk = append(s.talk, talk)
	return &api.AddTalkResponse{
		Talk: talk,
	}, nil
}

func (s *apiService) allocateId() string {
	result := s.nextId
	s.nextId += 1
	return fmt.Sprintf("%d", result)
}
