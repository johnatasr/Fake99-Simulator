package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Route struct {
	ID        string     `json:"routeID"`
	ClientID  string     `json:"clientID"`
	Positions []Position `json:"positions"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeID"`
	ClientID string    `json:"clientID"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {

	if r.ID == "" {
		return errors.New("ID not informed")
	}

	file, err := os.Open("destination/" + r.ID + ".txt")

	if err != nil {
		return nil
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)

		if err != nil {
			return errors.New("Latitude not defined")
		}

		lon, err := strconv.ParseFloat(data[1], 64)

		if err != nil {
			return errors.New("Longitude not defined")
		}

		r.Positions = append(r.Positions, Position{
			Lat: lat,
			Lon: lon,
		})
	}
	return nil
}

func (r *Route) ExportJson() ([]string, error) {
	var route PartialRoutePosition
	var result []string

	total := len(r.Positions)

	for i, j := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{j.Lat, j.Lon}
		route.Finished = false

		if total-1 == i {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)

		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}
