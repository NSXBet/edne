package parser_test

import (
	"testing"

	"github.com/NSXBet/edne/internal/parser"
	"github.com/NSXBet/edne/test"
	"github.com/stretchr/testify/require"
)

func TestParseNeighborhood(t *testing.T) {
	base := test.Fixture("base")
	require.NotEmpty(t, base)
	update := test.Fixture("update")
	require.NotEmpty(t, update)

	parser := parser.NewNeighborhoodParser()

	neighborhoods, err := parser.Parse(base, update)
	require.NoError(t, err)
	require.NotEmpty(t, neighborhoods)
	require.Len(t, neighborhoods, 55)

	// 5268@MG@3689@Santa Paula@Sta Paula
	require.Contains(t, neighborhoods, 5268)
	neighborhood := neighborhoods[5268]
	require.NotNil(t, neighborhood)
	require.Equal(t, 5268, neighborhood.ID)
	require.Equal(t, "Santa Paula", neighborhood.Name)

	// 75324@SP@9009@Área Industrial Senhor Antônio Gasparini@A Ind Sr Antônio Gasparini@UPD
	require.Contains(t, neighborhoods, 75324)
	neighborhood = neighborhoods[75324]
	require.NotNil(t, neighborhood)
	require.Equal(t, 75324, neighborhood.ID)
	require.Equal(t, "Área Industrial Senhor Antônio Gasparini", neighborhood.Name)
}
