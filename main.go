package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/siuyin/dflt"
	"github.com/siuyin/wthr/geo"
	"github.com/siuyin/wthr/nea"
	"github.com/siuyin/wthr/nea/daily"
	"github.com/siuyin/wthr/public"
)

func init() {
	msg := nea.Forecast2Hr()
	coords := nea.Coords(msg)
	geo.Load(coords)
}

func main() {
	http.Handle("/", http.FileServer(http.FS(public.Content)))
	http.HandleFunc("/fc", forecastHandler)
	http.HandleFunc("/nfc", neighbourhoodForecastHandler)
	http.HandleFunc("/daily", dailyForecastHandler)

	port := dflt.EnvString("PORT", "8080")
	log.Printf("starting webserver on PORT=%s\n", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func forecastHandler(w http.ResponseWriter, r *http.Request) {
	msg := nea.Forecast2Hr()
	fc := nea.AreaForecasts(msg)
	fmt.Fprintf(w, "<html><h2>Singapore Weather</h2>")
	fmt.Fprintf(w, "<p>%s</p>", nea.ForecastPeriod(msg))
	for i, f := range fc {
		fmt.Fprintf(w, "%d. %s: %s<br>", i+1, f.Area, f.Forecast)
	}
	fmt.Fprintf(w, "</html>")
}

func neighbourhoodForecastHandler(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(r.FormValue("lat"), 64)
	if err != nil {
		log.Println(err)
	}

	lng, err := strconv.ParseFloat(r.FormValue("lng"), 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("<html><h2>Local Forecasts (%.4f,%.4f)</h2>", lat, lng))

	fc := nea.NeighbourhoodForecast(nea.Forecast2Hr(), lat, lng)
	for i, f := range fc {
		fmt.Fprintf(w, "%d. %s: %s<br>", i+1, f.Area, f.Forecast)
	}
	fmt.Fprintf(w, "</html>")
}

func dailyForecastHandler(w http.ResponseWriter, r *http.Request) {
	msg := daily.Forecast()
	cf := daily.CurrentForecast(msg)
	fmt.Fprintf(w, "<h2>Daily Forecasts</h2>")
	g := msg.Data.Records[0].General
	fmt.Fprintf(w, "<p>min: %d°C, max: %d°C, %s</p>", g.Temperature.Low, g.Temperature.High, g.Forecast.Text)
	for _, f := range cf {
		fmt.Fprintf(w, "<p>%s<br>East: %s<br>West: %s<br>Central: %s<br>North: %s<br>South: %s</p>", f.TimePeriod.Text,
			f.Region.East.Text,
			f.Region.West.Text,
			f.Region.Central.Text,
			f.Region.North.Text,
			f.Region.South.Text,
		)
	}
}
