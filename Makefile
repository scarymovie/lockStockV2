prepare:
	mkdir -p development
	mkdir -p logs
	sudo chown -R $(whoami):$(whoami) logs
	sudo chmod +w logs
	touch logs/app.log

	cp build/docker/go/.env.example development/.go.env

build:
	docker-compose -f docker-compose-dev.yml build

build-no-cache:
	docker-compose -f docker-compose-dev.yml build --no-cache

up:
	docker-compose -f docker-compose-dev.yml up -d

.PHONY: prepare build up
