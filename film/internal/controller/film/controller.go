package film

import (
	"context"
	"errors"

	"github.com/cornelia247/nyxfilms/film/internal/gateway"
	"github.com/cornelia247/nyxfilms/film/pkg/model"
	metadatamodel "github.com/cornelia247/nyxfilms/metadata/pkg/model"
	ratingmodel "github.com/cornelia247/nyxfilms/rating/pkg/model"
)

// ErrNotFound is returned when the film metadata is not found.
var ErrNotFound = errors.New("film metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)

	// PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a new film service controller.
func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway, metadataGateway}
}

// Get returns the film details including the aggregated rating and film metadata
func (c *Controller) Get(ctx context.Context, id string) (*model.FilmDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.FilmDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeFilm)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// proceed, it's fine not to have ratings yet
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
