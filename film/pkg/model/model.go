package model

import (
	"github.com/cornelia247/nyxfilms/metadata/pkg/model"
)

// FilmDetails include film metadata and it's aggregated rating.

type FilmDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
