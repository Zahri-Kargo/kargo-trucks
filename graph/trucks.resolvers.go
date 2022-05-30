package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Zahri-Kargo/kargo-trucks/graph/generated"
	"github.com/Zahri-Kargo/kargo-trucks/graph/model"
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

func (r *queryResolver) PaginatedTrucks(ctx context.Context, id *string, plateNo *string, page int, first int) ([]*model.Truck, error) {
	// panic(fmt.Errorf("not implemented"))

	return r.Truck, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context, id *string, origin *string, destination *string, page int, first int) ([]*model.Shipment, error) {
	return r.Shipment, nil
	// panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
