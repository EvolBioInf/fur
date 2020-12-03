# `fur`
## Author
Bernhard Haubold, `haubold@evolbio.mpg.de`
## Description
Find unique genomic regions.
## Dependencies
In order to build `fur` from its sources, a number of packages need to
be installed on your system, in addition to the programs `macle` and
`phylonium`.
* Install packages
`sudo apt install install autoconf build-essential curl git gnuplot libbsd-dev \
    libbsd0 libdivsufsort-dev libdivsufsort3 libgsl-dev libgsl23 \
    libsdsl-dev libsdsl3 ncbi-blast+ noweb primer3 sudo \
    texlive-latex-extra texlive-latex-recommended texlive-pstricks \
    texlive-science`
* Install [`macle`](http://github.com/evolbioinf/macle)
* Install [`phylonium`](http://github.com/evolbioinf/phylonium)
## Compile
* Compile the `fur` sources using `make`; all executables are now in
  the directory `build`
## Documentation
The command `make doc` generates the manual `doc/fur.pdf`.
## Docker Container 
As an alternative to building `fur` yourself, we post the docker
  container `haubold/fox` on dockerhub. The container includes the
  documentation and all programs needed to work through the tutorial
  at the end. Once you have worked through the tutorial, you can start
  tackling your own analyses using the container.
  -  `docker pull haubold/fox`
  -  `docker container run --detach-keys="ctrl-@" -h fox -it haubold/fox`
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
