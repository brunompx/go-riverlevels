package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/brunompx/go-riverlevels/templates/pages"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCntLine = 7
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func HandleLineChart(w http.ResponseWriter, r *http.Request) {

	//hart := lineSmoothArea()
	//chart := lineSymbolsb()
	chart := lineForecat()

	pages.Linechart(chart).Render(r.Context(), w)
}

func lineSymbolsb() *charts.Line {

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "symbol options",
			Subtitle: "tooltip with 'axis' trigger",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}),
	)

	// Put data into instance
	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItemsb()).
		AddSeries("Category B", generateLineItemsb()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 15, Symbol: "diamond"},
		))

	return line
}

func generateLineItemsb() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

// smooth
func lineSmoothArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "smooth area"}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.2,
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: opts.Bool(true),
			}),
		)
	return line
}

// smooth
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

// func lineForecat(forecats []model.Forecast) *charts.Line {
func lineForecat() *charts.Line {

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "symbol options",
			Subtitle: "tooltip with 'axis' trigger",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}),
	)

	//first := []int{1, 2, 3, 4}
	sals1 := []opts.LineData{{Value: 2}, {Value: 1}, {Value: 3}, {Value: 3}, {Value: 4}, {Value: 5}}
	sals2 := []opts.LineData{{Value: 2}, {Value: 1}, {Value: 3}, {Value: 3}, {Value: 4}, {Value: 5}, {Value: 7}, {Value: 7}, {Value: 8}, {Value: 10}, {Value: 10}, {Value: 11}}
	sals3 := []opts.LineData{{Value: 2}, {Value: 1}, {Value: 3}, {Value: 3}, {Value: 4}, {Value: 5}, {Value: 5}, {Value: 6}, {Value: 7}, {Value: 8}, {Value: 8}, {Value: 9}}
	sals4 := []opts.LineData{{Value: 2}, {Value: 1}, {Value: 3}, {Value: 3}, {Value: 4}, {Value: 5}, {Value: 4}, {Value: 5}, {Value: 6}, {Value: 7}, {Value: 6}, {Value: 5}}

	//xaxis := opts.XAxis{Data: []string{"a", "b", "c", "f", "h", "j", "e", "r", "t", "y", "hu", "ji"}}
	t1 := stringT("2024-02-13")
	t2 := stringT("2024-02-14")
	t3 := stringT("2024-02-15")
	t4 := stringT("2024-02-16")
	t5 := stringT("2024-02-17")
	t6 := stringT("2024-02-18")
	t7 := stringT("2024-02-19")
	t8 := stringT("2024-02-20")
	t9 := stringT("2024-02-21")
	t10 := stringT("2024-02-22")
	t11 := stringT("2024-02-23")
	t12 := stringT("2024-02-24")

	//line.SetGlobalOptions(charts.WithXAxisOpts(
	//	opts.XAxis{Type: "time"},
	//))

	// Put data into instance
	//line.SetXAxis([]string{"a", "b", "c", "f", "h", "j", "e", "r", "t", "y", "hu", "ji"})
	//line.SetXAxis([]time.Time{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12})
	line.SetXAxis([]string{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12})

	line.AddSeries("Category A", sals4)
	line.AddSeries("Category b", sals2)
	line.AddSeries("Category c", sals3)
	line.AddSeries("Category Romeo", sals1)

	line.SetSeriesOptions(charts.WithLineChartOpts(
		opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 15, Symbol: "diamond"},
	))

	return line
}

func stringTw(stringValue string) time.Time {
	parsed, err := time.Parse("2006-01-02", stringValue)
	if err != nil {
		fmt.Println("Error parsing string date: ", err)
	}
	return parsed
}

func stringT(stringValue string) string {
	return stringValue
}

func generateFLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}
