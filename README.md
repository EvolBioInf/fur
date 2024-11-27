# [`fur`](https://owncloud.gwdg.de/index.php/s/6PpPRqNzLpSbtGq): Find Unique Regions
## Description
The program `fur` takes as input a set of target genome sequences and
a set of related genome sequences, the neighbors. It returns the
sequence regions common to all targets that are absent form the
neighbors. Such regions can be used as candidate genetic markers.
## Authors
[Bernhard Haubold](http://guanine.evolbio.mpg.de/)
and [Beatriz Vieira Mourato](https://beatrizvm.github.io/) 
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

## References

Haubold, B., Klötzl, F., Hellberg, L., Thompson, D., & Cavalar, M. (2021). [Fur: Find unique genomic regions for diagnostic PCR](https://academic.oup.com/bioinformatics/article/37/15/2081/6124302). *Bioinformatics*, 37(15), 2081-2087.

Vieira Mourato, B., Tsers, I., Denker, S., Klötzl, F., & Haubold, B. (2024). [Marker discovery in the large](https://academic.oup.com/bioinformaticsadvances/article/4/1/vbae113/7721998). *Bioinformatics Advances*, 4(1), vbae113.
