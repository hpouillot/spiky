package observer

import (
	"math"
	"spiky/pkg/core"
	"spiky/pkg/observer/widget"
	"spiky/pkg/utils"

	ui "github.com/gizak/termui/v3"
	"github.com/sirupsen/logrus"
)

type TrainingObserver struct {
	model   core.IModel
	dataset core.IDataset
	csts    *utils.Constants

	grid          *ui.Grid
	layersWidget  *widget.LayersWidget
	spikeWidget   *widget.SpikeWidget
	metricsWidget *widget.MetricsWidget

	speed     float64
	isStopped bool
	metrics   map[string]float64

	errors int
}

func (obs *TrainingObserver) OnStart(model core.IModel, dataset core.IDataset, metrics map[string]float64, iterations int) {
	if err := ui.Init(); err != nil {
		logrus.Fatalf("failed to initialize termui: %v", err)
	}
	obs.model = model
	obs.dataset = dataset

	obs.grid = ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	obs.grid.SetRect(0, 0, termWidth, termHeight)

	obs.spikeWidget = widget.NewSpikeWidget(obs.model.GetInput(), int(obs.csts.MaxTime))
	obs.layersWidget = widget.NewLayersWidget(obs.model.GetAllLayer())
	obs.metricsWidget = widget.NewMetricsWidget(&obs.metrics)

	obs.grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0*(1.0/6.0),
				ui.NewRow(1.0/2, obs.layersWidget),
				ui.NewRow(1.0/2, obs.metricsWidget),
			),
			ui.NewCol(1.0*(5.0/6.0), obs.spikeWidget),
		),
	)

	obs.render()
	go obs.observe()
}

func (app *TrainingObserver) observe() {
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<Down>":
				app.layersWidget.ScrollDown()
				app.spikeWidget.SetLayer(app.model.GetLayer(app.layersWidget.SelectedRow))
			case "<Up>":
				app.layersWidget.ScrollUp()
				app.spikeWidget.SetLayer(app.model.GetLayer(app.layersWidget.SelectedRow))
			case "<Left>":
				app.speed = utils.ClampFloat(math.Floor(app.speed*0.9-1), 0, 10000)
				app.metrics["speed"] = app.speed
			case "<Right>":
				app.speed = utils.ClampFloat(math.Ceil(app.speed*1.1+1), 0, 10000)
				app.metrics["speed"] = app.speed
			case "q", "<C-c>":
				app.OnStop()
			}
		} else if e.Type == ui.ResizeEvent {
			payload := e.Payload.(ui.Resize)
			app.grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			app.render()
		}
	}
}

func (app *TrainingObserver) OnUpdate(idx int) {
	app.render()
}

func (app *TrainingObserver) OnStop() {
	app.isStopped = true
	ui.Close()
}

func (app *TrainingObserver) render() {
	ui.Render(app.grid)
}

func NewTrainingObserver(csts *utils.Constants) *TrainingObserver {
	metrics := make(map[string]float64)
	metrics["speed"] = float64(100)
	app := &TrainingObserver{
		model:     nil,
		dataset:   nil,
		csts:      csts,
		isStopped: false,
		speed:     metrics["speed"],
		metrics:   metrics,
		errors:    0,
	}
	return app
}
