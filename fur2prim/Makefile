exe = fur2prim
all: $(exe)
$(exe): $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe) > $(exe)
	chmod +x $(exe)
clean:
	rm $(exe)
test:
	@echo "The program $(exe) is not tested."
