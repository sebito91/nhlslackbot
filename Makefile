all: clean
	go build -v -o nhlbot main.go

.PHONY: clean
clean:
	rm -rf ./build nhlbot
