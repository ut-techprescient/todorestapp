package todo

import (
	"encoding/json"
)

type TaskSerializers struct {
	Tasks *Tasks
}

func (s *TaskSerializers) Encode() ([]byte, error) {
	return json.Marshal(s.Tasks)

}
