package core

type IObserver interface {
	OnStart(model *Model, dataset IDataset)
	OnStep(metrics *map[string]float64)
	OnEpochStart(iterations int)
	OnEpochEnd()
	OnStop()
}
