package main

import (
	"log"
	"net/http"

	"github.com/siuyin/dflt"
	"github.com/siuyin/wthr/public"
)

func init() {
	//nea2HrForecast()
}

func main() {
	http.Handle("/", http.FileServer(http.FS(public.Content)))

	port := dflt.EnvString("PORT", "8080")
	log.Printf("starting webserver on PORT=%s\n", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
