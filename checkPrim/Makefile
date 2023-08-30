exe = checkPrim
all: $(exe)
$(exe): $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe) > $(exe)
	chmod +x $(exe)
clean:
	rm $(exe)
