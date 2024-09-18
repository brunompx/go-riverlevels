package handlers

import (
	"math/rand"
	"net/http"

	"github.com/brunompx/go-riverlevels/views"
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

	views.Linechart(chart).Render(r.Context(), w)
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
	// Put data into instance
	line.SetXAxis([]string{"a", "b", "c", "f", "h", "j", "e", "r", "t", "y", "hu", "ji"})

	line.AddSeries("Category A", sals4).
		AddSeries("Category b", sals2).AddSeries("Category c", sals3).AddSeries("Category d", sals1).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 15, Symbol: "diamond"},
		))

	return line
}

func generateFLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}
