go:
	go run main.go

test:
	./scripts/test.sh
	./scripts/clean.sh

dummy:
	./scripts/test.sh

clean:
	./scripts/clean.sh
