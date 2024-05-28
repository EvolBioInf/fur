date = $(shell git log | grep Date | head -n 1 | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
version = $(shell git describe)

exe = madis
vf = -X github.com/evolbioinf/fur/util.version=$(version)
df = -X github.com/evolbioinf/fur/util.date=$(date)

all : $(exe)
$(exe):
	go build -ldflags "$(vf) $(df)" $(exe).go
tangle: $(exe).go $(exe)_test.go
	go build -ldflags "$(vf) $(df)" $(exe).go
$(exe).go: $(exe).org
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe).go > $(exe).go
$(exe)_test.go: $(exe).go
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go	
test: $(exe)
	@go test
clean:
	rm -f $(exe) *.go

