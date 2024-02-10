# Get docker
FROM debian:stable-slim
# Get packages
RUN apt -y upgrade && apt -y update && apt install -y gnuplot libdivsufsort-dev libgsl-dev ncbi-blast+ phylonium sudo tar time wget
# Set up user
RUN useradd -m -p NbqDBxZy0F.tE -s /bin/bash jdoe
RUN usermod -aG sudo jdoe
COPY README furDoc.pdf /home/jdoe/
COPY bin/* /usr/local/bin/
RUN chown root:root /usr/bin/sudo && chmod 4755 /usr/bin/sudo && chown -R jdoe /home/jdoe/ && chgrp -R jdoe /home/jdoe/
USER jdoe
WORKDIR /home/jdoe
ENV HOME /home/jdoe

