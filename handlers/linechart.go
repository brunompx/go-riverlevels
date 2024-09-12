package handlers

import (
	"math/rand"
	"net/http"

	"github.com/brunompx/go-riverlevels/views"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCntLine = 6
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func HandleLineChart(w http.ResponseWriter, r *http.Request) {

	chart := lineSmoothArea()

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

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}
