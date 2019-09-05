.PHONY: run
run:
	export GO111MODULE=on && go build -o db .
	./db
