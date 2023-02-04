.PHONY: build
build: 
	@go build -o build/villianous

.PHONY: clean
clean:
	@rm -rf build/*