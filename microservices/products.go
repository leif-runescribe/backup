package data

import "time"

type Unit struct {
	unitID           int
	name             string
	description      string
	price            float32
	timeOfRecruiting string
	timeOfDisbanding string
}

var productsList = []*Product{
	&Product{
		productID:      10001,
		name:           "Stark bowmen",
		description:    "Band of 20 Stark bowmen, from a variety of Nothern banners of House Stark. Good mobility, ranged strength.",
		price:          "200 Talents",
		timeOfCreation: time.Now().UTC().String(),
		timeOfDisbanding: time.Now().UTC().String(),
	},
	&Product{
		productID:      10001,
		name:           "Stark bowmen",
		description:    "Band of 20 Stark bowmen, from a variety of Nothern banners of House Stark. Good mobility, ranged strength.",
		price:          "200 Talents",
		timeOfCreation: time.Now().UTC().String(),
		timeOfDisbanding: time.Now().UTC().String(),
	},
}