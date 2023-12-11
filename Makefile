M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: help
help: ## Prints this help message
	@grep -E '^[a-zA-Z_-]+:.*## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

### VARIABLES ###

APPNAME=petstore

ifdef APPWATCH
$(info $(M) GOW test for watch mode)
ifneq ($(@which gow),"")
$(info $(M) GOW found for watch mode)
WATCH=1
else
$(info $(M) GOW NOT found for watch mode. Using standart GO command)
endif
endif

### COMMANDS ###

.PHONY: run_client
run: ## Execute the go run command (can add a watch mode using APPWATCH=1)
ifeq ($(WATCH),1)
	$(info $(M) Starting an instance of $(APPNAME)_client in watch mode)
	@cd $(PWD)/client/; gow run cmd/${APPNAME}_client/main.go
else
	$(info $(M) Starting an instance of $(APPNAME)_client)
	@cd $(PWD)/client/; go run cmd/${APPNAME}_client/main.go
endif

.PHONY: setup
setup: ## Start dependencies setup
	$(info $(M) Setup Client dependencies)
	@cd $(PWD)/client/; \
	    rm -rf go.mod go.sum vendor; \
		go mod init ${APPNAME}_client; \
		go mod tidy; \
		go mod vendor

	@echo "Completed"

	$(info $(M) Setup Server dependencies)	
	$(info $(M) TODO)	