package uri

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUrlBuildPlaceMissing = errors.New("place value missing, fe 'kosice'")
	ErrUrlBuildSizeMissing  = errors.New("size value missing, ex. '3-izbove-byty'")
)

type Uri struct {
	root      string
	place     string
	size      string
	priceFrom int
	priceTo   int
	areaFrom  int
	areaTo    int
}

func NewUrlBuilder() *Uri {
	return &Uri{root: "https://www.nehnutelnosti.sk/vysledky/$size/$place/predaj"}
}

// place - city, ex. "kosice"
func (u *Uri) WithPlace(place string) *Uri {
	u.place = place
	return u
}

// size - type, ex. "3-izbove byty"
func (u *Uri) WithSize(size string) *Uri {
	u.size = size
	return u
}

// sets price from and to
func (u *Uri) WithPrice(from, to int) *Uri {
	u.priceFrom = from
	u.priceTo = to
	return u
}

// sets area from and to
func (u *Uri) WithArea(from, to int) *Uri {
	u.areaFrom = from
	u.areaTo = to
	return u
}

// builds url and returns it
func (u *Uri) Build() (string, error) {
	if u.place == "" {
		return "", ErrUrlBuildPlaceMissing
	}

	if u.size == "" {
		return "", ErrUrlBuildSizeMissing
	}

	var url string

	switch {
	case u.areaFrom != 0 && u.priceFrom != 0:
		url = fmt.Sprintf(
			"%s?areaTo=%d&priceTo=%d&areaFrom=%d&priceFrom=%d&order=NEWEST",
			u.root,
			u.areaTo,
			u.priceTo,
			u.areaFrom,
			u.priceFrom,
		)
	case u.areaFrom != 0:
		url = fmt.Sprintf(
			"%s?areaTo=%d&areaFrom=%d",
			u.root,
			u.areaTo,
			u.areaFrom,
		)
	case u.priceFrom != 0:
		url = fmt.Sprintf(
			"%s?priceTo=%d&priceFrom=%d",
			u.root,
			u.priceTo,
			u.priceFrom,
		)
	}

	url = strings.Replace(url, "$place", u.place, 1)
	url = strings.Replace(url, "$size", u.size, 1)

	fmt.Println("uri: ", url)

	return url, nil
}
