.PHONY: run-dectobin run-bintodec run-twoscomplementbin test

build:
	@(go build -o ./bin/binny main.go)

test:
	@(go test ./...)

run-dectobin: build
	./bin/binny dectobin 10

run-bintodec: build
	./bin/binny bintodec 00001010

run-twoscomplementbin: build
	./bin/binny twoscomplementbin 00000001

run-floattodec: build
	./bin/binny floattodec 01000010011101110000000000000000