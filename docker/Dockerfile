FROM debian:stable-slim
COPY install.sh .
RUN ./install.sh
# Set up user
RUN useradd -m -p NbqDBxZy0F.tE -s /bin/bash jdoe
RUN usermod -aG sudo jdoe
RUN mkdir /home/jdoe/data
COPY p1.fa /home/jdoe/data
COPY README /home/jdoe/
COPY furDoc.pdf /home/jdoe
COPY bin/* /usr/local/bin
RUN chown root:root /usr/bin/sudo && chmod 4755 /usr/bin/sudo
RUN chown -R jdoe /home/jdoe/
RUN chgrp -R jdoe /home/jdoe/
USER jdoe
ENV HOME /home/jdoe
WORKDIR /home/jdoe
