package reporter

import (
	"spiky/pkg/core"
	"spiky/pkg/reporter/widget"
	"spiky/pkg/utils"

	ui "github.com/gizak/termui/v3"
	"github.com/sirupsen/logrus"
)

type TrainingReporter struct {
	trainer *core.Trainer

	model   core.IModel
	dataset core.IDataset
	csts    *utils.Constants

	grid          *ui.Grid
	layersWidget  *widget.LayersWidget
	spikeWidget   *widget.SpikeWidget
	metricsWidget *widget.MetricsWidget
}

func (obs *TrainingReporter) OnStart(model core.IModel, dataset core.IDataset, iterations int) {
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
	obs.metricsWidget = widget.NewMetricsWidget()

	obs.grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0*(1.0/6.0),
				ui.NewRow(1.0/2, obs.layersWidget),
				ui.NewRow(1.0/2, obs.metricsWidget),
			),
			ui.NewCol(1.0*(5.0/6.0), obs.spikeWidget),
		),
	)

	go obs.observe()
}

func (app *TrainingReporter) observe() {
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
				app.trainer.SpeedDown()
				app.metricsWidget.Set("waiting time", float64(app.trainer.GetWaitingTime()))
				app.render()
			case "<Right>":
				app.trainer.SpeedUp()
				app.metricsWidget.Set("waiting time", float64(app.trainer.GetWaitingTime()))
				app.render()
			case "q", "<C-c>":
				app.trainer.Stop()
			}
		} else if e.Type == ui.ResizeEvent {
			payload := e.Payload.(ui.Resize)
			app.grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			app.render()
		}
	}
}

func (app *TrainingReporter) OnUpdate(metrics *map[string]float64) {
	for k, v := range *metrics {
		app.metricsWidget.Set(k, v)
	}
	app.render()
}

func (app *TrainingReporter) OnStop() {
	ui.Close()
}

func (app *TrainingReporter) render() {
	ui.Render(app.grid)
}

func NewTrainingReporter(trainer *core.Trainer, csts *utils.Constants) *TrainingReporter {
	metrics := make(map[string]float64)
	metrics["speed"] = float64(100)
	app := &TrainingReporter{
		trainer: trainer,
		model:   nil,
		dataset: nil,
		csts:    csts,
	}
	trainer.Subscribe(app)
	return app
}
