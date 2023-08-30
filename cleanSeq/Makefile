exe = cleanSeq
all: $(exe)
$(exe): $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go > $(exe).go
	go build $(exe).go
clean:
	rm $(exe)
