version="no_version"
g=$(which git)
if [[ $g != "" ]]; then
    version=$(git describe)
fi
echo $version
