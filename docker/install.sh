#!/bin/bash
# Bash "strict mode"
set -euo pipefail
# Noninteractive
export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get -y upgrade
apt-get -y install apt-utils curl libdivsufsort3 libgsl-dev libgslcblas0 \
    ncbi-blast+ phylonium primer3 sudo
apt-get clean
apt -y autoremove
rm -rf /var/lib/apt/lists/*
