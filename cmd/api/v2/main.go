package main

import (
	"github.com/vlbarou/sampleproject/config"
	"github.com/vlbarou/sampleproject/internal/database"
	"github.com/vlbarou/sampleproject/internal/server/v2"
	"log"
)

// https://chatgpt.com/share/67d0c938-6b14-8009-8b04-4b661fa5722c

// docker run --name mariadb-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -e MYSQL_USER=user -e MYSQL_PASSWORD=password -p 3306:3306 -d mariadb:latest
// docker rm -f mariadb-container

// barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ go build -o main cmd/api/v2/main.go
// barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ go build -o migrate cmd/migrate/migrate.go

// curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d @data.json
// curl http://localhost:8080/users
// curl http://localhost:8080/user?id=13623
func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	defer database.Close(db)

	srv := v2.NewServer(cfg, db)
	log.Fatal(srv.Start())
}
