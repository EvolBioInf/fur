# `fur`
## Author
Bernhard Haubold, `haubold@evolbio.mpg.de`
## Description
Find unique genomic regions.
## Dependencies
Building `fur` requires a number of packages, the Go compiler, and the two programs
`macle` & `phylonium`.
* Install packages  
`$ sudo apt install autoconf build-essential cmake git gnuplot libbsd-dev 
    libbsd0 libdivsufsort-dev libdivsufsort3 libgsl-dev libgsl23 
    libsdsl-dev libsdsl3 ncbi-blast+ noweb primer3 sudo 
    texlive-latex-extra texlive-latex-recommended texlive-pstricks 
    texlive-fonts-extra texlive-science wget`
* Install Go compiler
  - Download package  
  `$ curl https://go.dev/dl/go1.17.6.linux-amd64.tar.gz -o go1.17.6.linux-amd64.tar.gz`
  - Remove old installation  
  `$ sudo rm -rf /usr/local/go`
  - Unpack and install  
  `$ sudo tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz`
  - Add `/usr/local/bo/bin` to the PATH environment variable by adding,
  for example, `export PATH=$PATH:/usr/local/bo/bin` to
  `$HOME/.profile`and update the system by executing  
  `$ source $HOME/.profile`
  - Verify Go installation  
  `$ go version`
* [Install `macle`](http://github.com/evolbioinf/macle)
* [Install `phylonium`](http://github.com/evolbioinf/phylonium)
## Compile
* Compile `fur`  
  `$ make`  
  All executables are now in the directory `build`
## Documentation
The command  
`$ make doc`  
generates the manual `doc/fur.pdf`
## Docker Container 
As an alternative to building `fur`, we also post it as a [docker
  container](https://hub.docker.com/r/haubold/fox). The container
  includes all programs needed to work through the tutorial at the end
  of the documentation in `~/fox.pdf`.
  -  `$ docker pull haubold/fox`
  -  `$ docker container run --detach-keys="ctrl-@" -h fox -it haubold/fox`
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
