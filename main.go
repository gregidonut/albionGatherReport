package main

import (
	"github.com/gregidonut/albionGatherReport.git/gatherReport"
	"log"
)

func main() {
	payload := gatherReport.NewGatherReportRow(
		"[Screenshot_2023-09-18_11-46-48.png](https://lh3.googleusercontent.com/drive-viewer/AITFw-yDxE2y_Z_IY4FXIMsAD3rXX-1PDqszeFnrEMaRTKPtfRKtOz4x6D7-wWdlMZnJjuXEYqVPDFyA70F-KYrRXnbPQQWkUw=w1918-h1096)",
		0,
		1,
		0,
		92_000,
	)

	if err := payload.WriteReportRow(); err != nil {
		log.Fatal(err)
	}
}
