version = $(shell cat ../data/version.txt)
date = $(shell cat ../data/date.txt)

exe = makeFurDb
all : $(exe)

$(exe): $(exe).go
	go build -ldflags "-X github.com/evolbioinf/fur/util.version=$(version) -X github.com/evolbioinf/fur/util.date=$(date)" $(exe).go
$(exe).go: $(exe).org
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe).go | gofmt > $(exe).go
test: $(exe)_test.go
	@go test
$(exe)_test.go: $(exe).org
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go	
clean:
	rm -f *.go

