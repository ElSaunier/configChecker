GO = go 
BIN = bin/

all: build test install

test: verity_test.go
	$(GO) $@ $<

install: main.go
	$(GO) $@ $<

build: main.go | $(BIN)
	$(GO) $@ $<

run: main.go
	$(GO) $@ $<

$(BIN):
	mkdir -p $(BIN)

clean: 
	@echo "[-] Suppression des fichiers objets"
	rm *.o
	@echo "[-] Suppression de l'exécutable"
	rm main