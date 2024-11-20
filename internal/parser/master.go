package parser

import (
	"fmt"

	"github.com/NSXBet/edne/internal/models"
)

type MasterParser struct{}

func NewMasterParser() *MasterParser {
	return &MasterParser{}
}

func (p *MasterParser) Parse(base, update string) (map[int]models.Address, error) {
	neighborhoodParser := NewNeighborhoodParser()

	neighborhoods, err := neighborhoodParser.Parse(base, update)
	if err != nil {
		return nil, fmt.Errorf("error parsing neighborhoods: %w", err)
	}

	locationParser := NewLocationParser()
	locations, err := locationParser.Parse(base, update)
	if err != nil {
		return nil, fmt.Errorf("error parsing locations: %w", err)
	}

	streetParser := NewStreetParser()
	streets, err := streetParser.Parse(base, update)
	if err != nil {
		return nil, fmt.Errorf("error parsing streets: %w", err)
	}

	addresses := map[int]models.Address{}

	for zipCode, street := range streets {
		neighborhood, ok := neighborhoods[street.StartingNeighborhood.ID]
		if !ok {
			neighborhood = models.Neighborhood{}
		}

		location, ok := locations[street.LocationID]
		if !ok {
			location = models.Location{}
		}

		addresses[zipCode] = models.Address{
			StreetType:   street.Type,
			Street:       street.Name,
			Neighborhood: neighborhood.Name,
			City:         location.Name,
			CityIBGECode: location.IBGECode,
			State:        location.State,
			ZipCode:      zipCode,
		}
	}

	return addresses, nil
}
