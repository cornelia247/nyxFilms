package grpc

import (
	"context"
	"errors"

	"github.com/cornelia247/nyxfilms/metadata/pkg/model"
	"github.com/cornelia247/nyxfilms/gen"
	"github.com/cornelia247/nyxfilms/metadata/internal/controller/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a film metadata gRPC handler
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	ctrl *metadata.Controller
}

// New creates a new film metadata gRPC handler.
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}

}

// GetMetadata returns a film metadata by id.
func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil || req.FilmId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.FilmId)
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gen.GetMetadataResponse{Metadata: model.MetadataToProto(m)}, nil

}
