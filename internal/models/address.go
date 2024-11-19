package models

type LocationSituation int

const (
	LocationSituationNonCodified LocationSituation = iota
	LocationSituationCodifiedStreet
	LocationSituationCodifiedDistrict
)

type LocationType string

const (
	LocationTypeDistrict LocationType = "D"
	LocationTypeCity     LocationType = "M"
	LocationTypeVillage  LocationType = "P"
)

type Location struct {
	ID                    int
	State                 string
	Name                  string
	ZipCode               int
	Situation             LocationSituation
	Type                  LocationType
	SubordinateLocationID int
	IBGECode              string
}

type Street struct {
	ID                   int
	ZipCode              int
	State                string
	LocationID           int
	StartingNeighborhood *Neighborhood
	EndingNeighborhood   *Neighborhood
	Name                 string
	Complement           string
	Type                 string
}

type Neighborhood struct {
	ID   int
	Name string
}

func ZipCodeMap(addrs []Street) map[int]Street {
	m := make(map[int]Street)
	for _, addr := range addrs {
		m[addr.ZipCode] = addr
	}

	return m
}

func NeighborhoodMap(neighborhoods []Neighborhood) map[int]Neighborhood {
	m := make(map[int]Neighborhood)
	for _, neighborhood := range neighborhoods {
		m[neighborhood.ID] = neighborhood
	}
	return m
}

func LocationMap(locations []Location) map[int]Location {
	m := make(map[int]Location)
	for _, location := range locations {
		m[location.ID] = location
	}
	return m
}

type Address struct {
	Street       string
	Neighborhood string
	City         string
	CityIBGECode string
	State        string
	ZipCode      int
}
