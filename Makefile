.PHONY: all build run vendor

MAIN_PATH=cmd/useless_dragon/main.go

run:
	go run $(MAIN_PATH)
test:
	go test ./... -v
vendor:
	go mod tidy
	go mod vendor

prebuild:
	rm -rf bin
	mkdir -p bin/assets/enemies
	cp -r assets bin

build-linux: prebuild 
	env GOOS=linux GOARCH=amd64 go build -o bin/ $(MAIN_PATH)
	tar -zcvf linux.tgz bin/

build-windows: prebuild
	env GOOS=windows GOARCH=amd64 go build -o bin/ $(MAIN_PATH)
	zip -r windows bin

bundle-windows: prebuild build-windows

bundle-linux: prebuild build-linux

