package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
	"github.com/Zahri-Kargo/kargo-trucks/graph/generated"
	"github.com/Zahri-Kargo/kargo-trucks/graph/model"
	"encoding/csv"
	"log"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	// panic(fmt.Errorf("not implemented"))

	truck := &model.Truck{

		ID: fmt.Sprintf("TRUCK-%d", len(r.Truck)+1),

		PlateNo: plateNo,
	}

	r.Truck = append(r.Truck, truck)

	return truck, nil
}

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, truckID *string) (*model.Shipment, error) {
	shipment := &model.Shipment{
		ID:           fmt.Sprintf("SHIPMENT-%d", len(r.Shipment)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: deliveryDate,
		Trucks:       &model.Truck{ID: *truckID},
	}
	r.Shipment = append(r.Shipment, shipment)

	return shipment, nil
}

func (r *mutationResolver) DeleteTruck(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteShipment(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendTruckDatatoEmail(ctx context.Context, email string) (*model.Email, error) {
	fmt.Println(email)
	emails :=&model.Email{
		Email:	email,
	}

	from := "zahri.rusli@gmail.com"
	password := os.Getenv("EMAILPASSWORD")

	// Receiver email address.
	to := []string{
		email,
		"zahri.rusli@gmail.com",
	}

	// // smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// // Message.
	message := []byte("This is a test email message.")

	// // Authentication.
	auth := smtp.PlainAuth("", from, password,
		smtpHost)

	// // Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email Sent Successfully!")
	return emails, nil	
}
func (r *queryResolver) PaginatedTrucks(ctx context.Context, id *string, plateNo *string, page int, first int)([]*model.Truck, error) {
	// Sender data.

	return r.Truck, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context, id *string, origin *string, destination *string, page int, first int) ([]*model.Shipment, error) {
	return r.Shipment, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *Resolver) Init() {
	for i := 0; i < 100; i++ {
		truck := &model.Truck{
			ID:      fmt.Sprintf("TRUCK-%d", len(r.Truck)+1),
			PlateNo: fmt.Sprintf("B %d CD", len(r.Truck)+1),
		}
		r.Truck = append(r.Truck, truck)
	}
}
