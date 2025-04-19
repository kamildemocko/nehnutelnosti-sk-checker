package main

import "nehnutelnosti-sk/src/internal/uri"

func main() {
	uri1, err := uri.NewUrlBuilder().
		WithPlace("kosice").
		WithSize("3-izbove-byty").
		WithArea(50, 100).
		WithPrice(100000, 200000).
		Build()
	if err != nil {
		panic("wrong parameters for uri")
	}

	app := App{
		uri: []string{
			uri1,
		},
	}

	app.CheckUpdated()
}
