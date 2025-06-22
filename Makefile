GO_FILES := main.go
BIN := ./build/go-dumb-pw

all: $(BIN)

$(BIN):
	go build -o $(BIN) $(GO_FILES)

clean:
	rm -f $(BIN)

install:
	cp $(BIN) /usr/local/bin/go-dumb-pw && rm -fv /usr/local/bin/dumb-pw
	ln -s /usr/local/bin/go-dumb-pw /usr/local/bin/dumb-pw
