package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Position struct {
	pos_x, pos_y int
}

type SoilType struct {
	Name string
}

type Hydration struct {
	last_watered time.Time
}

type Garden struct {
	MaxX, MaxY, MinX, MinY int
	Soil                   *SoilType
	Grid                   [][]GridCell
}

func (self *Garden) Mulch() {
	for row := range self.Grid {
		for cell := range self.Grid[row] {
			self.Grid[row][cell].Mulched = true
		}
	}
}

type GridCell struct {
	Position  *Position
	PlantType string
	Mulched   bool
	Hyd       *Hydration
}

func GardenFunction(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Argh!")
	}

	var gd Garden
	err = json.Unmarshal(b, &gd)
	if err != nil {
		fmt.Println(err)
		panic("...damn")
	}

	s, err := json.Marshal(gd)
	if err != nil {
		fmt.Println(err)
		panic(":(")
	}
	ioutil.WriteFile(
		filename+".out.json",
		s,
		0644,
	)
}

func main() {
	GardenFunction("garden.json")
	GardenFunction("japanese_garden.json")
}
