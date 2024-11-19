package edne

import "github.com/NSXBet/edne/internal/models"

type (
	Address           = models.Address
	LocationSituation = models.LocationSituation
	LocationType      = models.LocationType
	Location          = models.Location
	Street            = models.Street
	Neighborhood      = models.Neighborhood
)

const (
	LocationSituationNonCodified      = models.LocationSituationNonCodified
	LocationSituationCodifiedStreet   = models.LocationSituationCodifiedStreet
	LocationSituationCodifiedDistrict = models.LocationSituationCodifiedDistrict
)

const (
	LocationTypeDistrict = models.LocationTypeDistrict
	LocationTypeCity     = models.LocationTypeCity
	LocationTypeVillage  = models.LocationTypeVillage
)
