binary = scratch-golang-gin

.DEFAULT_GOAL := run

run: clean code_quality test
	@go build
	@./$(binary)

test:
	@echo 'Running Scratch Tests'
	@go test $(go list ./... | grep -v /vendor/)
	@echo 'Tests Completed'

code_quality:
	@go fmt
	@go vet
	@go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

clean:
	@rm -f $(binary)
