# Project sampleproject

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```


## How to run sampleproject

- start a database container

```bash
docker run --name mariadb-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -e MYSQL_USER=user -e MYSQL_PASSWORD=password -p 3306:3306 -d mariadb:latest
```

- build binaries
  
```bash
barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ go build -o main cmd/api/main.go 
barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ go build -o migrate cmd/migrate/migrate.go
```

- open a terminal and run:

```
barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ ./migrate 
2025/03/12 20:14:15 Running database migrations...
2025/03/12 20:14:15 Migration completed successfully
barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ ./main 
Server running on :8080
```

- from another terminal run:

```
barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ curl -X POST http://localhost:8080/users -H "Content-Type: application/json"      -d @data.json
"post was successful"
barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ curl http://localhost:8080/users
[{"Name":"betty","Username":"fgtrfGG43","Email":"betty@example.com"},{"Name":"giannakis","Username":"vlasiboy666","Email":"giannakis@example.com"}]
barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$
```
