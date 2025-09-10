.DEFAULT := all

all: tidy build

include ./scripts/make-rules/golang.mk

build: go.tidy
	@make go.build

tidy:
	@make go.tidy

.PHONY: all build tidy