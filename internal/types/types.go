package types

import "time"

type Contact struct {
	Name string
	Icon string
	Link string
}

type Blog struct {
	PublishDate time.Time
	Title       string
	Tags        []string
	TimeToRead  time.Duration
}
type Project struct {
	PublishDate time.Time
	Title       string
	ProjectURL  string
	Tags        []string
}
