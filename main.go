package main

import (
	"github.com/gregidonut/albionGatherReport.git/gatherReport"
	"log"
)

func main() {
	payload := gatherReport.NewGatherReportRow(
		"[Screenshot_2023-09-18_17-00-46.png](https://lh3.googleusercontent.com/drive-viewer/AITFw-x_ksNsuqiHAiRMPiQLGiPL32ua0Ib-Jq0Fc4PNL5-Tj_4W1eWs0rbv0OJ0r3ABST7J66teKfHO2-tVC747Mc3bhqx3=w1081-h1072)",
		map[string]int{
			//"journeyBack": 2_000,
			"porkPie": 2,
			//"repair":      4_000,
			"revenue": 585_000,
		},
	)

	if err := payload.WriteReportRow(); err != nil {
		log.Fatal(err)
	}
}
