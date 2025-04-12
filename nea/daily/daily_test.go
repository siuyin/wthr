package daily

import (
	"os"
	"testing"
	"time"
)

func sampleMsg(t *testing.T) Msg {
	f, err := os.Open("./testdata/daily.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	return Decode(f)
}

func TestDailyForecast(t *testing.T) {
	msg := sampleMsg(t)

	t.Run("getMsg", func(t *testing.T) {
		if msg.Code != 0 {
			t.Error("bad code")
		}
		if len(msg.Data.Records) != 1 {
			t.Error("bad records len")
		}
		r := msg.Data.Records[0]
		if r.General.Temperature.High != 33 {
			t.Error("bad temp")
		}
		if r.General.Forecast.Text != "Thundery Showers" {
			t.Error("bad forecast")
		}
		p := r.Periods
		if len(p) != 4 {
			t.Error("bad periods len")
		}
		l := p[3]
		if l.TimePeriod.End != "2025-04-13T00:00:00+08:00" {
			t.Error("bad end time")
		}
		if l.Region.North.Text != "Cloudy" {
			t.Error("bad region forecast")
		}
	})
}

func TestTimePeriod(t *testing.T) {
	dat := []struct {
		tm time.Time
		p  int
	}{
		{time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), 0},
		{time.Date(2000, 1, 1, 6, 0, 0, 0, time.UTC), 1},
		{time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC), 2},
		{time.Date(2000, 1, 1, 18, 0, 0, 0, time.UTC), 3},

		{time.Date(2000, 1, 1, 5, 59, 59, 0, time.UTC), 0},
		{time.Date(2000, 1, 1, 11, 59, 59, 0, time.UTC), 1},
		{time.Date(2000, 1, 1, 17, 59, 59, 0, time.UTC), 2},
		{time.Date(2000, 1, 1, 23, 59, 59, 0, time.UTC), 3},
	}
	for _, d := range dat {
		if p := timePeriod(d.tm); p != d.p {
			t.Errorf("%s should be in period %d, got %d", d.tm, d.p, p)
		}
	}

}
