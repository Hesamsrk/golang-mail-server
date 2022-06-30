PORT=9090

prepare:
	go mod tidy

dev:
	 go run main.go --port ${PORT} 

dev-nodemon:
	nodemon --exec (go run main.go --port ${PORT})

db:
	docker run --name golang-email-server \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_DB=staging \
	-e POSTGRES_USER=pg \
	-p 9099:5432 \
	-d postgres