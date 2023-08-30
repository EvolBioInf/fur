wget guanine.evolbio.mpg.de/fur/eco105.tar.gz
tar -xvzf eco105.tar.gz
makeFurDb -t targets -n neighbors -d furDb
fur -d furDb > tmpl.fasta
# Step                    Sequences  Nucleotides  Mutations (N)
# -------------------------------------------------------------
# Sliding window               1005       681264              0
# Presence in targets           170        69407            151
# Absence from neighbors         91        46844           4309
head tmpl.fasta
count tmpl.fasta 
# 91 tmpl, 46844 nuc, 4309 N, 51153 total
fur -d furDb -u > unique1.fasta
# Step                    Sequences  Nucleotides  Mutations (N)
# -------------------------------------------------------------
# Sliding window               1005       681264              0
count unique1.fasta
# 1005 tmpl, 681264 nuc, 0 N, 681264 total
fur -d furDb -U > unique2.fasta
# Step                    Sequences  Nucleotides  Mutations (N)
# -------------------------------------------------------------
# Sliding window               1005       681264              0
# Presence in targets           170        69407            151
count unique2.fasta
# 170 tmpl, 69407 nuc, 151 N, 69558 total
fur -d furDb -w 1000 > tmpl.fasta
# Step                    Sequences  Nucleotides  Mutations (N)
# -------------------------------------------------------------
# Sliding window                111       634900              0
# Presence in targets            26        28730            108
# Absence from neighbors         18        18027           2990
fur -d furDb -w 90 > tmpl.fasta
# Step                    Sequences  Nucleotides  Mutations (N)
# -------------------------------------------------------------
# Sliding window               1610       609930              0
# Presence in targets           246        72184            167
# Absence from neighbors        174        53570           2922
fur -d furDb -m > tmpl.fasta
# Step                    Sequences  Nucleotides  Mutations (N)
# -------------------------------------------------------------
# Sliding window               1005       681264              0
# Presence in targets           170        69407            151
# Absence from neighbors        168        69052            137
fur -d furDb -e 1e-20 > tmpl.fasta
# Step                    Sequences  Nucleotides  Mutations (N)
# -------------------------------------------------------------
# Sliding window               1005       681264              0
# Presence in targets           170        69407            151
# Absence from neighbors        102        50516           2573
cleanSeq tmpl.fasta > tmpl2.fasta
senSpec -v query=tmpl2.fasta -v db=furDb
#S_n    S_p     C
1.000   0.995   0.962
fur2prim tmpl.fasta > prim.txt
primer3_core prim.txt > prim.out
prim2fasta -v file=primer prim.out 
grep PRIMER_PAIR_0_PENALTY prim.out |
    tr '=' '\t' |
    sort -k 2 -n |
    head -n 1
