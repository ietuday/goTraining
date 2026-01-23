package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInMemRepository()

	svc := service.NewService(inmemRepo)
	fare:= &domain.RideFareModel{
		UserID: "42",
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err!= nil {
		log.Println(err)

	}

	log.Println(t)
	//temporary 
	for {
		time.Sleep(time.Second)
	}
}
