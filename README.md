# `fur`
## Description
Iterated [`fur`](https://github.com/evolbioinf/fur).
## Author
[Bernhard Haubold](http://guanine.evolbio.mpg.de/), `haubold@evolbio.mpg.de`
## Make the Programs
Make sure you've installed the packages `git`, `golang`, `make`, and `noweb`.  
  `$ make`  
  The directory `bin` now contains the binaries, scripts are in
  `scripts`.
## Docker Container
As an alternative to building `fur`, we also post it as a [docker
  container](https://hub.docker.com/r/haubold/fox). The container
  includes all programs needed to work through the tutorial at the end
  of the documentation in `~/fox.pdf`.
  -  `$ docker pull haubold/fox`
  -  `$ docker container run --detach-keys="ctrl-@" -h fox -it haubold/fox`
## Make the Documentation
Make sure you've installed the packages `git`, `make`, `noweb`, `texlive-science`,
`texlive-pstricks`, `texlive-latex-extra`,
and `texlive-fonts-extra`.  
  `$ make doc`  
  The documentation is now in `doc/ifurDoc.pdf`.
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
