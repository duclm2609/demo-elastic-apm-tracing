.PHONY: all restart restart-app build-java

setup:
	@docker create volume maven

all:
	$(info Deploy all services...)
	@docker compose up -d --build

restart:
ifndef SERVICE
	$(error The SERVICE variable is required.)
endif
	@docker compose up -d --build --force-recreate --remove-orphans $(SERVICE)

restart-app: build-java
	@docker compose up -d --build --force-recreate micro-api-gateway micro-inventory micro-review micro-price

destroy:
	@docker compose down