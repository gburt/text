// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate gotext extract

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.English)

	p.Print("Hello world!\n")

	p.Println("Hello", "world!")

	person := "Sheila"
	place := "Zürich"

	p.Print("Hello ", person, " in ", place, "!\n")

	// Greet a city.
	p.Print("Hello city!\n")

	city := "Amsterdam"
	// Greet a city.
	p.Printf("Hello %s!\n", city)

	town := "Amsterdam"
	// Greet a town.
	p.Printf("Hello %s!\n",
		town, // Town
	)

	// Person visiting a place.
	p.Printf("%s is visiting %s!\n",
		person, // The person of matter.
		place,  // Place the person is visiting.
	)

	pp := struct {
		Person string // The person of matter. // TODO: get this comment.
		Place  string
	}{
		person, place,
	}

	// extract will drop this comment in favor of the one below.
	p.Printf("%s is visiting %s!\n", // Person visiting a place.
		pp.Person,
		pp.Place, // Place the person is visiting.
	)

	// Numeric literal
	p.Printf("%d files remaining!", 2)

	const n = 2

	// Numeric var
	p.Printf("%d more files remaining!", n)

	type referralCode int

	c := referralCode(5)
	p.Printf("Use the following code for your discount: %d\n", c)
}