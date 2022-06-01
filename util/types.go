package util

type WorkStatus struct {
	Open     string
	Occupied string
	Finished string
}

func GetWorkStatusMapping() WorkStatus {

	workStatus := WorkStatus{
		Open:     "open",
		Occupied: "occupied",
		Finished: "finished",
	}

	return workStatus
}
