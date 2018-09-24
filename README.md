#Request

## POST 
http://localhost:8080/tasks
```json
{
    "task": "laver le chien",
    "date" : "2016-11-02T08:18:20+02:00",
    "done": false
}
```

## GET 
http://localhost:8080/tasks/1

## PUT 
http://localhost:8080/tasks/1
```json
{   
    "task": "laver le chien",
    "date" : "2016-11-02T08:18:20+02:00",
    "done": false
}
``