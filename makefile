BIN=bin
SOURCE_SONG=src/*.java

dvd:
	@go run .

build-song:
	@javac -d $(BIN) $(SOURCE_SONG)

clean:
	@rm -rf $(BIN)

# rule to accept commands like [make build something]
# regardless of whether 'something' exists or not
%:
	@:
