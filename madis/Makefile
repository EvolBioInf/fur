date = $(shell bash ../scripts/getDate.sh)
version = $(shell bash ../scripts/getVersion.sh)

exe = madis
vf = -X github.com/evolbioinf/fur/util.version=$(version)
df = -X github.com/evolbioinf/fur/util.date=$(date)
nw = $(shell which noweb)

all : $(exe)
$(exe): $(exe).go
	go build -ldflags "$(vf) $(df)" $(exe).go
tangle: $(exe).go $(exe)_test.go
$(exe).go: $(exe).org
	if [ $(nw) != "" ]; then\
		awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe).go > $(exe).go;\
	fi
$(exe)_test.go:
	if [ $(nw) != "" ]; then\
		awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go;\
	fi
test: $(exe) $(exe)_test.go
	@go test
clean:
	rm -f $(exe) *.go

