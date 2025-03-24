package nea

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/siuyin/wthr/geo"
)

func TestDecode(t *testing.T) {
	msg := sampleMsg(t)
	if n := len(msg.Data.AreaMetadata); n != 47 {
		t.Errorf("bad parse: %d", n)
	}
}
func loadSample(t *testing.T) *os.File {
	f, err := os.Open("../output.sample.json")
	if err != nil {
		t.Error(err)
	}
	return f
}
func sampleMsg(t *testing.T) Msg {
	f := loadSample(t)
	defer f.Close()

	return Decode(f)
}

func TestMsgToString(t *testing.T) {
	msg := sampleMsg(t)
	if s := msg.String(); !strings.Contains(s, "Yishun Cloudy") {
		t.Errorf(`%s should have "Yishun Cloudy"`, s)
	}
}

func TestCoords(t *testing.T) {
	msg := sampleMsg(t)
	a := Coords(msg)
	if n := len(a); n == 0 {
		t.Error("bad len")
	}

	if s := fmt.Sprintf("%s, [%.3f %.3f]", a[0].Name, a[0].Long, a[0].Lat); s != "Ang Mo Kio, [103.839 1.375]" {
		t.Errorf("bad string compare: got %s", s)
	}
}

func TestForecast(t *testing.T) {
	msg := sampleMsg(t)
	af := AreaForecasts(msg)
	if n := len(af); n != 47 {
		t.Errorf("bad len, got: %d", n)
	}
	if af[46].Area != "Yishun" {
		t.Error("bad area")
	}
	if af[46].Forecast != "Cloudy" {
		t.Error("bad forecast")
	}
}

func TestNeighbourhoodForecast(t *testing.T) {
	f, err := os.CreateTemp("", "neatest-")
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
	msg := sampleMsg(t)
	coords := Coords(msg)
	geo.Load(coords)

	t.Run("numLocs", func(t *testing.T) {
		fc := NeighbourhoodForecast(msg, 1.023, 103.15)
		if n := len(fc); n != 3 {
			t.Errorf("bad len, got: %d", n)
		}

		area := fc[0].Area
		forecast := fc[0].Forecast
		if area != "Tuas" && forecast != "Cloudy" {
			t.Errorf("bad forecast, got: %s %s", area, forecast)
		}
	})

}
