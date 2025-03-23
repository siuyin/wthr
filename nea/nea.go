package nea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type latLng struct {
	Latitude  json.Number `json:"latitude"`
	Longitude json.Number `json:"longitude"`
}
type area struct {
	Name          string `json:"name"`
	LabelLocation latLng `json:"label_location"`
}
type item struct {
	UpdateTimestamp string `json:"update_timestamp"`
	Timestamp       string `json:"timestamp"`
	ValidPeriod     struct {
		Start string `json:"start"`
		End   string `json:"end"`
		Text  string `json:"text"`
	} `json:"valid_period"`
	Forecasts []struct {
		Area     string `json:"area"`
		Forecast string `json:"forecast"`
	}
	ErrorMsg string `json:"errorMsg"`
}
type data struct {
	AreaMetadata []area `json:"area_metadata"`
	Items        []item `json:"items"`
}
type Msg struct {
	Code json.Number `json:"code"`
	Data data        `json:"data"`
}

func (m Msg) String() string {
	b := &bytes.Buffer{}
	for i, d := range m.Data.AreaMetadata {
		fmt.Fprintf(b, "%d: %v\n", i, d)
	}

	fmt.Fprintf(b, "\nForecasts for: %s\n", m.Data.Items[0].ValidPeriod.Text)

	for i, d := range m.Data.Items[0].Forecasts {
		fmt.Fprintf(b, "%d: %v\n", i, d)
	}
	return b.String()
}

// Forecast2Hr provides a weather forecast covering a 2 hour interval.
func Forecast2Hr() Msg {
	const url = "https://api-open.data.gov.sg/v2/real-time/api/two-hr-forecast"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	return Decode(res.Body)
}

// Decode parses the message retured from an NEA API call.
func Decode(r io.Reader) Msg {
	var msg Msg
	dec := json.NewDecoder(r)
	if err := dec.Decode(&msg); err != nil {
		log.Println(err)
	}

	return msg
}

type AreaCoords struct {
	Name string
	Long float64
	Lat  float64
}

// AreaLongLat returns a list of area names, and their longitudes and latitudes.
func AreaLongLat(m Msg) []AreaCoords {
	a := []AreaCoords{}
	for _, v := range m.Data.AreaMetadata {
		var e AreaCoords
		var err error
		e.Name = v.Name

		e.Long, err = v.LabelLocation.Longitude.Float64()
		if err != nil {
			log.Fatalf("bad data: %s should be a float", v.LabelLocation.Longitude)
		}

		e.Lat, err = v.LabelLocation.Latitude.Float64()
		if err != nil {
			log.Fatalf("bad data: %s should be a float", v.LabelLocation.Longitude)
		}

		a = append(a, e)
	}
	return a
}
