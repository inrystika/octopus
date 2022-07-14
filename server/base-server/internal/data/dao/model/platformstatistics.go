package model

type PlatformStatTrainJob struct {
	PendingNum int
	RunningNum int
}

type PlatformStatSummary struct {
	TrainJob *PlatformStatTrainJob
}
