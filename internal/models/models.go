package models

import (
	"time"

	"github.com/google/uuid"
)

type SourcesData struct {
	ID          *uuid.UUID
	Name        *string
	URL         *string
	LastFetched *time.Time
}

type FeedData struct {
	ID              *uuid.UUID
	Title           *string
	Link            *string
	Description     *string
	PublicationDate *time.Time
	SourceId        *int
	Content         *string
	Author          *string
	Categories      *string
	GUID            *string
}

type FeedAttributeData struct {
	ID     *uuid.UUID
	FeedID *uuid.UUID
	Name   *string
	Value  *string
}
