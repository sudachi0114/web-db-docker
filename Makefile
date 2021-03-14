.PHONY: build
build:
	go build -o bin/main .

.PHONY: clean
clean:
	rm bin/main

.PHONY: compose/up
compose/up: compose/down
	docker-compose up -d

.PHONY: compose/down
compose/down:
	docker-compose down
