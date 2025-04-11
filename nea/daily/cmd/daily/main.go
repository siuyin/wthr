package main

import (
	"fmt"

	"github.com/siuyin/wthr/nea/daily"
)

func main() {
	msg := daily.Forecast()
	fmt.Printf("Forecast for %s: %s\n", msg.Data.Records[0].General.ValidPeriod.Text,
		msg.Data.Records[0].General.Forecast.Text)
	fmt.Printf(daily.Current(msg))
}
