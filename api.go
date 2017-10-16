package main

import (
	"errors"
	"fmt"

	"cloud.google.com/go/datastore"

	"github.com/soulplant/talk-tracker/api"
	context "golang.org/x/net/context"
)

// An apiService backed by datastore
type dsApiService struct {
	key    *datastore.Key
	client *datastore.Client
}

// TODO(james): Replace this with an interceptor that deserialises/serialises before/after the RPC.
func NewDsApiService(client *datastore.Client, key *datastore.Key) *dsApiService {
	return &dsApiService{key, client}
}

func (s *dsApiService) getData() (*apiService, error) {
	var data apiService
	err := s.client.Get(context.Background(), s.key, &data)
	return &data, err
}

func (s *dsApiService) FetchAll(ctx context.Context, req *api.FetchAllRequest) (*api.FetchAllResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	return api.FetchAll(ctx, req)
}

func (s *dsApiService) GetUsers(ctx context.Context, req *api.GetUsersRequest) (*api.GetUsersResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	return api.GetUsers(ctx, req)
}

func (s *dsApiService) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	resp, err := api.AddUser(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = s.client.Put(ctx, s.key, api)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *dsApiService) AddTalk(ctx context.Context, req *api.AddTalkRequest) (*api.AddTalkResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	resp, err := api.AddTalk(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = s.client.Put(ctx, s.key, api)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *dsApiService) Reorder(ctx context.Context, req *api.ReorderRequest) (*api.ReorderResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	resp, err := api.Reorder(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = s.client.Put(ctx, s.key, api)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *dsApiService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	resp, err := api.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = s.client.Put(ctx, s.key, api)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *dsApiService) RemoveUser(ctx context.Context, req *api.RemoveUserRequest) (*api.RemoveUserResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	resp, err := api.RemoveUser(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = s.client.Put(ctx, s.key, api)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *dsApiService) CompleteTalk(ctx context.Context, req *api.CompleteTalkRequest) (*api.CompleteTalkResponse, error) {
	api, err := s.getData()
	if err != nil {
		return nil, err
	}
	resp, err := api.CompleteTalk(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = s.client.Put(ctx, s.key, api)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type apiService api.ServerState

func NewApiService() *apiService {
	return &apiService{}
}

func (s *apiService) PopulateWithTestData() {
	s.AddUser(context.Background(), &api.AddUserRequest{Name: "James"})
	s.AddUser(context.Background(), &api.AddUserRequest{Name: "Anu"})
}

func (s *apiService) FetchAll(context.Context, *api.FetchAllRequest) (*api.FetchAllResponse, error) {
	return &api.FetchAllResponse{
		Talk: s.Talk,
		User: s.User,
	}, nil
}

func (s *apiService) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error) {
	user := &api.User{
		Id:   fmt.Sprintf("%d", s.NextId),
		Name: req.GetName(),
	}
	s.NextId += 1

	s.User = append(s.User, user)

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
	s.Talk = append(s.Talk, talk)
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
		tmp = append(tmp, s.User[:mi]...)
		tmp = append(tmp, s.User[mi+1:ins]...)
		tmp = append(tmp, s.User[mi])
		tmp = append(tmp, s.User[ins:]...)
		s.User = tmp
	} else {
		var tmp []*api.User
		tmp = append(tmp, s.User[:ins]...)
		tmp = append(tmp, s.User[mi])
		tmp = append(tmp, s.User[ins:mi]...)
		tmp = append(tmp, s.User[mi+1:]...)
		s.User = tmp
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
		s.User[i].Name = req.GetName()
	}
	if req.GetNextTalk() != "" {
		s.User[i].NextTalk = req.GetNextTalk()
	}
	return &api.UpdateUserResponse{}, nil
}

func (s *apiService) RemoveUser(ctx context.Context, req *api.RemoveUserRequest) (*api.RemoveUserResponse, error) {
	i, err := s.indexOfUser(req.GetUserId())
	if err != nil {
		return nil, err
	}
	s.User = append(s.User[:i], s.User[i+1:]...)
	return &api.RemoveUserResponse{}, nil
}

func (s *apiService) CompleteTalk(ctx context.Context, req *api.CompleteTalkRequest) (*api.CompleteTalkResponse, error) {
	i, err := s.indexOfUser(req.GetUserId())
	if err != nil {
		return nil, err
	}
	user := s.User[i]
	user.NextTalk = ""
	// TODO(james): Record the fact that the Talk happened.
	s.User = append(s.User[:i], s.User[i+1:]...)
	s.User = append(s.User, user)
	return &api.CompleteTalkResponse{}, nil
}

func (s *apiService) indexOfUser(userId string) (int, error) {
	for i, u := range s.User {
		if u.GetId() == userId {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Failed to get index of User %s", userId)
}

func (s *apiService) GetUsers(ctx context.Context, req *api.GetUsersRequest) (*api.GetUsersResponse, error) {
	return &api.GetUsersResponse{
		User: s.User,
	}, nil
}

func (s *apiService) allocateId() string {
	result := s.NextId
	s.NextId += 1
	return fmt.Sprintf("%d", result)
}
