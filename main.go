package main

import (
	"github.com/gregidonut/albionGatherReport.git/gatherReport"
	"log"
	"time"
)

func main() {
	payload := &gatherReport.GatherReport{
		Date: time.Now(),
		Cost: gatherReport.Cost{
			JourneyBack: 10,
			PorkPie:     20,
			Repair:      5,
		},
		Profit:  100,
		Revenue: 200,
	}

	if err := payload.WriteReport(); err != nil {
		log.Fatal(err)
	}
}
