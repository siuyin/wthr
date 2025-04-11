package daily

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Msg struct {
	Code     int    `json:"code"`
	Data     data   `json:"data"`
	ErrorMsg string `json:"errorMsg"`
}
type data struct {
	Records []record `json:"records"`
}
type record struct {
	Date             string   `json:"date"`
	UpdatedTimestamp string   `json:"updatedTimestamp"`
	General          general  `json:"general"`
	Periods          []period `json:"periods"`
	Timestamp        string   `json:"timestamp"`
}
type general struct {
	Temperature struct {
		Low  int    `json:"low"`
		High int    `json:"high"`
		Unit string `json:"unit"`
	} `json:"temperature"`
	RelativeHumidity struct {
		Low  int    `json:"low"`
		High int    `json:"high"`
		Unit string `json:"unit"`
	} `json:"relativeHumidity"`
	Forecast struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"forecast"`
	ValidPeriod struct {
		Start string `json:"start"`
		End   string `json:"end"`
		Text  string `json:"text"`
	} `json:"validPeriod"`
	Wind struct {
		Speed struct {
			Low  int `json:"low"`
			High int `json:"high"`
		} `json:"speed"`
		Direction string `json:"direction"`
	} `json:"wind"`
}
type period struct {
	TimePeriod struct {
		Start string `json:"start"`
		End   string `json:"end"`
		Text  string `json:"text"`
	} `json:"timePeriod"`
	Regions struct {
		West struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"west"`
		East struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"east"`
		Central struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"central"`
		South struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"south"`
		North struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"north"`
	}
}

func Decode(r io.Reader) Msg {
	var msg Msg
	dec := json.NewDecoder(r)
	if err := dec.Decode(&msg); err != nil {
		log.Println(err)
	}

	return msg
}

func Forecast() Msg {
	const url = "https://api-open.data.gov.sg/v2/real-time/api/twenty-four-hr-forecast"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	return Decode(res.Body)
}

func Current(msg Msg) string {
	p := timePeriod(time.Now())
	r := msg.Data.Records[0].Periods[p].Regions
	fc0 := fmt.Sprintf("West: %s\nEast: %s\nCentral: %s\nSouth: %s\nNorth: %s\n",
		r.West.Text, r.East.Text, r.Central.Text, r.South.Text, r.North.Text)
	if p == 3 {
		return fc0
	}

	r = msg.Data.Records[0].Periods[p+1].Regions
	fc1 := fmt.Sprintf("\nLater:\nWest: %s\nEast: %s\nCentral: %s\nSouth: %s\nNorth: %s\n",
		r.West.Text, r.East.Text, r.Central.Text, r.South.Text, r.North.Text)
	return fc0 + fc1
}

func timePeriod(tm time.Time) int {
	h := tm.Hour()
	switch {
	case h >= 0 && h < 6:
		return 0
	case h >= 6 && h < 12:
		return 1
	case h >= 12 && h < 18:
		return 2
	case h >= 18 && h < 24:
		return 3
	default:
		return -1
	}
}
