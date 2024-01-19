# demo task intended for corey


to run the api server
```
cd cmd
```

and

```
go run .
```


## available end points

### /contact
example curl request
```sh
curl --request POST \
  --url http://localhost:8000/contact \
  --header 'Content-Type: application/json' \ \
  --data '{
	"name": "John Doe 3",
	"email": "john.doe3@example.com"
}
'
```
response
```json
{
	"ID": 6,
	"name": "John Doe 3",
	"email": "john.doe3@example.com",
	"Tasks": null
}
```
### /contact/{id}
example curl request
```sh
curl --request GET \
  --url http://localhost:8000/contact/6 \
  --header 'Content-Type: application/json' \
```
response
```json
{
	"ID": 6,
	"name": "John Doe 3",
	"email": "john.doe3@example.com",
	"Tasks": null
}
```
### /task
example curl request
```sh
curl --request POST \
  --url http://localhost:8000/task \
  --header 'Content-Type: application/json' \ 
  --data '{
	"title": "Task 3",
	"description": "Complete task 3",
	"reminder": "2024-01-18T12:00:00Z",
	"contact_id": 6
}
'
```
response
```json
{
	"ID": 5,
	"title": "Task 3",
	"description": "Complete task 3",
	"reminder": "2024-01-18T12:00:00Z",
	"contact_id": 6
}
```
### /contact
example curl request
```sh
curl --request GET \
  --url http://localhost:8000/contact \
  --header 'Content-Type: application/json' \
```
response
```json
[
	{
		"ID": 1,
		"name": "John Doe",
		"email": "john.doe@example.com",
		"Tasks": null
	}
]
```