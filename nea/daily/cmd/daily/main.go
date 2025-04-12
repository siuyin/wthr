package main

import (
	"fmt"

	"github.com/siuyin/wthr/nea/daily"
)

func main() {
	msg := daily.Forecast()
	fmt.Printf("Forecast for %s: %s\n", msg.Data.Records[0].General.ValidPeriod.Text,
		msg.Data.Records[0].General.Forecast.Text)
	// fmt.Printf(daily.Current(msg))

	cn := daily.CurrentForecast(msg)
	for i, p := range cn {
		if i == 1 {
			fmt.Printf("\nLater:\n")
		}
		fmt.Printf("Period: %s\n", p.TimePeriod.Text)
		fmt.Printf("West: %s\nEast: %s\nCentral: %s\nSouth: %s\nNorth: %s\n",
			p.Region.West.Text, p.Region.East.Text, p.Region.Central.Text, p.Region.South.Text, p.Region.North.Text)
	}
}
