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

type LocationParser struct{}

func NewLocationParser() *LocationParser {
	return &LocationParser{}
}

func (p *LocationParser) Parse(base, update string) (map[int]models.Location, error) {
	// Read base file
	baseLocations, err := p.parseFile(base)
	if err != nil {
		return nil, fmt.Errorf("error parsing base file: %w", err)
	}

	// Read update file if it exists
	if update != "" {
		updateLocations, err := p.parseFile(update)
		if err != nil {
			return nil, fmt.Errorf("error parsing update file: %w", err)
		}

		for id, location := range updateLocations {
			baseLocations[id] = location
		}
	}

	return baseLocations, nil
}

func (p *LocationParser) parseFile(basePath string) (map[int]models.Location, error) {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", basePath, err)
	}

	var locations []models.Location

	for _, entry := range entries {
		if !strings.HasPrefix(strings.ToUpper(entry.Name()), "LOG_LOCALIDADE") &&
			!strings.HasPrefix(strings.ToUpper(entry.Name()), "DELTA_LOG_LOCALIDADE") {
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

			zipCode := 0
			if record[3] != "" {
				zipCode, err = strconv.Atoi(strings.TrimSpace(record[3]))
				if err != nil {
					return nil, fmt.Errorf("error parsing zip code: %w", err)
				}
			}

			situation, err := strconv.Atoi(strings.TrimSpace(record[4]))
			if err != nil {
				return nil, fmt.Errorf("error parsing situation: %w", err)
			}

			subordinateLocationID := 0
			if record[6] != "" {
				subordinateLocationID, err = strconv.Atoi(strings.TrimSpace(record[6]))
				if err != nil {
					return nil, fmt.Errorf("error parsing subordinate location ID: %w", err)
				}
			}

			// Parse each record into a Neighborhood object
			location := models.Location{
				ID:                    id,
				State:                 strings.TrimSpace(record[1]),
				Name:                  strings.TrimSpace(record[2]),
				ZipCode:               zipCode,
				Situation:             models.LocationSituation(situation),
				Type:                  models.LocationType(record[5]),
				SubordinateLocationID: subordinateLocationID,
				IBGECode:              record[8],
			}

			locations = append(locations, location)
		}
	}

	return models.LocationMap(locations), nil
}
