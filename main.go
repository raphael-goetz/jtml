package main

import "jtml/jtml"

type Road struct {
	Name   string
	Number int
}

func main() {
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	html, err := jtml.ParseAny(roads)
	if err != nil {
		panic(err)
	}

	print(html)
}
