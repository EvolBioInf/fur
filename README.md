# `fur`
## Author
Bernhard Haubold, `haubold@evolbio.mpg.de`
## Description
Find unique genomic regions.
## Dependencies
* `gnuplot`
* gsl (Gnu Scientific Library)
* `latex`
* libbsd
* [`macle`](https://github.com/evolbioinf/macle)
* ncbi-blast+
* `noweb`
* [`phylonium`](https://github.com/evolbioinf/phylonium)
* pst-tools
* `wget`
## Compile
* Compile the sources using `make`; all executables are now in the
  directory `build`
* If this does not work on your system, consider using the "fox"
  docker container (for "fur box"). It contains an installation of all
  programs needed to work through the tutorial in the documentation.
  -  `docker pull haubold/fox`
  -  `docker container run --detach-keys="ctrl-@" -h fox -it haubold/fox`
## Documentation
The command `make doc` generates the manual `doc/fur.pdf`, or use
  the [typeset version](http://guanine.evolbio.mpg.de/fur/fur.pdf).
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
