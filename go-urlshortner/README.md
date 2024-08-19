# go-urlshortner
A simple golang based url shortner


# Requirments to run
1. Golang [How to install Golang](https://go.dev/doc/install)
2. Docker [How to install Docker?](https://docs.docker.com/engine/install/)


# Setting up the project
Run the following commands to setup the project


``` bash
git clone https://github.com/keploy/samples-go.git && cd samples-go/go-urlshortner
go mod download
```


# Running app

## Let's start the MySql Instance

``` bash
sudo docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=uss -p 3306:3306 --rm mysql:latest
```

# Capture the Testcases

``` bash
sudo -E env PATH=$PATH oss record -c "./go-urlshortner"
```

To generate testcases we just need to make some API calls. You can use Postman, Hoppscotch, or simply curl


1. Root Endpoint:
```bash
-> curl -X GET http://localhost:9090/
```


2. Health Check:
```bash
-> curl -X GET http://localhost:9090/healthcheck
```


3. Short URL:
```bash
-> curl -X POST http://localhost:9090/shorten -H "Content-Type: application/json" -d '{"url": "https://github.com"}'
```


4. Resolve short code:

```bash
-> curl -X GET http://localhost:9090/resolve/4KepjkTT
```

Now both these API calls were captured as a testcase and should be visible on the Keploy CLI. You should be seeing an app named keploy folder with the test cases we just captured and data mocks created.

![alt text](https://github.com/Hermione2408/samples-go/blob/app/go-urlshortner/img/keploy_record.png?raw=true)

# Run the captured testcases

Now that we have our testcase captured, run the test file.

```bash
sudo -E env PATH=$PATH oss test -c "./go-urlshortner" --delay 20
```

So no need to setup dependencies like mySql, web-go locally or write mocks for your testing.

The application thinks it's talking to mysql 😄

We will get output something like this:

![alt text](https://github.com/Hermione2408/samples-go/blob/app/go-urlshortner/img/keploy_test.png?raw=true)