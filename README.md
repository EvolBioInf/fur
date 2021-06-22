# `fur`
## Author
Bernhard Haubold, `haubold@evolbio.mpg.de`
## Description
Find unique genomic regions.
## Dependencies
Building `fur` requires a number of packages and the two programs
`macle` & `phylonium`.
* Install packages:
`sudo apt install autoconf build-essential curl git gnuplot golang libbsd-dev 
    libbsd0 libdivsufsort-dev libdivsufsort3 libgsl-dev libgsl23 
    libsdsl-dev libsdsl3 ncbi-blast+ noweb primer3 sudo 
    texlive-latex-extra texlive-latex-recommended texlive-pstricks 
    texlive-science`
* [Install `macle`](http://github.com/evolbioinf/macle)
* [Install `phylonium`](http://github.com/evolbioinf/phylonium)
## Compile
* Compile `fur` with the command `make`; all executables are now in
  the directory `build`
## Documentation
The command `make doc` generates the manual `doc/fur.pdf`.
## Docker Container 
As an alternative to building `fur`, we also post it as a [docker
  container](https://hub.docker.com/r/haubold/fox). The container
  includes all programs needed to work through the tutorial at the end
  of the documentation in `~/fox.pdf`.
  -  `docker pull haubold/fox`
  -  `docker container run --detach-keys="ctrl-@" -h fox -it haubold/fox`
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
