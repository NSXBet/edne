package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/NSXBet/edne/internal/models"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type State string

type LogradouroParserOption func(opts *ParserOptions)

type ParserOptions struct {
	States []State
}

var defaultStates = []State{
	"AC", // Acre
	"AL", // Alagoas
	"AP", // Amapá
	"AM", // Amazonas
	"BA", // Bahia
	"CE", // Ceará
	"DF", // Distrito Federal
	"ES", // Espírito Santo
	"GO", // Goiás
	"MA", // Maranhão
	"MT", // Mato Grosso
	"MS", // Mato Grosso do Sul
	"MG", // Minas Gerais
	"PA", // Pará
	"PB", // Paraíba
	"PR", // Paraná
	"PE", // Pernambuco
	"PI", // Piauí
	"RJ", // Rio de Janeiro
	"RN", // Rio Grande do Norte
	"RS", // Rio Grande do Sul
	"RO", // Rondônia
	"RR", // Roraima
	"SC", // Santa Catarina
	"SP", // São Paulo
	"SE", // Sergipe
	"TO", // Tocantins
}

func WithStates(states ...State) LogradouroParserOption {
	return func(opts *ParserOptions) {
		opts.States = states
	}
}

type LogradouroParser struct {
	states []State
}

func NewLogradouroParser(opts ...LogradouroParserOption) *LogradouroParser {
	options := &ParserOptions{}

	parser := &LogradouroParser{}
	for _, opt := range opts {
		opt(options)
	}

	parser.states = options.States
	if len(parser.states) == 0 {
		parser.states = defaultStates
	}

	return parser
}

func (p *LogradouroParser) Parse(basePath, updatePath string) (map[string]models.Address, error) {
	baseAddresses, err := p.parseFiles(basePath, "LOG")
	if err != nil {
		return nil, fmt.Errorf("error parsing base file: %w", err)
	}

	if updatePath != "" {
		updateAddresses, err := p.parseFile(updatePath, "DELTA_LOG_LOGRADOURO.TXT")
		if err != nil {
			return nil, fmt.Errorf("error parsing update file: %w", err)
		}

		for _, address := range updateAddresses {
			baseAddresses[address.ZipCode] = address
		}
	}

	return baseAddresses, nil
}

func (p *LogradouroParser) parseFiles(basePath, prefix string) (map[string]models.Address, error) {
	var addresses []models.Address

	// Create a map of valid states for O(1) lookup
	validStates := make(map[State]bool)
	for _, state := range p.states {
		validStates[state] = true
	}

	// Read directory entries
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", basePath, err)
	}

	// Process each file that matches our pattern
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), fmt.Sprintf("%s_LOGRADOURO_", prefix)) ||
			!strings.HasSuffix(entry.Name(), ".TXT") {
			continue
		}

		fileAddresses, err := p.parseFile(basePath, entry.Name())
		if err != nil {
			return nil, fmt.Errorf("error parsing file %s: %w", entry.Name(), err)
		}

		addresses = append(addresses, fileAddresses...)
	}

	return models.ZipCodeMap(addresses), nil
}

func (p *LogradouroParser) parseFile(basePath, filename string) ([]models.Address, error) {
	addresses := make([]models.Address, 0)

	filepath := path.Join(basePath, filename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filepath, err)
	}
	defer file.Close()

	dec := transform.NewReader(file, charmap.Windows1252.NewDecoder())

	reader := csv.NewReader(dec)
	reader.Comma = '@'
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record from %s: %w", filepath, err)
		}

		// Ensure we have minimum required fields
		if len(record) < 9 {
			continue
		}

		address := models.Address{
			ID:         strings.TrimSpace(record[0]),
			State:      strings.TrimSpace(record[1]),
			LocationID: strings.TrimSpace(record[2]),
			StartingNeighborhood: &models.Neighborhood{
				ID: strings.TrimSpace(record[3]),
			},
			EndingNeighborhood: &models.Neighborhood{
				ID: strings.TrimSpace(record[4]),
			},
			Name:       strings.TrimSpace(record[5]),
			Complement: strings.TrimSpace(record[6]),
			ZipCode:    strings.TrimSpace(record[7]),
			Type:       strings.TrimSpace(record[8]),
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}
