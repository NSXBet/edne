package parser_test

import (
	"testing"

	"github.com/NSXBet/edne/internal/parser"
	"github.com/NSXBet/edne/test"
	"github.com/stretchr/testify/require"
)

func TestMasterParser(t *testing.T) {
	base := test.Fixture("base")
	require.NotEmpty(t, base)
	update := test.Fixture("update")
	require.NotEmpty(t, update)

	parser := parser.NewMasterParser()
	addresses, err := parser.Parse(base, update)
	require.NoError(t, err)
	require.NotEmpty(t, addresses)
	require.Len(t, addresses, 450)

	zipCode := 6415235
	require.Contains(t, addresses, zipCode)
	addr := addresses[zipCode]
	require.NotNil(t, addr)

	require.Equal(t, "Rua", addr.StreetType)
	require.Equal(t, "São Bento", addr.Street)
	require.Equal(t, "Mosteiro Bento", addr.City)
	require.Equal(t, "Cruz de São Bento", addr.Neighborhood)
	require.Equal(t, zipCode, addr.ZipCode)
	require.Equal(t, "SP", addr.State)
	require.Equal(t, "06415235", addr.CityIBGECode)
}
