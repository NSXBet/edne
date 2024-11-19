package models

type Address struct {
	ID                   string
	ZipCode              string
	State                string
	LocationID           string
	StartingNeighborhood *Neighborhood
	EndingNeighborhood   *Neighborhood
	Name                 string
	Complement           string
	Type                 string
}

type Neighborhood struct {
	ID   string
	Name string
}

func ZipCodeMap(addrs []Address) map[string]Address {
	m := make(map[string]Address)
	for _, addr := range addrs {
		m[addr.ZipCode] = addr
	}

	return m
}

func NeighborhoodMap(neighborhoods []Neighborhood) map[string]Neighborhood {
	m := make(map[string]Neighborhood)
	for _, neighborhood := range neighborhoods {
		m[neighborhood.ID] = neighborhood
	}
	return m
}
