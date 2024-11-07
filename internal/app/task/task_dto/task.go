package task_dto

type AddTaskReq struct {
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	TotalTime int    `json:"total_time"`
}
type DelTaskReq struct {
	ID int `json:"id"`
}
type UpdateTaskReq struct {
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	TotalTime int    `json:"total_time"`
	Recover   bool   `json:"recover"`
}
