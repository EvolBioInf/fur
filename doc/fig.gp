set terminal postscript eps
set size 0.6, 0.6
set xlabel "C_m"
set ylabel "CDF"

f(x) = 0.95
set arrow from 1.0197856732271406,0 to 1.0197856732271406,1 nohead
set output "cdf.eps"
plot [][] "cdf.dat" title "" wi li lw 5,\
f(x) title "" wi li ls 1
