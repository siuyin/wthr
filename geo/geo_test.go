package geo

import (
	"log"
	"os"
	"testing"
)

func TestGeo(t *testing.T) {
	var testCoords = []Coord{
		Coord{"PlaceA", 1.033, 103.233},
		Coord{"PlaceB", 1.234, 103.456},
	}

	f, err := os.CreateTemp("", "geotest-*")
	if err != nil {
		t.Error(err)
	}

	t.Setenv("DB_FILE", f.Name())
	t.Cleanup(func() {
		os.Setenv("DB_FILE", "")
		if err := os.Remove(f.Name()); err != nil {
			log.Println(err)
		}
	})

	t.Run("openDB", func(t *testing.T) {
		openDB()
		if db == nil {
			t.Error("did not initialize db")
		}
	})

	t.Run("Load", func(t *testing.T) {
		err := Load(testCoords)
		if err != nil {
			t.Errorf("failed to load coords: %v", err)
		}

		coord := get("PlaceB")
		if coord != "[103.456000 1.234000]" {
			t.Errorf("incorrect get: %s", coord)
		}
	})

	t.Run("NearestK", func(t *testing.T) {
		Load(testCoords)
		coords := Nearest(1, Coord{"myPlace", 1.234, 103.456}, 1.0)
		if len(coords) != 1 {
			t.Error("bad num of coords")
		}

		if nm := coords[0].Name; nm != "PlaceB" {
			t.Errorf("bad sort: got %s", nm)
		}

		coords = Nearest(0, Coord{"yourPlace", 1.012, 103.234}, 1.0)
		if n := len(coords); n != 2 {
			t.Errorf("bad num of coords, got: %d", n)
		}

		if nm := coords[0].Name; nm != "PlaceA" {
			t.Errorf("bad sort: got %s", nm)
		}
	})
}
