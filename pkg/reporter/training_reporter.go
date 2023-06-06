package reporter

import (
	"spiky/pkg/core"
	"spiky/pkg/reporter/widget"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/sirupsen/logrus"
)

type TrainingReporter struct {
	trainer *core.Trainer

	model   *core.Model
	dataset core.IDataset
	config  *core.ModelConfig

	grid          *ui.Grid
	layersWidget  *widget.LayersWidget
	spikeWidget   *widget.SpikeWidget
	metricsWidget *widget.MetricsWidget
}

func NewTrainingReporter(trainer *core.Trainer, config *core.ModelConfig) *TrainingReporter {
	app := &TrainingReporter{
		trainer: trainer,
		model:   nil,
		dataset: nil,
		config:  config,
	}
	trainer.Subscribe(app)
	return app
}

func (obs *TrainingReporter) OnStart(model *core.Model, dataset core.IDataset) {
	if err := ui.Init(); err != nil {
		logrus.Fatalf("failed to initialize termui: %v", err)
	}
	obs.model = model
	obs.dataset = dataset

	obs.grid = ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	obs.grid.SetRect(0, 0, termWidth, termHeight)

	obs.spikeWidget = widget.NewSpikeWidget(obs.model.GetInput(), int(obs.config.MaxTime))
	obs.layersWidget = widget.NewLayersWidget(obs.model.GetAllLayer())
	obs.metricsWidget = widget.NewMetricsWidget()
	obs.metricsWidget.Set("speed", float64(obs.trainer.GetSpeed()))

	noticeWidget := widgets.NewParagraph()
	noticeWidget.Text = "Select layer: [↑ or ↓](fg:cyan), Change speed: [← or →](fg:green)"

	obs.grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0*(1.0/6.0),
				ui.NewRow(1.0/2, obs.layersWidget),
				ui.NewRow(1.0/2, obs.metricsWidget),
			),
			ui.NewCol(1.0*(5.0/6.0),
				ui.NewRow(11.0/12.0, obs.spikeWidget),
				ui.NewRow(1.0/12.0, noticeWidget),
			),
		),
	)

	go obs.observe()
}

func (app *TrainingReporter) OnEpochStart(iterations int) {

}

func (app *TrainingReporter) OnEpochEnd() {

}

func (app *TrainingReporter) OnStep(metrics *map[string]float64) {
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

func (app *TrainingReporter) observe() {
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<Down>":
				app.layersWidget.ScrollDown()
				app.spikeWidget.SetLayer(app.model.GetLayer(app.layersWidget.SelectedRow))
				app.render()
			case "<Up>":
				app.layersWidget.ScrollUp()
				app.spikeWidget.SetLayer(app.model.GetLayer(app.layersWidget.SelectedRow))
				app.render()
			case "<Left>":
				app.trainer.SpeedDown()
				app.metricsWidget.Set("speed", app.trainer.GetSpeed())
				app.render()
			case "<Right>":
				app.trainer.SpeedUp()
				app.metricsWidget.Set("speed", app.trainer.GetSpeed())
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
