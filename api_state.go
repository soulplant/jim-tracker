package main

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/golang/protobuf/ptypes"
	"github.com/soulplant/talk-tracker/api"
	context "golang.org/x/net/context"
)

// An apiService backed by datastore
type dsApiService struct {
	key    *datastore.Key
	client *datastore.Client
}

// NewDsApiService constructs an api service that is backed by the given.
func NewDsApiService(client *datastore.Client) *dsApiService {
	key := datastore.NameKey("order", "order", nil)
	return &dsApiService{key, client}
}

type rootData struct {
	Order []int64
}

func (s *dsApiService) Close() {
	s.client.Close()
}

// Init initialises the datastore with an empty root object.
func (s *dsApiService) Init() error {
	_, err := s.client.RunInTransaction(context.Background(), func(tx *datastore.Transaction) error {
		var root rootData
		err := tx.Get(s.key, &root)
		if err == nil {
			return nil
		}

		if err == datastore.ErrNoSuchEntity {
			root = rootData{Order: []int64{}}
			_, err := s.client.Put(context.Background(), s.key, &root)
			return err
		}
		return err
	})
	return err
}

func (s *dsApiService) FetchAll(ctx context.Context, req *api.FetchAllRequest) (*api.FetchAllResponse, error) {
	users, err := s.GetUsers(ctx, &api.GetUsersRequest{})
	if err != nil {
		return nil, err
	}
	return &api.FetchAllResponse{
		User: users.User,
	}, nil
}

func (s *dsApiService) GetUsers(ctx context.Context, req *api.GetUsersRequest) (*api.GetUsersResponse, error) {
	var resp *api.GetUsersResponse
	_, err := s.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var data rootData
		if err := tx.Get(s.key, &data); err != nil {
			return fmt.Errorf("Failed to get data %s", err)
		}
		keys := []*datastore.Key{}
		for _, id := range data.Order {
			keys = append(keys, userKey(id))
		}
		users := make([]*api.User, len(keys))
		if err := tx.GetMulti(keys, users); err != nil {
			return fmt.Errorf("GetMulti failed: %s", err)
		}
		for i, u := range users {
			u.Id = keys[i].ID
		}
		resp = &api.GetUsersResponse{
			User: users,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *dsApiService) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error) {
	key := userKey(0)
	user := api.User{
		Name: req.GetName(),
	}
	key, err := s.client.Put(ctx, key, &user)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a new user: %v", err)
	}
	_, err = s.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var data rootData
		err := tx.Get(s.key, &data)
		if err != nil {
			return err
		}
		data.Order = append(data.Order, key.ID)
		_, err = tx.Put(s.key, &data)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &api.AddUserResponse{
		UserId: key.ID,
	}, nil
}

func (s *dsApiService) AddTalk(ctx context.Context, req *api.AddTalkRequest) (*api.AddTalkResponse, error) {
	key, err := s.client.Put(ctx, talkKey(""), &api.Talk{
		SpeakerId: req.UserId,
		Name:      req.Name,
		Completed: ptypes.TimestampNow(),
	})
	if err != nil {
		return nil, err
	}
	return &api.AddTalkResponse{
		TalkId: key.ID,
	}, nil
}

// indexOf returns the index of needle in haystack, or -1 if it isn't in haystack.
func indexOf(haystack []int64, needle int64) int {
	for i, h := range haystack {
		if h == needle {
			return i
		}
	}
	return -1
}

// removeItem removes r from ns if present, and returns an error otherwise
func removeItem(ns []int64, r int64) ([]int64, error) {
	i := indexOf(ns, r)
	if i == -1 {
		return nil, fmt.Errorf("Item not found in list %d %v", r, ns)
	}
	tmp := ns[:i]
	tmp = append(tmp, ns[i+1:]...)
	return tmp, nil
}

