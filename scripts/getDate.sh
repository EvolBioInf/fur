date="no_date"
g=$(which git)
if [[ $g != "" ]]; then
	date=$(git log | grep Date | head -n 1 | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
fi
echo $date
