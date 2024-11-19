package parser_test

import (
	"testing"

	"github.com/NSXBet/edne/internal/parser"
	"github.com/NSXBet/edne/test"
	"github.com/stretchr/testify/require"
)

func TestParseLogradouro(t *testing.T) {
	base := test.Fixture("base")
	require.NotEmpty(t, base)
	update := test.Fixture("update")
	require.NotEmpty(t, update)

	parser := parser.NewLogradouroParser()
	addresses, err := parser.Parse(base, update)
	require.NoError(t, err)
	require.NotEmpty(t, addresses)
	require.Len(t, addresses, 450)

	require.Contains(t, addresses, "70800122")
	addr := addresses["70800122"]
	require.NotNil(t, addr)

	// 1005314@DF@1778@1128@@SCEN Trecho 2 Conjunto 4@@70800122@Trecho@N@SCEN Tr 2 Cj 4
	require.Equal(t, "1005314", addr.ID)
	require.Equal(t, "DF", addr.State)
	require.Equal(t, "1778", addr.LocationID)

	require.NotNil(t, addr.StartingNeighborhood)
	require.Equal(t, "1128", addr.StartingNeighborhood.ID)

	require.NotNil(t, addr.EndingNeighborhood)
	require.Equal(t, "", addr.EndingNeighborhood.ID)

	require.Equal(t, "SCEN Trecho 2 Conjunto 4", addr.Name)
	require.Equal(t, "", addr.Complement)
	require.Equal(t, "70800122", addr.ZipCode)
	require.Equal(t, "Trecho", addr.Type)

	// 1303878@SP@9052@17217@@Otávio Gouveia@@15810115@Rua@S@R Otávio Gouveia@INS@
	require.Contains(t, addresses, "15810115")
	addr = addresses["15810115"]
	require.NotNil(t, addr)
	require.Equal(t, "1303878", addr.ID)
	require.Equal(t, "SP", addr.State)
	require.Equal(t, "9052", addr.LocationID)
	require.Equal(t, "17217", addr.StartingNeighborhood.ID)
	require.Equal(t, "Otávio Gouveia", addr.Name)
	require.Equal(t, "", addr.Complement)
	require.Equal(t, "15810115", addr.ZipCode)
	require.Equal(t, "Rua", addr.Type)
}
