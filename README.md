# `fur`: Find Unique Regions
## Description
The program `fur` takes as input a set of target genome sequences and
a set of related genome sequences, the neighbors. It returns the
sequence regions common to all targets that are absent form the
neighbors. Such regions can be used as candidate genetic markers.
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
  of the documentation in `~/furDoc.pdf`.
  -  `$ docker pull haubold/fox`
  -  `$ docker run -it --env="DISPLAY" --net=host -v ~/fox_share:/home/jdoe/fox_share --detach-keys="ctrl-@" fox`  
  This constructs the directory `fox_share` in your home directory and
  in the container's home directory, for sharing files between the two
  environments.
## Make the Documentation
Make sure you've installed the packages `git`, `make`, `noweb`, `texlive-science`,
`texlive-pstricks`, `texlive-latex-extra`,
and `texlive-fonts-extra`.  
  `$ make doc`  
  The documentation is now in `doc/furDoc.pdf`.
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
