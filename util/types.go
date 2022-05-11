package util

type WorkStatus struct {
	Open     string
	Occupied string
	Finished string
}

var workStatus *WorkStatus

func GetWorkStatusMapping() *WorkStatus {
	if workStatus == nil {
		workStatus = &WorkStatus{
			Open:     "open",
			Occupied: "occupied",
			Finished: "finished",
		}
	}
	return workStatus
}
