package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"encoding/csv"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/opts"
)

func main() {
	names := make([]opts.BarData, 0)
	points := make([]opts.BarData, 0)
	names, points = readCsv()  
	barg(names, points)
	gauge()
}

func readCsv() ([]opts.BarData, []opts.BarData ){
	csvfile, err := os.Open("Data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	names := make([]opts.BarData, 0)
	points := make([]opts.BarData, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, opts.BarData{Value: record[0]})
		points = append(points, opts.BarData{Value: record[1]})
		fmt.Printf("%s %s\n", record[0], record[1])
	}
	return names, points
}


func barg(names []opts.BarData, points []opts.BarData) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Stats",
			Subtitle: "Details",
		}),
	)
	f, err := os.Create("barg.html")
	if err != nil {
		panic(err)
	}
	bar.SetXAxis(names).AddSeries("Points", points)
	bar.Render(f)
}

func gauge() {
	gauge := charts.NewGauge()
    gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Age",
		}),
	)
	age := make([]opts.GaugeData, 0)
	age = append(age, opts.GaugeData{Value: 23})
    gauge.AddSeries("Kshitiz", age)
    f, err := os.Create("gauge.html")
	if err != nil {
		panic(err)
	}
	gauge.Render(f)
}