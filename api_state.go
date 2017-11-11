package main

import (
	"time"

	"cloud.google.com/go/datastore"
	"github.com/soulplant/jim-tracker/api"
	context "golang.org/x/net/context"
)

// An apiService backed by datastore
type dsApiService struct {
	key    *datastore.Key
	client *datastore.Client
}

// NewDsApiService constructs an api service that is backed by the given.
func NewDsApiService(client *datastore.Client) *dsApiService {
	key := datastore.NameKey("jimtracker", "root", nil)
	key.Namespace = *namespace
	return &dsApiService{key, client}
}

func (s *dsApiService) Close() {
	s.client.Close()
}

var recentQueries = datastore.NewQuery("Delivery").Namespace(*namespace).Order("-date").Limit(50)

func deliveryKey(name string) *datastore.Key {
	key := datastore.NameKey("Delivery", name, nil)
	key.Namespace = *namespace
	return key
}

func (s *dsApiService) FetchAll(ctx context.Context, req *api.FetchAllRequest) (*api.FetchAllResponse, error) {
	q := datastore.NewQuery("Delivery").Namespace(*namespace).Order("-Date").Limit(50)
	var resp []*api.Delivery
	if _, err := s.client.GetAll(ctx, q, &resp); err != nil {
		return nil, err
	}
	return &api.FetchAllResponse{
		Delivery: resp,
	}, nil
}

func (s *dsApiService) RecordDelivery(ctx context.Context, in *api.RecordDeliveryRequest) (*api.RecordDeliveryResponse, error) {
	date := time.Unix(in.Date, 0).Format("20060102")
	_, err := s.client.Put(ctx, deliveryKey(date), &api.Delivery{
		Date: date,
		Time: in.Date,
	})
	if err != nil {
		return nil, err
	}
	return &api.RecordDeliveryResponse{}, nil
}
