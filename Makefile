GO = go 
BIN = bin/
MOD = configChecker

all: build test install

test: verify_test.go
	$(GO) $@ $<

install:
	$(GO) $@ $<

build: $(BIN)
	$(GO) $@
	mv $(MOD) $(BIN)

run: main.go
	$(GO) $@ $<

$(BIN):
	mkdir -p $(BIN)

clean: 
	@echo "[-] Suppression de l'exécutable"
	rm -r $(BIN)