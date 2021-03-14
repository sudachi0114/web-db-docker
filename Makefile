.PHONY: build
build:
	go build -o bin/main .

.PHONY: clean
clean:
	rm bin/main

.PHONY: compose/up
compose/up: compose/up/db
	docker-compose up -d app

.PHONY: compose/up/db
compose/up/db: compose/down
	docker-compose up -d db

.PHONY: compose/down
compose/down:
	docker-compose down
