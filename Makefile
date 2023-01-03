TAG=default

test:
	./scripts/setup_test.sh
	go test

publish:
	./scripts/publish.sh $(TAG)