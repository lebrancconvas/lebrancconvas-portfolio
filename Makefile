run:
	gin -a 2000 -p 2019 run server.go

dev:
	go run server.go

gin:
	gin run server.go