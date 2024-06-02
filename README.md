# [`fur`](https://owncloud.gwdg.de/index.php/s/ZJrSZ10O97fAV2j): Find Unique Regions
## Description
The program `fur` takes as input a set of target genome sequences and
a set of related genome sequences, the neighbors. It returns the
sequence regions common to all targets that are absent form the
neighbors. Such regions can be used as candidate genetic markers.
## Author
[Bernhard Haubold](http://guanine.evolbio.mpg.de/), `haubold@evolbio.mpg.de`
## Make the Programs
If you are on an Ubuntu system like Ubuntu on
[wsl](https://learn.microsoft.com/en-us/windows/wsl/install) under
MS-Windows or the [Ubuntu Docker
container](https://hub.docker.com/_/ubuntu), you can clone the
repository and change into it.

`git clone https://github.com/evolbioinf/fur`  
`cd fur`

Then install the additional dependencies by running the script
[`setup.sh`](scripts/setup.sh).

`bash scripts/setup.sh`

Make the programs.

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
