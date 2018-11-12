.PHONY: clean env-up run

all: clean env-up run

dev: run

##### BUILD
build:
	@echo "Build ..."
	@dep ensure
	@go build
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@cd fixtures && docker-compose up --force-recreate -d
	@echo "Sleep 10 seconds in order to let the environment setup correctly"
	@sleep 10
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fixtures && docker-compose down
	@echo "Environment down"

##### RUN
run:
	@echo "Start app ..."
	@bee run

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/loveHome-* bill
	@docker rm -f -v `docker ps -a --no-trunc | grep "bill" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "bill" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"
