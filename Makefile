include VERSION

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
REPOPATH = github.com/miltalex/Mikasa

.PHONY: vendor

vendor:
	@echo -e "\033[1;33mgo mod vendor\033[0m"
	@go mod vendor