package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
	"github.com/siuyin/wthr/geo"
	"github.com/siuyin/wthr/nea"
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

	port := dflt.EnvString("PORT", "8080")
	log.Printf("starting webserver on PORT=%s\n", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func forecastHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>")
	fc := nea.AreaForecasts(nea.Forecast2Hr())
	for i, f := range fc {
		fmt.Fprintf(w, "%d. %s: %s<br>", i+1, f.Area, f.Forecast)
	}
	fmt.Fprintf(w, "</html>")
}
