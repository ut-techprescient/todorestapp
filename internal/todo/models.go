package todo

import (
	"time"

	"todorestapp.ut/internal/utils"
)

type Task struct {
	ID          uint64    `json:"id"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	Time        time.Time `json:"time"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func (t *Tasks) Insert(task Task) uint64 {
	task.ID = uint64(len(t.Tasks))
	t.Tasks = append(t.Tasks, task)
	return task.ID
}

func (t *Tasks) GetAll() *Tasks {
	return t
}

func stringToTime(val string) time.Time {
	times, err := time.Parse(val, time.UTC.String())
	if err != nil {
		return time.Now().UTC()
	}
	return times
}

func (t *Tasks) CreateTaskFromFile(filePath string) (bool, error) {

	data, err := utils.ReadCsvFile(filePath)
	if err != nil {
		return false, nil
	}

	lenTask := len(t.Tasks)
	for i, line := range data {
		lenTask += i
		if i < 1 {
			continue
		}
		var task Task
		task.ID = uint64(lenTask)
		for j, field := range line {
			switch j {
			case 0:
				task.Description = string(field)
			case 1:
				task.IsCompleted = field == "true"
			case 2:
				task.Time = stringToTime(field)
			}
		}
		t.Tasks = append(t.Tasks, task)
	}

	return true, nil
}

func (t *Tasks) SearchById(id int) *Task {
	for _, task := range t.Tasks {
		if int(task.ID) == id {
			return &task
		}
	}
	return nil
}
