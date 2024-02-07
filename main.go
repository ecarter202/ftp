package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

const KG_IN_A_POUND = 0.4535924

func main() {
	lbs := os.Args[1]

	if lbs == "" {
		log.Fatal("weight required")
	}

	lbsFloat, err := strconv.ParseFloat(lbs, 64)
	if err != nil {
		log.Fatalf("invalid weight [ERR: %s]", err)
	}

	kilos := lbsFloat * KG_IN_A_POUND

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Watt/kg", "Watts")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for i := 0.5; i <= 7; i += 0.5 {
		watts := kilos * i
		tbl.AddRow(i, fmt.Sprintf("%.2f", watts))
	}

	tbl.Print()
}
