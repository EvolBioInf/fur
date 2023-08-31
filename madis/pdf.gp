set terminal postscript eps color size 10cm,8cm
set output "pdf.ps"
set xlabel "Match Length"
set ylabel "P(Match Length)"
set arrow from 9,0 to 9,0.5 nohead
set arrow from 12,0 to 12,0.5 nohead
plot[*:*][*:*] "-" t "exp" w l, "-" t "obs" w l
8	0.000483005
9	0.148372
10	0.471992
11	0.266851
12	0.0830738
13	0.0218282
14	0.00561806
15	0.00132701
16	0.000335003
17	8.60009e-05
18	2.40002e-05
19	5.00005e-06
20	3.00003e-06
21	1.00001e-06
22	1.00001e-06
e
5	8.246265736869019e-213
6	9.636491400545489e-54
7	5.575490769680773e-14
8	0.00048594792825268216
9	0.14798744005634895
10	0.47227018080050875
11	0.2668784393711048
12	0.08301537949496296
13	0.021939718468916847
14	0.005561982363703644
15	0.0013953586287646091
16	0.00034914434068800126
17	8.730513953270069e-05
18	2.1827475962576948e-05
19	5.456943435566686e-06
20	1.3642405116698342e-06
21	3.4106042023918093e-07
22	8.526512229600769e-08
23	2.131627996337926e-08
24	5.329072849669103e-09
25	1.3322670744386755e-09
26	4.4408943189466754e-10