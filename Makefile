# Makefile actions need to be initiated by a tab

run:
	go run app/web/main.go

test:
	go test ./...