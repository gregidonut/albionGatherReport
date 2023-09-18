package main

import (
	"github.com/gregidonut/albionGatherReport.git/gatherReport"
	"log"
)

func main() {
	payload := gatherReport.NewGatherReportRow(
		"[Screenshot_2023-09-18_11-31-29.png](https://lh3.googleusercontent.com/drive-viewer/AITFw-w-Jb8lcvOk0cHH7LOtOoYb3CMw8t7g_bmusxdazIlElLuYXgY-sODDykfSmmjWZFu1lYGfoMD41rprAQRsMq04Tmci=w1918-h1072)",
		490_000,
		6,
		98_000,
		1_2000_000,
	)

	if err := payload.WriteReportRow(); err != nil {
		log.Fatal(err)
	}
}
