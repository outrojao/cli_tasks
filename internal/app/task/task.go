package task

type Task struct {
	//para interpretação da lib encoding/json é necessário que o nome dos atributos estejam em maisculos em refenciados em json type
	Id   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func CreateTask(id int, taskName string) *Task {
	return &Task{
		Id:   id,
		Name: taskName,
		Done: false,
	}
}
