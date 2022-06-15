FROM debian:stable-slim
COPY install.sh .
RUN ./install.sh
# Set up user
RUN useradd -m -p NbqDBxZy0F.tE -s /bin/bash jdoe
RUN usermod -aG sudo jdoe
RUN mkdir /home/jdoe/data
COPY p1.fa /home/jdoe/data
COPY README /home/jdoe/
COPY furTut.sh /home/jdoe/
COPY fur.pdf /home/jdoe
RUN chown -R jdoe /home/jdoe/
RUN chgrp -R jdoe /home/jdoe/
USER jdoe
ENV HOME /home/jdoe
WORKDIR /home/jdoe
