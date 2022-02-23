package ammunition

import (
	"bytes"
	"fmt"

	"github.com/dblencowe/CheekyBreekiBot/helper"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type AmmoGraph struct {
}

func generatePenetrationItems(ammo []TarkovAmmunition) []opts.ScatterData {
	items := make([]opts.ScatterData, 0)
	for i := 0; i < len(ammo); i++ {
		items = append(items, opts.ScatterData{
			Value:        ammo[i].Ballistics.PenetrationPower,
			Symbol:       "roundRect",
			SymbolSize:   20,
			SymbolRotate: 10,
		})
	}

	return items
}

func generateDamageItems(ammo []TarkovAmmunition) []opts.ScatterData {
	items := make([]opts.ScatterData, 0)
	for i := 0; i < len(ammo); i++ {
		items = append(items, opts.ScatterData{
			Value:        ammo[i].Ballistics.Damage,
			Symbol:       "roundRect",
			SymbolSize:   20,
			SymbolRotate: 10,
		})
	}

	return items
}

func makeChart(ammo []TarkovAmmunition) *charts.Scatter {
	chart := charts.NewScatter()
	chart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Ammo Chart",
		}),
	)
	chart.
		AddSeries("Penetration", generatePenetrationItems(ammo)).
		AddSeries("Damage", generateDamageItems(ammo))

	return chart
}

func NewAmmoGraph(ammo []TarkovAmmunition) *bytes.Buffer {
	chart := makeChart(ammo)
	chartBuffer := &bytes.Buffer{}
	chart.Render(chartBuffer)
	helper.CreateFile("/tmp/example.html", chartBuffer.Bytes())
	imageBuffer := &bytes.Buffer{}
	bytes := GenerateImage(chartBuffer.String())
	imageBuffer.Write(bytes)
	return imageBuffer
}
