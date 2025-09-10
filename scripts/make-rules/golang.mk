GO_LDFLAGS := $(shell ./scripts/ldflags.sh)

BUILD_TARGETS := $(notdir $(wildcard ./cmd/*))
BUILD_EXAMPLE_TARGETS := $(notdir $(wildcard ./example/*))

go.build:
	@echo "==========> Building targets: $(BUILD_TARGETS)"
	@mkdir -p ./_output/bin
	@for target in $(BUILD_TARGETS); do \
		echo "==========> Building target: $$target"; \
		go build -ldflags "$(GO_LDFLAGS)" -o ./_output/bin/$$target ./cmd/$$target; \
	done

go.build.%:
	@echo "==========> Building target: $*"
	@go build -ldflags "$(GO_LDFLAGS)" -o ./_output/bin/$* ./cmd/$*
	@echo "==========> Done"

go.build.example.%:
	@echo "==========> Building target: $*"
	@go build -ldflags "$(GO_LDFLAGS)" -o ./_output/bin/$* ./example/$*
	@echo "==========> Done"


go.tidy:
	@echo "==========> Running go mod tidy..."
	@go mod tidy

.PHONY: go.build go.tidy