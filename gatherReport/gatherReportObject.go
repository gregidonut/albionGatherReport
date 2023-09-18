package gatherReport

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type Cost struct {
	JourneyBack int
	PorkPie     int
	Repair      int
	total       int
}

func newCost(journeyBack, porkPie, repair int) *Cost {
	payload := &Cost{
		JourneyBack: journeyBack,
		PorkPie:     porkPie,
		Repair:      repair,
	}
	payload.total = (porkPie * 5_000) + journeyBack + repair
	return payload

}

type GatherReport struct {
	Date    time.Time
	Cost    *Cost
	Profit  int
	Revenue int
	URL     string
}

func NewGatherReportRow(url string, journeyback, porkpie, repair, revenue int) *GatherReport {
	payload := &GatherReport{
		Date: time.Now(),
		Cost: newCost(
			journeyback,
			porkpie,
			repair,
		),
		Revenue: revenue,
		URL:     url,
	}

	payload.Profit = revenue - payload.Cost.total
	return payload
}

func (gr *GatherReport) WriteReportRow() error {
	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		return fmt.Errorf("%v:%v", parsingTemplateErr, err)
	}
	// Open the file in append mode, creating it if it doesn't exist.
	file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("%v:%v", creatingReadMeErr, err)
	}
	defer file.Close()
	if err = tmpl.Execute(file, []*GatherReport{gr}); err != nil {
		return fmt.Errorf("%v:%v", writingReadMeErr, err)
	}
	fmt.Println("README.md updated successfully.")
	return nil
}

const readmeTemplate = `
{{range .}}| {{.Date.Format "2006-01-02"}} | {{.Cost.JourneyBack}} | {{.Cost.PorkPie}} | {{.Cost.Repair}} | {{.Profit}} | {{.Revenue}} | {{.URL}} |{{end}}`
