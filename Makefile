ROLE := infer-vllm
.PHONY: build test
build:
	go build -o bin/cofiswarm-infer-vllm ./cmd/cofiswarm-infer-vllm
test: build test-standalone-layout
test-standalone-layout:
	./test/scripts/assert-layout.sh $(ROLE)
