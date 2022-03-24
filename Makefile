.PHONY: build test clean docker


build:
	go mod tidy
	go build -o ./bin/auto-check main.go
	@echo  "[INFO] go build successful"
	@echo  "[INFO] you can cd bin to execute it"

run :
	./bin/auto-check

clean :
	rm -f ./bin/*