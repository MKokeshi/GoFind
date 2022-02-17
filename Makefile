.PHONY: install
install:
	go install user/gfind

.PHONY: test
test:
	go test user/gfind

.PHONY: format
format:
	go fmt user/gfind

.PHONY: clean
clean:
	rm -rf ./bin/*
	rm -rf ./pkg/*