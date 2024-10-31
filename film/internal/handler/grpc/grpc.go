package grpc

import (
	"context"
	"errors"

	"github.com/cornelia247/nyxfilms/film/internal/controller/film"
	"github.com/cornelia247/nyxfilms/metadata/pkg/model"

	"github.com/cornelia247/nyxfilms/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler is the gRPC film service handler.
type Handler struct {
	gen.UnimplementedFilmServiceServer
	ctrl *film.Controller
}

// New creates a new instance of the gRPC handler for film service.
func New(ctrl *film.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetFilmDetails returns all datafrom the ratings and film metadata
func (h *Handler) GetFilmDetails(ctx context.Context, req *gen.GetFilmDetailsRequest) (*gen.GetFilmDetailsResponse, error) {
	if req == nil || req.FilmId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.FilmId)
	if err != nil && errors.Is(err, film.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())

	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetFilmDetailsResponse{FilmDetails: &gen.FilmDetails{
		Metadata: model.MetadataToProto(&m.Metadata),
		Rating:   float32(*m.Rating),
	}}, nil

}
