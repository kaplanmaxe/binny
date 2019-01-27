.PHONY: run-dectobin run-bintodec run-twoscomplementbin

build:
	@(go build -o ./bin/binny main.go)

run-dectobin: build
	./bin/binny dectobin 10

run-bintodec: build
	./bin/binny bintodec 00001010

run-twoscomplementbin: build
	./bin/binny twoscomplementbin 00000001