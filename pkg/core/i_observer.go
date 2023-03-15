package core

type IObserver interface {
	OnStart(model IModel, dataset IDataset, metrics map[string]float64, iterations int)
	OnUpdate(idx int)
	OnStop()
}
