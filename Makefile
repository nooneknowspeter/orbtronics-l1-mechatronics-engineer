.PHONY: all clean test

format:
	treefmt
composeup:
	docker compose up -d
composedown:
	docker compose down