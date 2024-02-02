# 当前目录
CUR_DIR=$(shell pwd)
OUT_DIR=$(CUR_DIR)/out

compile:
	mkdir -p $(OUT_DIR)
	#go tool compile -S -N -l -o $(OUT_DIR)/originfun.o ./origingunc > $(OUT_DIR)/originfun.s
	#go tool compile -S -N -l -o $(OUT_DIR)/main.o . > $(OUT_DIR)/main.s
	go build -gcflags="-S -N -l" -o $(OUT_DIR)/main . 2> $(OUT_DIR)/main.s
build:
	mkdir -p $(OUT_DIR)
	go build -gcflags="-N -l" -o $(OUT_DIR)/main .

run:
	$(OUT_DIR)/main

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
