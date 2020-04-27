package config

type ID int

const (
	_ ID = iota // idRed
	_           // idGreen
	_           // idBlue
	_           // idBar
	idSofa
	_ // idUV
	idShoe
	idLP
	idKitchenSink
)

const (
	raspiLight        = 0
	raspiLightKitchen = 1
)

type Light struct {
	Name string
	Pi   int
	GPIO uint8
}

var Lights = map[ID]*Light{
	idSofa: {
		Name: "Journey Light",
		Pi:   raspiLight,
		GPIO: 10,
	},
	idShoe: {
		Name: "Shoe Light",
		Pi:   raspiLight,
		GPIO: 9,
	},
	idLP: {
		Name: "TV Light",
		Pi:   raspiLight,
		GPIO: 4,
	},
	idKitchenSink: {
		Name: "Sink Light",
		Pi:   raspiLightKitchen,
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
