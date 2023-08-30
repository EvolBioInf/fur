/^ *!/ {
  l = $0
  sub(/!/, "", l)
  printf "\\textbf{%s}\n", l
}
!/^ *!/ {
  print
}
