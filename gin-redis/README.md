# Redis Sample Application

This is a sample application demonstrating basic operations with Redis using Go and the Gin web framework. 

## Setup and Run

 Start the Redis server using the following command:
```bash
docker-compose up -d
 ```
 Navigate to the project directory and run the server using the following command:

```bash
go run main.go
 ```

## Usage 

GET Value
To get the value of mykey, you can use this curl command:

```bash
curl -X GET "http://localhost:8083/api/mykey"
```

SET Value
To set the value of mykey, you can use this curl command:

```bash
curl -X POST "http://localhost:8083/api/mykey" -H "Content-Type: application/json" -d '{ "value": "your_value" }'
 ```

DELETE Value
To delete mykey, you can use this curl command:

```bash
curl -X DELETE "http://localhost:8083/api/mykey"
 ```

UPDATE Value
To update the value of mykey, you can use this curl command:

```bash
curl -X PUT "http://localhost:8083/api/mykey" -H "Content-Type: application/json" -d '{ "value": "new_value" }'
 ```
