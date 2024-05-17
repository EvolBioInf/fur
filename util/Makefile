all : util

util:
	go build util.go
tangle: util.go
	go build util.go
util.go: util.org
	awk -f ../scripts/preTangle.awk util.org | bash ../scripts/org2nw | notangle -Rutil.go | gofmt > util.go
.PHONY: test
test: 
	@echo "The package util is not tested."
util_test.go:
	awk -f ../scripts/preTangle.awk util.org | bash ../scripts/org2nw | notangle -Rutil_test.go | gofmt > util_test.go

clean:
	rm -f *.go
