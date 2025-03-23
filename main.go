package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
	"github.com/siuyin/wthr/public"
)

func init() {
	nea2HrForecast()
}

func main() {
	http.Handle("/", http.FileServer(http.FS(public.Content)))

	port := dflt.EnvString("PORT", "8080")
	log.Printf("starting webserver on PORT=%s\n", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
type neaMsg struct {
	Code json.Number `json:"code"`
	Data data        `json:"data"`
}

func nea2HrForecast() neaMsg {
	const url = "https://api-open.data.gov.sg/v2/real-time/api/two-hr-forecast"
	var msg neaMsg
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&msg); err != nil {
		log.Println(err)
	}

	for i, d := range msg.Data.AreaMetadata {
		fmt.Printf("%d: %v\n", i, d)
	}

	fmt.Printf("\nForecasts for: %s\n", msg.Data.Items[0].ValidPeriod.Text)

	for i, d := range msg.Data.Items[0].Forecasts {
		fmt.Printf("%d: %v\n", i, d)
	}

	return msg
}
