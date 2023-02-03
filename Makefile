.PHONY = run run-config

run:
	go run main.go

run-config:
	go run main.go --config ./.env
