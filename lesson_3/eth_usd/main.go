package main

import (
	"eth_usd/csv_process"
	"eth_usd/plotter"
	"fmt"
	"gonum.org/v1/plot/vg"
	"os"
)

func main() {
	pricePairs, err := csv_process.LoadDataFrom("prices.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading prices from CSV file: %v\n", err)
		os.Exit(1)
	}

	pricesPlot, err := plotter.GeneratePlotFor(pricePairs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plot: %v\n", err)
		os.Exit(1)
	}

	if err := pricesPlot.Save(15*vg.Inch, 4*vg.Inch, "ethereum_prices.png"); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving plot: %v\n", err)
		os.Exit(1)
	}
}
