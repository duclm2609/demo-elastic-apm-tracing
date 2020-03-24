.PHONY: all restart build-java-docker

build-java:
	$(info Building Java docker image)
	@cd micro-inventory-service; \
	docker-compose -f docker-compose.build.yml up \


all: build-java
	$(info Deploy all services...)
	@docker-compose up -d --build --force-recreate --remove-orphans --quiet-pull

restart:
ifndef SERVICE
	$(error The SERVICE variable is required.)
endif
	@docker-compose up -d --build --force-recreate --remove-orphans $(SERVICE)