package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Graph struct {
	Metadata Metadata `json:"metadata"`
	Objects  Objects  `json:"objects"`
}

type Metadata struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Objects struct {
	Box  []Box  `json:"box"`
	Line []Line `json:"line"`
	Text []Text `json:"text"`
}

type Line struct {
	PrimitiveCoords
	Type string `json:"type"`
}
type Text struct {
	PrimitiveCoords
	Text string `json:"text"`
}

type PrimitiveCoords struct {
	Coords []int `json:"coords"`
}

type Primitive interface {
	Draw() (int, int, RuneMap)
	// TODO
	// Validate() error
}

type Canvas [][]string

// Map function over slice
func map2[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))
	for _, e := range data {
		res = append(res, f(e))
	}
	return res
}

// Represent [][]rune as string (with newlines)
func (rm RuneMap) String() string {
	return strings.Join(map2(rm, func(r []rune) string {
		return string(r)
	}), "\n")
}

type RuneMap [][]rune

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("example.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var graph Graph

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &graph)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	for _, v := range graph.Objects.Box {
		x, y, text := v.Draw()
		fmt.Printf("Box at (%v,%v):\n%v\n", x, y, text)
	}

}
