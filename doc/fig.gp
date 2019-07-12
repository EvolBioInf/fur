#set terminal epslatex color solid
set terminal epslatex monochrome
set size 5/5., 4/3.
set format xy "\\large$%g$"
#set format y "\Large $%.0t\times 10^{%T}$"
set xlabel "\\Large$C_{\\rm m}$"
set ylabel "\\Large$\\mbox{CDF}$"
#set format "\Large$%g$"
#set logscale xy
#set pointsize 2

f(x) = 0.95
set arrow from 1.0197856732271406,0 to 1.0197856732271406,1 nohead
set output "cdf.tex"
plot [][] "cdf.dat" title "" wi li lw 5,\
f(x) title "" wi li ls 1
