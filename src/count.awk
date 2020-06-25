{
  if (/^>/)
    c++
  else {
    t += length($1)
    n += gsub("N", "")
  }
}
END {
  printf "# %s tmpl, %d nuc, %d N, %d total\n", c, t - n, n, t
}
