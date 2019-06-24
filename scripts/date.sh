git log |
    grep Date |
    head -n 1 |
    awk '{printf "%s\\ %s\\ %s\\ %s,\\ %s\\\n", $2, $3, $4, $6, $5}'
