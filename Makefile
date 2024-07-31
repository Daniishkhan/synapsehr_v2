run:
	templ generate && go run cmd/main.go

generate:
	templ generate

.PHONY: run generate