COMPOSE_PACKAGE=nottify
COMPOSE=docker-compose -p $(COMPOSE_PACKAGE)

coverage:
	go test -race -coverprofile=/tmp/cov ./...
	go tool cover -html=/tmp/cov -o ./coverage.html

down:
	$(COMPOSE) down

install:
	go get ./...

logs:
	$(COMPOSE) logs -f

stop:
	$(COMPOSE) stop

test:
	go test -race ./...

up:
	$(COMPOSE) up -d nottify

