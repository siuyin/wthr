package nea

import (
	"fmt"
	"os"
	"strings"
	"testing"
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

func TestAreaLongLat(t *testing.T) {
	msg := sampleMsg(t)
	a := AreaLongLat(msg)
	if n := len(a); n == 0 {
		t.Error("bad len")
	}

	if s := fmt.Sprintf("%s, [%.3f %.3f]", a[0].Name, a[0].Long, a[0].Lat); s != "Ang Mo Kio, [103.839 1.375]" {
		t.Errorf("bad string compare: got %s", s)
	}
}
