package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent"
)

// Pointers to int/string to allow for 'null' value in JSON
type Image struct {
	Id           int64   `json:"id" db:"id"`
	Width        int     `json:"width" db:"width"`
	Height       int     `json:"height" db:"height"`
	Url          string  `json:"url" db:"url"`
	Title        *string `json:"title" db:"title"`
	Artist       *string `json:"artist" db:"artist"`
	Gallery      *string `json:"gallery" db:"gallery"`
	Organization *string `json:"organization" db:"organization"`
}

var ImageSegment = newrelic.DatastoreSegment{
	Collection:         "images",
	Operation:          "SELECT",
	ParameterizedQuery: makeQuery(Image{}, "FROM images i WHERE i.media IS NOT NULL"),
}

type ImageMap map[int64]Image

func (ds *DbService) loadImages(txn newrelic.Transaction) {
	var err error
	var image Image
	var rows *sqlx.Rows

	if txn != nil {
		ImageSegment.StartTime = newrelic.StartSegmentNow(txn)
		defer ImageSegment.End()
	}

	// Query images
	Log.Debug.Printf("ImageQuery: %s\n", ImageSegment.ParameterizedQuery)
	rows, err = ds.Queryx(ImageSegment.ParameterizedQuery)
	if err != nil {
		Log.Error.Printf("Load Images: %v\n", err)
		return
	}

	// Load into ImageMap
	ds.ImageMap = make(ImageMap)
	for rows.Next() {
		err = rows.StructScan(&image)
		if err != nil {
			Log.Error.Printf("Load Images: %v\n", err)
			continue
		}
		ds.ImageMap[image.Id] = image
	}
	Log.Info.Printf("Load Images: %d entries total\n", len(ds.ImageMap))
}
