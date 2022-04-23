// Binary geojson converts a set of CSV files into a GeoJSON feature
// collection.
package main

// SPDX-License-Identifier: MIT

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type Region struct {
	Region, LocationName string
	Lat, Lon             float64
	OsmID, CountryTLD    string
}

func main() {
	fc := geojson.NewFeatureCollection()

	for _, file := range os.Args[1:] {
		data, err := parseCSV(file)
		if err != nil {
			panic(err)
		}

		provider := filepath.Base(filepath.Dir(file))

		for _, region := range data {
			point := geojson.NewFeature(orb.Point{region.Lon, region.Lat})
			point.Properties["provider"] = provider
			point.Properties["region"] = region.Region
			point.Properties["location_name"] = region.LocationName
			point.Properties["osm_id"] = region.OsmID
			point.Properties["country"] = region.CountryTLD
			fc.Append(point)
		}
	}

	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(rawJSON)
}

func parseCSV(file string) ([]Region, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip header
	_, err = r.Read()
	if err != nil {
		return nil, err
	}
	data := []Region{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lat, _ := strconv.ParseFloat(record[2], 64)
		lon, _ := strconv.ParseFloat(record[3], 64)
		data = append(data, Region{record[0], record[1], lat, lon, record[4], record[5]})
	}
	return data, nil
}
