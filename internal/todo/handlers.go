package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"todorestapp.ut/cmd/config"
	"todorestapp.ut/internal/utils"
)

var tasks = &Tasks{}

func TodoTaskListHandler(appConfig *config.AppEnv) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		appConfig.InfoLog.Print(ps)

		serializer := TaskSerializers{Tasks: tasks.GetAll()}
		res, err := serializer.Encode()

		if err != nil {
			appConfig.AppError.ServerError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}

}

func TodoTaskGetHandler(appConfig *config.AppEnv) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		appConfig.InfoLog.Println(ps)
		taskId, _ := strconv.Atoi(ps.ByName("id"))

		task := tasks.SearchById(taskId)

		res, err := json.Marshal(task)

		if err != nil {
			appConfig.AppError.ServerError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}

}

func TodoTaskCreateHandler(appConfig *config.AppEnv) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		decoder := json.NewDecoder(r.Body)
		var task Task
		err := decoder.Decode(&task)
		if err != nil {
			appConfig.ErrorLog.Printf("Failed while processing task request %s", err)
			appConfig.AppError.ClientError(w, 400)
			return
		}
		appConfig.InfoLog.Print(task)
		tasks.Insert(task)

		response := map[string]string{
			"task_id": strconv.Itoa(int(task.ID)),
			"message": "Task Registered Successfully",
		}

		res, err := json.Marshal(response)

		if err != nil {
			appConfig.AppError.ServerError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func TodoTaskUploadHandler(appConfig *config.AppEnv) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tempFile, err := utils.UploadFile(r)

		if err != nil {
			appConfig.ErrorLog.Printf("Failed while processing task file request %s", err)
			appConfig.AppError.ServerError(w, err)
			return
		}

		isCreated, err := tasks.CreateTaskFromFile(tempFile)

		if err != nil {
			appConfig.ErrorLog.Printf("Failed while processing task file request %s", err)
			appConfig.AppError.ServerError(w, err)
			return
		}

		response := map[string]any{
			"created": isCreated,
			"message": "Task File Processed Successfully",
		}

		res, err := json.Marshal(response)

		if err != nil {
			appConfig.AppError.ServerError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)

	}
}
