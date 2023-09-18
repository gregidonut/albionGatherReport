package gatherReport

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

const (
	currentPorkPieMarketPrice = 5_000
)

type Cost struct {
	JourneyBack int
	PorkPie     int
	Repair      int
	total       int
}

func newCost(kwargs map[string]int) *Cost {
	var payLoad = new(Cost)

	journeyBack, ok := kwargs["journeyBack"]
	if ok {
		payLoad.JourneyBack = journeyBack
	}
	porkPie, ok := kwargs["porkPie"]
	if ok {
		payLoad.PorkPie = porkPie
	}
	repair, ok := kwargs["repair"]
	if ok {
		payLoad.Repair = repair
	}

	payLoad.total = (porkPie * currentPorkPieMarketPrice) + journeyBack + repair
	return payLoad
}

type GatherReport struct {
	Date    time.Time
	Cost    *Cost
	Profit  int
	Revenue int
	URL     string
}

func NewGatherReportRow(url string, kwargs map[string]int) *GatherReport {
	payLoad := &GatherReport{
		Date: time.Now(),
		Cost: newCost(kwargs),
		URL:  url,
	}

	revenue, ok := kwargs["revenue"]
	if ok {
		payLoad.Revenue = revenue
	}
	payLoad.Profit = payLoad.Revenue - payLoad.Cost.total

	return payLoad
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
