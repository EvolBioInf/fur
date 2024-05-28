# [`fur`](https://owncloud.gwdg.de/index.php/s/ZJrSZ10O97fAV2j): Find Unique Regions
## Description
The program `fur` takes as input a set of target genome sequences and
a set of related genome sequences, the neighbors. It returns the
sequence regions common to all targets that are absent form the
neighbors. Such regions can be used as candidate genetic markers.
## Author
[Bernhard Haubold](http://guanine.evolbio.mpg.de/), `haubold@evolbio.mpg.de`
## Make the Programs on Ubuntu
Setup the environment by running the [setup script](scripts/setup.sh).

`bash scripts/setup.sh`

Then make the programs.

`make`

The directory `bin` now contains the binaries.
## Docker Container
As an alternative to building `fur`, we also post it as a [docker
  container](https://hub.docker.com/r/haubold/fox). The container
  includes all programs needed to work through the tutorial at the end
  of the [documentation](https://owncloud.gwdg.de/index.php/s/ZJrSZ10O97fAV2j) in `~/furDoc.pdf`.
  
`docker pull haubold/fox`  
`docker run -it --detach-keys="ctrl-@" fox`
  
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
