date = $(shell git log | grep Date | head -n 1 | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
version = $(shell git describe)

exe = prim2tab
all : $(exe)

$(exe): $(exe).go
	go build -ldflags "-X github.com/evolbioinf/fur/util.version=$(version) -X github.com/evolbioinf/fur/util.date=$(date)" $(exe).go
$(exe).go: $(exe).org
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe).go | gofmt > $(exe).go
test: $(exe)_test.go
	@go test
$(exe)_test.go: $(exe).go
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go	
clean:
	rm -f *.go

