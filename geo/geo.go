// Package go provides geo-spatial functions.
package geo

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/siuyin/dflt"
	"github.com/tidwall/buntdb"
)

// Coord represents the latitude and longitude of a place.
type Coord struct {
	Name string
	Lat  float64
	Long float64
}

var db *buntdb.DB

func init() {
	var err error
	f := dflt.EnvString("DB_FILE", "/tmp/geo.db")
	db, err = buntdb.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(`Opened geo-spatial DB_FILE="%s"`, f)
}

// Load loads coordinates into the internal geospatial database.
func Load(coords []Coord) error {
	for _, c := range coords {
		val := fmt.Sprintf("[%f %f]", c.Long, c.Lat)
		set(db, c.Name, val)
	}

	if err := db.CreateSpatialIndex("geo", "*", buntdb.IndexRect); err != nil {
		return err
	}

	return nil
}

// Close closes the internal geospatial databse.
func Close() {
	db.Close()
}

func Nearest(k int, c Coord, dist float64) []Coord {
	coords := []Coord{}
	db.View(func(tx *buntdb.Tx) error {
		tx.Nearby("geo", lngLat(c), func(k, v string, dist float64) bool {
			coords = append(coords, latLng(k, v))
			return true
		})
		return nil
	})

	if k == 0 {
		return coords
	}
	if k > len(coords) {
		return coords
	}
	return coords[:k]
}
func lngLat(c Coord) string {
	return fmt.Sprintf("[%f %f]", c.Long, c.Lat)
}
func latLng(name, lngLatStr string) Coord {
	s := strings.ReplaceAll(lngLatStr, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	a := strings.Split(s, " ")
	lng, err := strconv.ParseFloat(a[0], 64)
	if err != nil {
		log.Fatalf("not a float: %s: %v", a[0], err)
	}
	lat, err := strconv.ParseFloat(a[1], 64)
	if err != nil {
		log.Fatalf("not a float: %s: %v", a[0], err)
	}
	return Coord{name, lat, lng}
}

func set(db *buntdb.DB, k, v string) {
	if err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(k, v, nil)
		return err
	}); err != nil {
		log.Println(err)
	}
}

func get(db *buntdb.DB, k string) string {
	var (
		val string
		err error
	)

	db.View(func(tx *buntdb.Tx) error {
		val, err = tx.Get(k)
		return err
	})
	if err != nil {
		log.Printf("%s: %v\n", k, err)
		return ""
	}

	return val
}

func list(db *buntdb.DB) {
	if err := db.View(func(tx *buntdb.Tx) error {
		return tx.Ascend("", func(k, v string) bool {
			fmt.Println(k, v)
			return true
		})
	}); err != nil {
		log.Println(err)
	}
}
