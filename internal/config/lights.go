package config

import (
	"fmt"
	"strconv"
)

type ID int

const (
	idRed ID = iota
	idGreen
	idBlue
	idBar
	idSofa
	idUV
	idShoe
	idLP

	idLast
)

func ParseID(raw string) (ID, *Light, error) {
	i, err := strconv.Atoi(raw)
	if err != nil {
		return -1, nil, err
	}
	id := ID(i)
	l := Lights[id]
	if l == nil {
		return -1, nil, fmt.Errorf("unknown ID %d", i)
	}
	return id, l, nil
}

type Light struct {
	Name string
	Pi   int
	GPIO uint8
}

var Lights = map[ID]*Light{
	idRed: {
		Name: "Red",
		Pi:   0,
		GPIO: 13,
	},
	idGreen: {
		Name: "Green",
		Pi:   0,
		GPIO: 19,
	},
	idBlue: {
		Name: "Blue",
		Pi:   0,
		GPIO: 18,
	},
	idBar: {
		Name: "Bar",
		Pi:   0,
		GPIO: 16,
	},
	idSofa: {
		Name: "Sofa",
		Pi:   0,
		GPIO: 20,
	},
	idUV: {
		Name: "UV",
		Pi:   0,
		GPIO: 21,
	},
	idShoe: {
		Name: "Shoe Light",
		Pi:   0,
		GPIO: 12,
	},
	idLP: {
		Name: "LP Light",
		Pi:   0,
		GPIO: 26,
	},
}

var Pis = func(lights map[ID]*Light) map[int]struct{} {
	res := map[int]struct{}{}
	for _, l := range lights {
		res[l.Pi] = struct{}{}
	}
	return res
}(Lights)
