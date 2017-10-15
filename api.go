package main

import (
	"errors"
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
	return &apiService{}
}

func NewTestApiService() *apiService {
	a := NewApiService()
	a.PopulateWithTestData()
	return a
}

func (s *apiService) PopulateWithTestData() {
	s.AddUser(context.Background(), &api.AddUserRequest{Name: "James"})
	s.AddUser(context.Background(), &api.AddUserRequest{Name: "Anu"})
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
	if len(req.GetUserId()) == 0 {
		return nil, errors.New("user_id required")
	}
	talk := &api.Talk{
		Id:        s.allocateId(),
		SpeakerId: req.GetUserId(),
	}
	s.talk = append(s.talk, talk)
	return &api.AddTalkResponse{
		Talk: talk,
	}, nil
}

func (s *apiService) Reorder(ctx context.Context, req *api.ReorderRequest) (*api.ReorderResponse, error) {
	ai, err := s.indexOfUser(req.AnchorUserId)
	if err != nil {
		return nil, err
	}
	mi, err := s.indexOfUser(req.MoveUserId)
	if err != nil {
		return nil, err
	}
	ins := ai
	if !req.GetBefore() {
		ins += 1
	}
	if mi < ai {
		var tmp []*api.User
		tmp = append(tmp, s.user[:mi]...)
		tmp = append(tmp, s.user[mi+1:ins]...)
		tmp = append(tmp, s.user[mi])
		tmp = append(tmp, s.user[ins:]...)
		s.user = tmp
	} else {
		var tmp []*api.User
		tmp = append(tmp, s.user[:ins]...)
		tmp = append(tmp, s.user[mi])
		tmp = append(tmp, s.user[ins:mi]...)
		tmp = append(tmp, s.user[mi+1:]...)
		s.user = tmp
	}

	return &api.ReorderResponse{
		Accepted: true,
	}, nil
}

func (s *apiService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	i, err := s.indexOfUser(req.GetUserId())
	if err != nil {
		return nil, err
	}
	if req.GetName() != "" {
		s.user[i].Name = req.GetName()
	}
	if req.GetNextTalkName() != "" {
		s.user[i].NextTalk = req.GetNextTalkName()
	}
	return &api.UpdateUserResponse{}, nil
}

func (s *apiService) RemoveUser(ctx context.Context, req *api.RemoveUserRequest) (*api.RemoveUserResponse, error) {
	i, err := s.indexOfUser(req.GetUserId())
	if err != nil {
		return nil, err
	}
	s.user = append(s.user[:i], s.user[i+1:]...)
	return &api.RemoveUserResponse{}, nil
}

func (s *apiService) CompleteTalk(ctx context.Context, req *api.CompleteTalkRequest) (*api.CompleteTalkResponse, error) {
	i, err := s.indexOfUser(req.GetUserId())
	if err != nil {
		return nil, err
	}
	user := s.user[i]
	user.NextTalk = ""
	// TODO(james): Record the fact that the talk happened.
	s.user = append(s.user[:i], s.user[i+1:]...)
	s.user = append(s.user, user)
	return &api.CompleteTalkResponse{}, nil
}

func (s *apiService) indexOfUser(userId string) (int, error) {
	for i, u := range s.user {
		if u.GetId() == userId {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Failed to get index of user %s", userId)
}

func (s *apiService) GetUsers(ctx context.Context, req *api.GetUsersRequest) (*api.GetUsersResponse, error) {
	return &api.GetUsersResponse{
		User: s.user,
	}, nil
}

func (s *apiService) allocateId() string {
	result := s.nextId
	s.nextId += 1
	return fmt.Sprintf("%d", result)
}
