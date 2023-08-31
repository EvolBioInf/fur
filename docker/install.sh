#!/bin/bash
# Bash "strict mode"
set -euo pipefail
# Noninteractive
export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get -y upgrade
apt-get -y install apt-utils build-essential curl git gnuplot golang libbsd-dev \
    libbsd0 libdivsufsort-dev libdivsufsort3 libgsl-dev libgslcblas0 \
    libsdsl-dev libsdsl3 ncbi-blast+ noweb phylonium primer3 sudo \
    texlive-latex-extra texlive-latex-recommended texlive-pstricks \
    texlive-science
apt-get clean
rm -rf /var/lib/apt/lists/*
# Install fur
git clone https://github.com/evolbioinf/fur
cd fur
make
cp bin/* /usr/local/bin
# make doc
# cp doc/fur.pdf /usr/local/share/
cd ..
rm -rf fur
# Clean up
apt-get -y remove autoconf build-essential git gnuplot golang noweb \
    texlive-latex-extra texlive-latex-recommended texlive-pstricks \
    texlive-science
apt clean
apt -y autoremove
