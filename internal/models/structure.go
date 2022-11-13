package models

import (
	"encoding/json"

	"01.alem.school/git/Nurzhas/groupie-tracker/internal/service"
)

type InfoErr struct {
	StatusNumberText string
	StatusNumber     int
}

const (
	ArtistsURL = "https://groupietrackers.herokuapp.com/api/artists"
)

type Artists struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	ConcertDates string              `json:"concertDates"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relation     string              `json:"relations"`
	Concerts     map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Concerts map[string][]string `json:"datesLocations"`
}

func (artist *Artists) InfoConcert() error {
	r := Relation{}
	body, err := service.GetJson(artist.Relation)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return err
	}
	artist.Concerts = r.Concerts
	return nil
}
