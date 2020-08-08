package entry

type WorkerEntry struct {
	Number int `json:"number" form:"number" binding:"required"`
}

func NewWorkerEntry() *WorkerEntry {
	return &WorkerEntry{}
}
