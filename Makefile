all: run

BIN=2048
MAIN=src/main.go src/state.go src/utils.go src/keys.go src/item.go

run:
	go run ${MAIN}

build:
	go build -o '2048' ${MAIN}

clean:
	rm -rf ${BIN}
