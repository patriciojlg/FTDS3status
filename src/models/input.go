package models

type Request struct {
	Command     string                 `json:"command"`
	Account     string                 `json:"account"`
	BatchID     string                 `json:"batch_id"`
	Status      string                 `json:"status"`
	TaskID      string                 `json:"task_id"`
	ServiceTask string                 `json:"service_task"`
	Payload     map[string]interface{} `json:"payload"`
}
