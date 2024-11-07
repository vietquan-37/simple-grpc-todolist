package enum

type Status string

const (
	Pending   Status = "PENDING"
	Progress  Status = "INPROGRESS"
	Completed Status = "COMPLETED"
)
