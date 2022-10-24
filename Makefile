test:
	./scripts/clean.sh
	./scripts/test.sh

tests:
	./scripts/clean.sh
	./scripts/test.sh
	./scripts/downloads.sh

dummy:
	./scripts/test.sh

clean:
	./scripts/clean.sh

downloads:
	./scripts/downloads.sh

