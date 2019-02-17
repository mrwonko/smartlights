package config

type ID int

const (
	idRed ID = iota
	idGreen
	idBlue
	idBar
	idSofa
	idUV
)

type Light struct {
	Name string
	Pi   int
	GPIO int
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
}
