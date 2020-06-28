FROM ubuntu
RUN apt update --fix-missing
RUN yes | apt install curl libbsd0 libdivsufsort3 libgomp1 libgsl23 libgslcblas0 ncbi-blast+ primer3 sudo
RUN useradd -m -p NbqDBxZy0F.tE -s /bin/bash jdoe
RUN usermod -aG sudo jdoe
COPY bin/* /usr/local/bin/
RUN chmod +x /usr/local/bin/*
RUN mkdir /home/jdoe/data
COPY p1.fa /home/jdoe/data
COPY README /home/jdoe/
RUN chown -R jdoe /home/jdoe/
RUN chgrp -R jdoe /home/jdoe/
USER jdoe
ENV HOME /home/jdoe
WORKDIR /home/jdoe

