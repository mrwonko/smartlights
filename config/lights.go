package config

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
)

type Light struct {
	Name string
	Pi   int
	GPIO uint8
}

var Lights = map[ID]*Light{
	idSofa: {
		Name: "Journey Light",
		Pi:   0,
		GPIO: 10,
	},
	idShoe: {
		Name: "Shoe Light",
		Pi:   0,
		GPIO: 9,
	},
	idLP: {
		Name: "TV Light",
		Pi:   0,
		GPIO: 4,
	},
}

var Pis = func(lights map[ID]*Light) map[int]struct{} {
	res := map[int]struct{}{}
	for _, l := range lights {
		res[l.Pi] = struct{}{}
	}
	return res
}(Lights)
