#!/bin/bash
# Bash "strict mode"
set -euo pipefail
# Noninteractive
export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get -y upgrade
apt-get -y install autoconf build-essential curl git gnuplot golang libbsd-dev \
    libbsd0 libdivsufsort-dev libdivsufsort3 libgsl-dev libgsl23 \
    libsdsl-dev libsdsl3 ncbi-blast+ noweb primer3 sudo \
    texlive-latex-extra texlive-latex-recommended texlive-pstricks \
    texlive-science
apt-get clean
rm -rf /var/lib/apt/lists/*
# Install macle
git clone https://github.com/evolbioinf/macle
cd macle
make
cp build/macle /usr/local/bin
cd ..
rm -rf macle
# Install phylonium
git clone https://github.com/evolbioinf/phylonium
cd phylonium
autoreconf -fi -Im4
./configure
make
cp src/phylonium /usr/local/bin
cd ..
rm -rf phylonium
# Install and test fur
git clone https://github.com/evolbioinf/fur
cd fur
make
cp build/* /usr/local/bin
make doc
cp doc/fur.pdf /usr/local/share/
cd ..
rm -rf fur
# Clean up
apt-get -y remove autoconf build-essential git gnuplot golang libbsd-dev \
    libdivsufsort-dev libgsl-dev libsdsl-dev noweb \
    texlive-latex-extra texlive-latex-recommended texlive-pstricks \
    texlive-science
apt clean
apt -y autoremove
