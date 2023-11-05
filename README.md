
# TODO REST API


## Simple Rest App (In Memory Data Storage)
### App Libraries

- Server build on top of base library `net/http`

- Vendor Libraries
    -
    * `httprouter`
    * `alice`

        


## Features
- Create task using post api
- Create task using csv task files
- Search task by ID



### Run Dev App
Running on port 4000
```sh
go run ./cmd/web
```

### Docker Deployment
```sh
docker build -f Dockerfile -t goapp . 
docker run -p 4000:4000 --name goappweb goapp:latest  
```


## API Docs
Checkout Postman Collections

[Download Postman Collections](docs/TODO APP.postman_collection.json)

* Check Health
```curl
curl --location 'http://localhost:4000'
```

* Get ALL Task
```curl
curl --location 'http://localhost:4000/todo/list'
```

* Search Task By Id
```curl
curl --location 'http://localhost:4000/todo/get/:id'
```

* Create Task
```curl
curl --location 'http://localhost:4000/todo/create' \
--header 'Content-Type: application/json' \
--data '{
    "description": "task2",
    "is_completed": false,
    "time": "2023-11-04T15:04:05Z"
}'
```

* Create Task From CSV File

```curl
curl --location 'http://localhost:4000/todo/upload' \
--form 'file=@"task_list.csv"'
```