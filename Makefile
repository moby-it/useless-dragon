.PHONY: all build run

run:
	go run .

prebuild:
	rm -rf bin
	mkdir -p bin/assets/enemies
	cp -r assets bin

build-linux: prebuild 
	env GOOS=linux GOARCH=amd64 go build -o bin/ .

build-windows: prebuild
	env GOOS=windows GOARCH=amd64 go build -o bin/ .

zip: 
	tar -zcvf useless_dragon.tgz bin/

bundle-windows: prebuild build-windows zip

bundle-linux: prebuild build-linux zip

