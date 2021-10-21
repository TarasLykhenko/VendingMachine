GOTEST_OUTPUT_RESULTS=testresults.out
GOTEST_OUTPUT_COVERAGE=testcoverageresults.out

# Check where are we running the build
ifdef BUILD_BUILDNUMBER
    # CI Build
	GOTEST_OUTPUT=-coverprofile=$(GOTEST_OUTPUT_COVERAGE) > $(GOTEST_OUTPUT_RESULTS)
else
    # Manual build
	GOTEST_OUTPUT=-cover
endif

.DEFAULT_GOAL := default

.PHONY: all

# test.lint:
# 	go install golang.org/x/lint/golint

# 	golint -set_exit_status ./...

test.vet:
	go vet ./...

test.unit:
	go test -v ./... $(GOTEST_OUTPUT)

# test.generate.report:
# 	go install github.com/jstemmer/go-junit-report
# 	go install github.com/axw/gocov/gocov
# 	go install github.com/AlekSi/gocov-xml

# 	gocov convert $(GOTEST_OUTPUT_COVERAGE) | gocov-xml > coverage.xml
# 	cat $(GOTEST_OUTPUT_RESULTS) | go-junit-report > report.xml


build.vendingmachine:
	$(MAKE) -C ./cmd/vendingmachine

build: build.vendingmachine

# default: test.lint build test.vet test.unit

default: build test.vet test.unit