package main

import (
	"flag"
	"github.com/ddeep2007/KopCapital/pkg/reporting"
	"github.com/ddeep2007/KopCapital/pkg/stock"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	avAPIKey = "SPLMZGZRF1JE1U1G"
)

var (
	ticker = flag.String("ticker", "SNOW", "ticker symbol to list data for")
	report = flag.String("report", "Saas", "the report type which needs to be generated")
)

func main() {
	flag.Parse()

	logger := log.New()
	logger.Out = os.Stdout

	av := stock.InitAlphaVantage(logger, avAPIKey)
	r := reporting.Generator{Logger: logger, Av: av}

	err := r.GenerateTickerReport(*ticker, *report)
	if err != nil {
		return
	}
	return
}
