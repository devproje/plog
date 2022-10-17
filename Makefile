TAG=default

init:
	go mod tidy

publish:
	./scripts/publish.sh $(TAG)