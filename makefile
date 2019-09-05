.PHONY: run
run:
	export GO111MODULE=on && go build .
	./pg
