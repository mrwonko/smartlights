package config

type ID int

const (
	idShelfRed ID = iota
	idShelfGreen
	idShelfBlue
	_ // idBar
	idSofa
	_ // idUV
	idShoe
	idLP
	idKitchenSink
	_ // this index is cursed for some reason, Google won't send us requests for it (WTF?!?)
	idLiquorGreen
	idLiquorBlue
	idLiquorRed
)

// IDs of the Raspberries
const (
	RaspiLight        = 0
	RaspiLightKitchen = 1
)

type Light struct {
	Name string
	Pi   int
	GPIO uint8
}

// TODO: rename to Devices, support different types
var Lights = map[ID]*Light{
	idShelfRed: {
		Name: "Red Shelf Light",
		Pi:   RaspiLight,
		GPIO: 17,
	},
	idShelfGreen: {
		Name: "Green Shelf Light",
		Pi:   RaspiLight,
		GPIO: 22,
	},
	idShelfBlue: {
		Name: "Blue Shelf Light",
		Pi:   RaspiLight,
		GPIO: 27,
	},
	idSofa: {
		Name: "Journey Light",
		Pi:   RaspiLight,
		GPIO: 10,
	},
	idShoe: {
		Name: "Shoe Light",
		Pi:   RaspiLight,
		GPIO: 9,
	},
	idLP: {
		Name: "TV Light",
		Pi:   RaspiLight,
		GPIO: 4,
	},
	idKitchenSink: {
		Name: "Sink Light",
		Pi:   RaspiLightKitchen,
		GPIO: 4,
	},
	idLiquorRed: {
		Name: "Red Liquor Light",
		Pi:   RaspiLightKitchen,
		GPIO: 27,
	},
	idLiquorGreen: {
		Name: "Green Liquor Light",
		Pi:   RaspiLightKitchen,
		GPIO: 22,
	},
	idLiquorBlue: {
		Name: "Blue Liquor Light",
		Pi:   RaspiLightKitchen,
		GPIO: 17,
	},
}

var Pis = func(lights map[ID]*Light) map[int]struct{} {
	res := map[int]struct{}{}
	for _, l := range lights {
		res[l.Pi] = struct{}{}
	}
	return res
}(Lights)
