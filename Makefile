BINARYNAME	:= nhlbot
BINARYENDING	:=
ifeq ($(OS),Windows_NT)
	BINARYENDING = .exe
endif

all: clean
	go build -v -o $(BINARYNAME)$(BINARYENDING) main.go

.PHONY: clean
clean:
	rm -rf ./build $(BINARYNAME)$(BINARYENDING)
