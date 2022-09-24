

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info falemais-telzir backend commands:)
	$(info -> test                    runs available tests)
	$(info -> run                     runs application)

test:
	@go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

run:
	@echo "Running application... waiting a message:  Server running on port..."
	@go run main.go


