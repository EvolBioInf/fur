# Convert each FASTA entry to two lines
sed 's/>.*$/>/' |
    tr -d '\n' |
    sed 's/>/\n>\n/g' |
    # Remove runs of Ns at the beginning of lines
    sed -E 's/^N+//' |
    # Remove runs of Ns at the end of lines
    sed -E 's/N*$//' |
    # Break sequences at internal runs of N
    sed -E 's/N+/\n>S\n/g' |
    # Print sequences at least 50 nucleotides long
    awk '!/^>/ {s[n++] = $1} END {for (i = 0; i < n; i++) if (length(s[i]) >= 50) printf ">S%d\n%s\n", ++c, s[i]}'
