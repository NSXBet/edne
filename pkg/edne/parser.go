package edne

import "github.com/NSXBet/edne/internal/parser"

type Parser struct{}

func (p *Parser) Parse(base, update string) (map[int]Address, error) {
	masterParser := parser.NewMasterParser()

	addresses, err := masterParser.Parse(base, update)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}
