git describe | sed -E 's/^[vV]//; s/(.)$/\1\\/; s/-g(.......)/\\ \\(\1\\)/' | tr -d '\n'
