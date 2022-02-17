.PHONY: install test format clean
install:
	go install user/gfind
test:
	go test user/gfind
format:
	go fmt user/gfind
clean:
	rm -rf ./bin/*
	rm -rf ./pkg/*