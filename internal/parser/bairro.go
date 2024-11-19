package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/NSXBet/edne/internal/models"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type BairroParser struct{}

func NewNeighborhoodParser() *BairroParser {
	return &BairroParser{}
}

func (p *BairroParser) Parse(base, update string) (map[int]models.Neighborhood, error) {
	// Read base file
	baseNeighborhoods, err := p.parseFile(base)
	if err != nil {
		return nil, fmt.Errorf("error parsing base file: %w", err)
	}

	// Read update file if it exists
	if update != "" {
		updateNeighborhoods, err := p.parseFile(update)
		if err != nil {
			return nil, fmt.Errorf("error parsing update file: %w", err)
		}

		for id, neighborhood := range updateNeighborhoods {
			baseNeighborhoods[id] = neighborhood
		}
	}

	return baseNeighborhoods, nil
}

func (p *BairroParser) parseFile(basePath string) (map[int]models.Neighborhood, error) {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", basePath, err)
	}

	var neighborhoods []models.Neighborhood

	for _, entry := range entries {
		if !strings.HasPrefix(strings.ToUpper(entry.Name()), "LOG_BAIRRO") &&
			!strings.HasPrefix(strings.ToUpper(entry.Name()), "DELTA_LOG_BAIRRO") {
			continue
		}

		filepath := path.Join(basePath, entry.Name())
		file, err := os.Open(filepath)
		if err != nil {
			return nil, fmt.Errorf("error opening file %s: %w", filepath, err)
		}
		defer file.Close()

		dec := transform.NewReader(file, charmap.Windows1252.NewDecoder())

		reader := csv.NewReader(dec)
		reader.Comma = '@'
		reader.FieldsPerRecord = -1 // Allow variable number of fields
		reader.TrimLeadingSpace = true

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("error reading file %s: %w", filepath, err)
			}

			id, err := strconv.Atoi(strings.TrimSpace(record[0]))
			if err != nil {
				return nil, fmt.Errorf("error parsing ID: %w", err)
			}

			// Parse each record into a Neighborhood object
			neighborhood := models.Neighborhood{
				ID:   id,
				Name: strings.TrimSpace(record[3]),
			}
			neighborhoods = append(neighborhoods, neighborhood)
		}
	}

	return models.NeighborhoodMap(neighborhoods), nil
}
