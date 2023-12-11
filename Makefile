
run_test_v:
	go test -v ./...

run_test:
	go test ./...

run:
	go run main.go

curr_version:
	@echo "Current version: $(shell git describe --abbrev=0 --tags)"

new_version:
	@echo "Digite a nova versão: "
	@read version; \
	git tag $$version; \
	git push origin $$version
	go list -m -json | curl -X PATCH -d @- https://proxy.golang.org/$(go list -m)

