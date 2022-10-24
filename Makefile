go:
	go run main.go

test:
	./scripts/clean.sh
	./scripts/test.sh

dummy:
	./scripts/test.sh

clean:
	./scripts/clean.sh

downloads:
	./scripts/downloads.sh

