package grpc

import (
	"context"

	"github.com/cornelia247/nyxfilms/gen"
	"github.com/cornelia247/nyxfilms/internal/grpcutil"
	"github.com/cornelia247/nyxfilms/metadata/pkg/model"
	"github.com/cornelia247/nyxfilms/pkg/discovery"
)

// Gateway defines a film's metadata gRPC gateway
type Gateway struct {
	registry discovery.Registry
}

// New creates a new gRPC fgateway for  film metadata service
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}

}

//  Get returns film metadata by film id
func(g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{FilmId: id})
	if err != nil {
		return nil, err
	}
	return model.MetadataFromProto(resp.Metadata), nil
}
