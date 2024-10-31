package grpc

import (
	"context"

	"github.com/cornelia247/nyxfilms/gen"
	"github.com/cornelia247/nyxfilms/internal/grpcutil"
	"github.com/cornelia247/nyxfilms/pkg/discovery"
	"github.com/cornelia247/nyxfilms/rating/pkg/model"
)

// Gateway defines a gRPC gateway for rating service.
type Gateway struct {
	registry discovery.Registry
}

// New creates a new instance of the gRPC gateway for rating service.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}

}

// GetAggregatedRating returns the aggregated rating for a record or ErrNotFound.
func(g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64,error) {
	conn, err := grpcutil.ServiceConnection(ctx, "rating",g.registry)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	client := gen.NewRatingServiceClient(conn)
	resp, err := client.GetAggregatedRating(ctx, &gen.GetAggregatedRatingRequest{RecordId: string(recordID),RecordType: string(recordType)})
	if err != nil {
		return 0, err
	}
	return resp.RatingValue, nil
}
