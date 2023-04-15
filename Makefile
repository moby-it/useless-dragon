.PHONY: all build run vendor

run:
	go run .

vendor:
	go mod tidy
	go mod vendor

prebuild:
	rm -rf bin
	mkdir -p bin/assets/enemies
	cp -r assets bin

build-linux: prebuild 
	env GOOS=linux GOARCH=amd64 go build -o bin/ .
	tar -zcvf linux.tgz bin/

build-windows: prebuild
	env GOOS=windows GOARCH=amd64 go build -o bin/ .
	zip -r windows bin

bundle-windows: prebuild build-windows

bundle-linux: prebuild build-linux

