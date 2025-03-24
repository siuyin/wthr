package geo

import "testing"

var testCoords = []Coord{
	Coord{"PlaceA", 1.033, 103.233},
	Coord{"PlaceB", 1.234, 103.456},
}

func TestInit(t *testing.T) {
	if db == nil {
		t.Error("did not initialize db")
	}
}

func TestLoad(t *testing.T) {
	err := Load(testCoords)
	if err != nil {
		t.Errorf("failed to load coords: %v", err)
	}

	coord := get(db, "PlaceB")
	if coord != "[103.456000 1.234000]" {
		t.Errorf("incorrect get: %s", coord)
	}
}

func TestNearestK(t *testing.T) {
	Load(testCoords)
	coords := Nearest(1, Coord{"myPlace", 1.234, 103.456}, 1.0)
	if len(coords) != 1 {
		t.Error("bad num of coords")
	}

	if nm := coords[0].Name; nm != "PlaceB" {
		t.Errorf("bad sort: got %s", nm)
	}

	coords = Nearest(0, Coord{"yourPlace", 1.012, 103.234}, 1.0)
	if len(coords) != 2 {
		t.Error("bad num of coords")
	}

	if nm := coords[0].Name; nm != "PlaceA" {
		t.Errorf("bad sort: got %s", nm)
	}
}
