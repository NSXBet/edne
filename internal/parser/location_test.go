package parser_test

import (
	"testing"

	"github.com/NSXBet/edne/internal/models"
	"github.com/NSXBet/edne/internal/parser"
	"github.com/NSXBet/edne/test"
	"github.com/stretchr/testify/require"
)

func TestParseLocation(t *testing.T) {
	base := test.Fixture("base")
	require.NotEmpty(t, base)
	update := test.Fixture("update")
	require.NotEmpty(t, update)

	parser := parser.NewLocationParser()

	locations, err := parser.Parse(base, update)
	require.NoError(t, err)
	require.NotEmpty(t, locations)
	require.Len(t, locations, 22)

	require.Contains(t, locations, 16)
	location := locations[16]
	require.NotNil(t, location)

	// 16@AC@Rio Branco@@1@M@@Rio Branco@1200401
	require.Equal(t, 16, location.ID)
	require.Equal(t, "AC", location.State)
	require.Equal(t, "Rio Branco", location.Name)
	require.Equal(t, 0, location.ZipCode)
	require.Equal(t, "1200401", location.IBGECode)
	require.Equal(t, models.LocationSituationCodifiedStreet, location.Situation)
	require.Equal(t, models.LocationTypeCity, location.Type)
	require.Equal(t, 0, location.SubordinateLocationID)

	// 9858@RJ@Rio de Janeiro@@1@M@@Rio de Janeiro@1200401
	require.Contains(t, locations, 9858)
	location = locations[9858]
	require.NotNil(t, location)

	require.Equal(t, 9858, location.ID)
	require.Equal(t, "RJ", location.State)
	require.Equal(t, "Rio de Janeiro", location.Name)
	require.Equal(t, 0, location.ZipCode)
	require.Equal(t, "1200401", location.IBGECode)
	require.Equal(t, models.LocationSituationCodifiedStreet, location.Situation)
	require.Equal(t, models.LocationTypeCity, location.Type)
	require.Equal(t, 0, location.SubordinateLocationID)
}
