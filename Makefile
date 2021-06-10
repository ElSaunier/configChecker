GO = go 
BIN = bin/
MOD = configChecker

all: install test run

test: verify_test.go
	$(GO) $@ $<

install:
	$(GO) $@ $<

build: $(BIN)
	$(GO) $@
	mv $(MOD) $(BIN)

run: 
	$(GO) $@ .

$(BIN):
	mkdir -p $(BIN)

clean: 
	@echo "[-] Suppression de l'exécutable"
	rm -r $(BIN)