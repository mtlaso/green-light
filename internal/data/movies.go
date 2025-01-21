package data

import "time"

// Movie represents a movie record.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`                        // When added to the database.
	Title     string    `json:"title"`                    // Movie title.
	Year      int32     `json:"year,omitempty"`           // Release year.
	Runtime   int32     `json:"runtime,omitempty,string"` // Movie runtime in minutes.
	Genres    []string  `json:"genres,omitempty"`         // List of genres.
	Version   int32     `json:"version"`                  // The version number starts at 1 and will be incremented each time the movie information is updated.
}
