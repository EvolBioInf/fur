FROM ubuntu
RUN apt update --fix-missing
RUN yes | apt install libdivsufsort3 libgsl23 libgslcblas0 libgomp1 libbsd0 wget sudo
RUN useradd -m -p NbqDBxZy0F.tE -s /bin/bash jdoe
RUN usermod -aG sudo jdoe
COPY bin/* /usr/local/bin/
RUN chmod +x /usr/local/bin/*
COPY README /home/jdoe/
RUN chown -R jdoe /home/jdoe/README
RUN chgrp -R jdoe /home/jdoe/README
USER jdoe
ENV HOME /home/jdoe
WORKDIR /home/jdoe

