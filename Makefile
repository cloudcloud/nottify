# use docker-compose
COMPOSE=docker-compose -p nottify -f docker-compose.yml
NOTTIFY=./nott
NOTTIFY_BKP=./nott.bkp

# primarily for local usage
install:
	go get ./...

test:
	go test ./...

coverage: INT?=0
coverage: OUT?=../coverage
coverage: PA?=.
coverage:
	@if [ -f "$(OUT).json" ]; then rm $(OUT).json; fi
	@if [ -f "$(OUT).html" ]; then rm $(OUT).html; fi
	@RUNINTEGRATION="$(INT)" gocov test `go list $(PA)/... | grep -v /vendor/` > "$(OUT).json"
	@gocov-html "$(OUT).json" > "$(OUT).html"
	@echo "mode: count" > "$(OUT).coverage" && echo "Putting coverage metrics for codecov.io"
	@for file in `go list $(PA)/... | grep -v /vendor/`; do \
		go test -coverprofile="$(OUT).cover" -covermode=count "$$file"; \
		grep -h -v "^mode: " "$(OUT).cover" >> "$(OUT).coverage"; \
	done
	@sed -i 's#github.com/cloudcloud/nottify/##' "$(OUT).coverage" && mv "$(OUT).coverage" "$(OUT).out"

clean:
	@if [ -f "$(NOTTIFY)" ]; then rm "$(NOTTIFY)"; fi
	@if [ -f "$(NOTTIFY_BKP)" ]; then rm "$(NOTTIFY_BKP)"; fi

# primarily for containered usage
compile:
	@if [ -f "$(NOTTIFY)" ]; then mv "$(NOTTIFY)" "$(NOTTIFY_BKP)"; fi
	@GOOS=linux go build -a -tags netgo -ldflags '-w' -o "$(NOTTIFY)" . 2>&1
	@if [ -f "$(NOTTIFY)" -a -f "$(NOTTIFY_BKP)" ]; then rm "$(NOTTIFY_BKP)"; \
		elif [ ! -f "$(NOTTIFY)" ]; then mv "$(NOTTIFY_BKP)" "$(NOTTIFY)"; fi

build: compile
	$(COMPOSE) pull
	$(COMPOSE) build

up:
	$(COMPOSE) up -d nottify

down:
	$(COMPOSE) kill
	$(COMPOSE) rm --force
