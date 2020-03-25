.PHONY: all restart restart-app build-java

build-java:
	$(info Building Java docker image)
	@cd micro-inventory-service; \
	docker-compose -f docker-compose.build.yml up \


all: build-java
	$(info Deploy all services...)
	@docker-compose up -d --build

restart:
ifndef SERVICE
	$(error The SERVICE variable is required.)
endif
	@docker-compose up -d --build --force-recreate --remove-orphans $(SERVICE)

restart-app: build-java
	@docker-compose up -d --build --force-recreate micro-api-gateway micro-inventory micro-review micro-price