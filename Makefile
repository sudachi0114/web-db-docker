.PHONY: build
build:
	go build -o bin/main .

.PHONY: clean
clean: compose/down
	rm bin/main

.PHONY: migrate
migrate: compose/up
	docker-compose exec app go run /app/tools/db/migrate.go

.PHONY: compose/up
compose/up: compose/up/db
	docker-compose up -d app

.PHONY: compose/up/db
compose/up/db: compose/down
	docker-compose up -d db

.PHONY: compose/down
compose/down:
	docker-compose down
