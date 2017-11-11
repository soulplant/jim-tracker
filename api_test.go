package main

import (
	"testing"
	"time"

	context "golang.org/x/net/context"

	"github.com/soulplant/jim-tracker/api"
)

// MakeNewTestService creates and inits a datastore-backed api service.
func MakeNewTestService(t *testing.T) *dsApiService {
	ds := InitTestDatastore(t)
	a := NewDsApiService(ds)
	return a
}

func TestStuff(t *testing.T) {
	a := MakeNewTestService(t)
	defer a.Close()
	ctx := context.Background()
	date := time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local)
	dateStr := date.Format(dateFormat)
	if _, err := a.RecordDelivery(ctx, &api.RecordDeliveryRequest{
		Delivery: &api.Delivery{
			Date: dateStr,
			Time: &api.LocalTime{Hour: 8, Minute: 30, Second: 0},
		},
	}); err != nil {
		t.Fatal("Failed to record delivery", err)
	}

	data := fetchAll(t, a)
	if len(data) != 1 {
		t.Fatalf("Expected 1 delivery, got %d", len(data))
	}
	if data[0].Date != dateStr {
		t.Errorf("Got '%s', but expected '%s' for the timestamp", data[0].Date, dateStr)
	}

	// Record a second delivery for one second later - it should get de-duped.
	if _, err := a.RecordDelivery(ctx, &api.RecordDeliveryRequest{
		Delivery: &api.Delivery{
			Date: dateStr,
			Time: &api.LocalTime{Hour: 8, Minute: 30, Second: 1},
		},
	}); err != nil {
		t.Fatal("Failed record delivery", err)
	}

	data = fetchAll(t, a)
	if len(data) != 1 {
		t.Fatalf("Expected 1 delivery, got %d", len(data))
	}
	if data[0].Time.Second != 1 {
		t.Error("Date is not the new date")
	}

	if _, err := a.ClearDelivery(ctx, &api.ClearDeliveryRequest{Date: dateStr}); err != nil {
		t.Fatal("Failed to clear the delivery", err)
	}
	data = fetchAll(t, a)
	if len(data) != 0 {
		t.Fatalf("Expected no deliveries to be recorded, got %d", len(data))
	}
}

func fetchAll(t *testing.T, a *dsApiService) []*api.Delivery {
	resp, err := a.FetchAll(context.Background(), &api.FetchAllRequest{})
	if err != nil {
		t.Fatal("Failed to fetch deliveries", err)
	}
	return resp.Delivery
}