func (s *dsApiService) Reorder(ctx context.Context, req *api.ReorderRequest) (*api.ReorderResponse, error) {
	_, err := s.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var data rootData
		err := tx.Get(s.key, &data)
		if err != nil {
			return err
		}
		order := data.Order
		ai := indexOf(order, req.AnchorUserId)
		if ai == -1 {
			return fmt.Errorf("Couldn't find anchor item in order")
		}
		mi := indexOf(order, req.MoveUserId)
		if mi == -1 {
			return fmt.Errorf("Couldn't find move item in order")
		}

		ins := ai
		if !req.GetBefore() {
			ins += 1
		}
		if mi < ai {
			var tmp []int64
			tmp = append(tmp, order[:mi]...)
			tmp = append(tmp, order[mi+1:ins]...)
			tmp = append(tmp, order[mi])
			tmp = append(tmp, order[ins:]...)
			order = tmp
		} else {
			var tmp []int64
			tmp = append(tmp, order[:ins]...)
			tmp = append(tmp, order[mi])
			tmp = append(tmp, order[ins:mi]...)
			tmp = append(tmp, order[mi+1:]...)
			order = tmp
		}
		data.Order = order
		_, err = tx.Put(s.key, &data)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &api.ReorderResponse{
		Accepted: true,
	}, nil
}

func (s *dsApiService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	_, err := s.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		return s.updateUser(ctx, tx, req)
	})
	if err != nil {
		return nil, err
	}
	return &api.UpdateUserResponse{}, nil
}

// updateUser is the same as UpdateUser, but it doesn't run in a transaction so
// its operations can be run as part of a larger transaction.
func (s *dsApiService) updateUser(ctx context.Context, tx *datastore.Transaction, req *api.UpdateUserRequest) error {
	var user api.User
	key := userKey(req.GetUserId())
	err := tx.Get(key, &user)
	if err != nil {
		return err
	}
	if req.GetHasName() {
		user.Name = req.GetName()
	}
	if req.GetHasNextTalk() {
		user.NextTalk = req.GetNextTalk()
	}
	_, err = tx.Put(key, &user)
	return err
}

func (d *rootData) MoveUserToEnd(userId int64) error {
	i := indexOf(d.Order, userId)
	if i == -1 {
		return fmt.Errorf("User %d is not present in the rotation", userId)
	}
	tmp := d.Order[:i]
	tmp = append(tmp, d.Order[i+1:]...)
	tmp = append(tmp, userId)
	d.Order = tmp
	return nil
}

func (d *rootData) RemoveUser(userId int64) error {
	tmp, err := removeItem(d.Order, userId)
	if err != nil {
		return err
	}
	d.Order = tmp
	return nil
}

// RemoveUser removes the specified user from the talk rotation. Note that it
// doesn't delete the user object itself.
func (s *dsApiService) RemoveUser(ctx context.Context, req *api.RemoveUserRequest) (*api.RemoveUserResponse, error) {
	_, err := s.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var data rootData
		if err := tx.Get(s.key, &data); err != nil {
			return err
		}
		if err := data.RemoveUser(req.UserId); err != nil {
			return err
		}
		_, err := tx.Put(s.key, &data)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &api.RemoveUserResponse{}, nil
}

func (s *dsApiService) CompleteTalk(ctx context.Context, req *api.CompleteTalkRequest) (*api.CompleteTalkResponse, error) {
	_, err := s.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var data rootData
		if err := tx.Get(s.key, &data); err != nil {
			return err
		}
		if err := data.MoveUserToEnd(req.UserId); err != nil {
			return err
		}
		if _, err := tx.Put(s.key, &data); err != nil {
			return err
		}
		return s.updateUser(ctx, tx, &api.UpdateUserRequest{
			UserId:      req.GetUserId(),
			HasNextTalk: true,
			NextTalk:    "",
		})
	})
	if err != nil {
		return nil, err
	}
	return &api.CompleteTalkResponse{}, nil
}
