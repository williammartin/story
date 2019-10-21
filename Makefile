test:
	ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race

test-unit:
	ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race --skipPackage acceptance

test-acceptance:
	ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race cmd/acceptance

build:
	go build -o story cmd/story/main.go

.PHONY: test test-unit test-acceptance build
