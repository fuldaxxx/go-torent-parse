.PHONY: go-build
go-build:
	go build -o go-torent-parse ./cmd/main.go

.PHONY: run-app
run-app:
	./go-torent-parse
