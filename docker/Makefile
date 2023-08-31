all:
	cd ../ && make && cd - && cp -r ../bin .
	cd ../doc && make && cd - && cp ../doc/furDoc.pdf .
	sudo docker build -t fox .
