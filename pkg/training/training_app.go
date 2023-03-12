package training

import (
	"math"
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/utils"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/sirupsen/logrus"
)

type TrainignApp struct {
	model   core.Model
	dataset *data.Dataset
	csts    *utils.Constants

	grid          *ui.Grid
	layersWidget  *LayersWidget
	spikeWidget   *SpikeWidget
	metricsWidget *MetricsWidget

	speed     float64
	isStopped bool
	metrics   map[string]float64
}

func (app *TrainignApp) Open() {
	if err := ui.Init(); err != nil {
		logrus.Fatalf("failed to initialize termui: %v", err)
	}
	app.grid = ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	app.grid.SetRect(0, 0, termWidth, termHeight)

	app.spikeWidget = NewSpikeWidget(app.model.GetInput(), int(app.csts.MaxTime))
	app.layersWidget = NewLayersWidget(app.model.GetAllLayer())
	app.metricsWidget = NewMetricsWidget(&app.metrics)

	app.grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0*(1.0/6.0),
				ui.NewRow(1.0/2, app.layersWidget),
				ui.NewRow(1.0/2, app.metricsWidget),
			),
			ui.NewCol(1.0*(5.0/6.0), app.spikeWidget),
		),
	)

	app.Render()
}

func (app *TrainignApp) observe() {
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<Down>":
				app.layersWidget.ScrollDown()
				app.spikeWidget.layer = app.model.GetLayer(app.layersWidget.SelectedRow)
			case "<Up>":
				app.layersWidget.ScrollUp()
				app.spikeWidget.layer = app.model.GetLayer(app.layersWidget.SelectedRow)
			case "<Left>":
				app.speed = utils.ClampFloat(math.Floor(app.speed*0.9), 10, 1000)
				app.metrics["speed"] = app.speed
			case "<Right>":
				app.speed = utils.ClampFloat(math.Floor(app.speed*1.1), 10, 1000)
				app.metrics["speed"] = app.speed
			case "q", "<C-c>":
				app.Stop()
			}
		} else if e.Type == ui.ResizeEvent {
			payload := e.Payload.(ui.Resize)
			app.grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			app.Render()
		}
	}
}

func (app *TrainignApp) Start() {
	go app.observe()
	for sample := range app.dataset.Cycle(1000) {
		app.model.Clear()
		app.model.Fit(sample.X, sample.Y)
		app.Render()
		time.Sleep(time.Duration(app.speed) * time.Millisecond)
		if app.isStopped {
			break
		}
	}
	app.isStopped = false
}

func (app *TrainignApp) Stop() {
	app.isStopped = true
}

func (app *TrainignApp) Render() {
	ui.Render(app.grid)
}

func (app *TrainignApp) Close() {
	ui.Close()
}

func NewTrainingApp(model core.Model, dataset *data.Dataset, csts *utils.Constants) *TrainignApp {
	metrics := make(map[string]float64)
	metrics["speed"] = float64(100)
	app := &TrainignApp{
		model:     model,
		dataset:   dataset,
		csts:      csts,
		isStopped: false,
		speed:     metrics["speed"],
		metrics:   metrics,
	}
	return app
}
