package core

type IObserver interface {
	OnStart(model IModel, dataset IDataset, iterations int)
	OnUpdate(metrics *map[string]float64)
	OnStop()
}
