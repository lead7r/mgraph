package main

import "time"

type film struct {
	title            string
	description      string
	shortDescription string
}

type genre struct {
	aType       string
	name        string
	description string
}

type series struct {
	title        string
	seasonCount  int
	chapterCount int
}

type season struct {
	number int
	name   string
}

type chapter struct {
	number      int
	name        string
	description string
	images      []string
}

type user struct {
	id             int64
	username       string
	profilePicture []byte
}

type viewed struct {
	rating byte
}

type isFriendOf struct {
	since time.Time
}

type isComposedBy struct {
}

type isOfGenre struct {
}
