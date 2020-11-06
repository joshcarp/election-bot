package main

import (
	"fmt"
	"regexp"
	"strings"
)
const text = `----------------------  --------------------------------------------------------------------------------
Last updated:           2020-11-06 03:59 UTC
Latest batch received:  2020-11-06 03:45 UTC
Prettier web version:   https://alex.github.io/nyt-2020-election-scraper/battleground-state-changes.html
----------------------  --------------------------------------------------------------------------------

Alaska (EV: 3) Total Votes: (Trump: 80,999, Biden: 45,758)
----------------  ---  -----------------------------  -------------------------  -------------------------------------------  ------------------  ----------------------------  -------------------------
2020-11-04 18:40  ---  Trump leading by 51,382 votes  Remaining (est.): 193,975  Change:  39,556 (Trump 70.4% / 29.6% Biden)  Precincts: 360/441  Biden needs 63.24% [5.699%]   Biden recent trend 29.60%
2020-11-04 13:28  ---  Trump leading by 35,241 votes  Remaining (est.): 233,531  Change:       0 (n/a)                        Precincts: 201/441  Biden needs 57.55% [57.545%]  Biden recent trend n/a
----------------  ---  -----------------------------  -------------------------  -------------------------------------------  ------------------  ----------------------------  -------------------------

Arizona (EV: 11) Total Votes: (Biden: 1,528,319, Trump: 1,482,062)
----------------  ---  ------------------------------  ---------------------------  -------------------------------------------  --------------------  ----------------------------  -------------------------
2020-11-06 03:37  ***  Biden leading by 46,667 votes   Remaining (est.): 314,747    Change:   1,180 (Biden 67.4% / 32.6% Trump)  Precincts: 1417/1489  Trump needs 57.41% [0.093%]   Trump recent trend 56.64%
2020-11-06 02:07  ---  Biden leading by 46,257 votes   Remaining (est.): 315,927    Change:  75,314 (Biden 43.0% / 57.0% Trump)  Precincts: 1415/1489  Trump needs 57.32% [0.058%]   Trump recent trend 57.02%
2020-11-06 01:44  ---  Biden leading by 56,833 votes   Remaining (est.): 391,241    Change:       0 (n/a)                        Precincts: 1386/1489  Trump needs 57.26% [-0.293%]  Trump recent trend 51.76%
2020-11-06 01:33  ---  Biden leading by 56,833 votes   Remaining (est.): 376,072    Change:   1,957 (Biden 24.2% / 75.8% Trump)  Precincts: 1386/1489  Trump needs 57.56% [-0.095%]  Trump recent trend 51.76%
2020-11-06 01:15  ---  Biden leading by 57,844 votes   Remaining (est.): 378,029    Change:   2,485 (Biden 76.1% / 23.9% Trump)  Precincts: 1383/1489  Trump needs 57.65% [0.220%]   Trump recent trend 50.23%
2020-11-06 01:04  ---  Biden leading by 56,547 votes   Remaining (est.): 380,514    Change:  28,344 (Biden 47.5% / 52.5% Trump)  Precincts: 1380/1489  Trump needs 57.43% [0.339%]   Trump recent trend 53.50%
2020-11-06 00:55  ---  Biden leading by 57,986 votes   Remaining (est.): 408,858    Change:   5,990 (Biden 41.9% / 58.1% Trump)  Precincts: 1366/1489  Trump needs 57.09% [-0.014%]  Trump recent trend 66.08%
2020-11-06 00:46  ---  Biden leading by 58,953 votes   Remaining (est.): 414,848    Change:   5,169 (Biden 27.5% / 72.5% Trump)  Precincts: 1362/1489  Trump needs 57.11% [-0.190%]  Trump recent trend 61.14%
2020-11-06 00:35  ---  Biden leading by 61,280 votes   Remaining (est.): 420,017    Change:   7,456 (Biden 23.9% / 76.1% Trump)  Precincts: 1361/1489  Trump needs 57.29% [-0.329%]  Trump recent trend 60.45%
2020-11-05 23:07  ---  Biden leading by 65,179 votes   Remaining (est.): 427,473    Change:  12,227 (Biden 38.8% / 61.2% Trump)  Precincts: 1359/1489  Trump needs 57.62% [-0.098%]  Trump recent trend 58.95%
2020-11-05 22:27  ---  Biden leading by 67,906 votes   Remaining (est.): 439,700    Change:       0 (n/a)                        Precincts: 1352/1489  Trump needs 57.72% [0.172%]   Trump recent trend 58.54%
2020-11-05 20:55  ---  Biden leading by 67,906 votes   Remaining (est.): 449,700    Change:   1,185 (Biden 40.6% / 59.4% Trump)  Precincts: 1352/1489  Trump needs 57.55% [-0.005%]  Trump recent trend 58.54%
2020-11-05 19:53  ---  Biden leading by 68,129 votes   Remaining (est.): 450,885    Change:     184 (Biden 20.1% / 79.9% Trump)  Precincts: 1352/1489  Trump needs 57.56% [-0.009%]  Trump recent trend 58.53%
2020-11-05 19:49  ---  Biden leading by 68,239 votes   Remaining (est.): 451,069    Change:     632 (Biden 23.4% / 76.6% Trump)  Precincts: 1351/1489  Trump needs 57.56% [-0.027%]  Trump recent trend 58.47%
2020-11-05 19:21  ---  Biden leading by 68,575 votes   Remaining (est.): 451,701    Change:   1,936 (Biden 54.8% / 45.2% Trump)  Precincts: 1350/1489  Trump needs 57.59% [0.053%]   Trump recent trend 58.29%
2020-11-05 07:45  ---  Biden leading by 68,390 votes   Remaining (est.): 453,637    Change:       0 (n/a)                        Precincts: 1349/1489  Trump needs 57.54% [0.000%]   Trump recent trend 58.69%
2020-11-05 07:43  ---  Biden leading by 68,390 votes   Remaining (est.): 453,637    Change:  62,002 (Biden 41.3% / 58.7% Trump)  Precincts: 1325/1489  Trump needs 57.54% [-0.139%]  Trump recent trend 58.69%
2020-11-05 02:53  ---  Biden leading by 79,173 votes   Remaining (est.): 515,639    Change:       0 (n/a)                        Precincts: 1325/1489  Trump needs 57.68% [0.096%]   Trump recent trend 59.03%
2020-11-05 02:19  ---  Biden leading by 79,173 votes   Remaining (est.): 522,139    Change:       0 (n/a)                        Precincts: 1325/1489  Trump needs 57.58% [-0.462%]  Trump recent trend 59.03%
2020-11-05 02:03  ---  Biden leading by 79,173 votes   Remaining (est.): 492,139    Change:  75,537 (Biden 41.0% / 59.0% Trump)  Precincts: 1325/1489  Trump needs 58.04% [-0.131%]  Trump recent trend 59.03%
2020-11-05 00:44  ---  Biden leading by 92,817 votes   Remaining (est.): 567,676    Change:   5,614 (Biden 48.2% / 51.8% Trump)  Precincts: 1296/1489  Trump needs 58.18% [0.063%]   Trump recent trend 64.17%
2020-11-04 23:49  ---  Biden leading by 93,016 votes   Remaining (est.): 573,290    Change:       0 (n/a)                        Precincts: 1294/1489  Trump needs 58.11% [3.635%]   Trump recent trend 64.71%
2020-11-04 23:41  ---  Biden leading by 93,016 votes   Remaining (est.): 1,038,753  Change:       0 (n/a)                        Precincts: 1294/1489  Trump needs 54.48% [-4.394%]  Trump recent trend 64.71%
2020-11-04 23:16  ---  Biden leading by 93,016 votes   Remaining (est.): 524,269    Change:       0 (n/a)                        Precincts: 1294/1489  Trump needs 58.87% [-3.065%]  Trump recent trend 64.71%
2020-11-04 20:28  ---  Biden leading by 93,016 votes   Remaining (est.): 389,660    Change:     785 (Biden 18.0% / 82.0% Trump)  Precincts: 1294/1489  Trump needs 61.94% [-0.040%]  Trump recent trend 64.71%
2020-11-04 19:44  ---  Biden leading by 93,518 votes   Remaining (est.): 390,445    Change:      66 (Biden 50.0% / 50.0% Trump)  Precincts: 1294/1489  Trump needs 61.98% [0.002%]   Trump recent trend 64.60%
2020-11-04 18:56  ---  Biden leading by 93,518 votes   Remaining (est.): 390,511    Change:     210 (Biden 52.1% / 47.9% Trump)  Precincts: 1294/1489  Trump needs 61.97% [0.008%]   Trump recent trend 64.61%
2020-11-04 15:28  ---  Biden leading by 93,509 votes   Remaining (est.): 390,721    Change:       0 (n/a)                        Precincts: 1294/1489  Trump needs 61.97% [0.178%]   Trump recent trend 64.64%
2020-11-04 14:47  ---  Biden leading by 93,509 votes   Remaining (est.): 396,611    Change: 126,904 (Biden 35.4% / 64.6% Trump)  Precincts: 1294/1489  Trump needs 61.79% [-0.691%]  Trump recent trend 64.64%
2020-11-04 13:28  ---  Biden leading by 130,665 votes  Remaining (est.): 523,515    Change:       0 (n/a)                        Precincts: 1243/1489  Trump needs 62.48% [62.480%]  Trump recent trend n/a
----------------  ---  ------------------------------  ---------------------------  -------------------------------------------  --------------------  ----------------------------  -------------------------

Georgia (EV: 16) Total Votes: (Trump: 2,447,337, Biden: 2,445,540)
----------------  ---  ------------------------------  -------------------------  --------------------------------------------  --------------------  ----------------------------  -------------------------
2020-11-06 03:29  ***  Trump leading by 1,775 votes    Remaining (est.): 10,534   Change:      34 (Trump 17.6% / 82.4% Biden)   Precincts: 2643/2655  Biden needs 58.43% [-0.077%]  Biden recent trend 66.52%
2020-11-06 03:06  ---  Trump leading by 1,797 votes    Remaining (est.): 10,568   Change:     337 (Trump 34.4% / 65.6% Biden)   Precincts: 2643/2655  Biden needs 58.50% [-0.219%]  Biden recent trend 66.50%
2020-11-06 02:33  ---  Trump leading by 1,902 votes    Remaining (est.): 10,905   Change:   1,031 (Trump 21.1% / 78.9% Biden)   Precincts: 2643/2655  Biden needs 58.72% [-1.739%]  Biden recent trend 66.51%
2020-11-06 02:00  ---  Trump leading by 2,497 votes    Remaining (est.): 11,936   Change:   1,337 (Trump 13.0% / 87.0% Biden)   Precincts: 2643/2655  Biden needs 60.46% [-2.672%]  Biden recent trend 66.11%
2020-11-06 01:44  ---  Trump leading by 3,486 votes    Remaining (est.): 13,273   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 63.13% [1.880%]   Biden recent trend 65.20%
2020-11-06 01:35  ---  Trump leading by 3,486 votes    Remaining (est.): 15,491   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 61.25% [-0.001%]  Biden recent trend 65.20%
2020-11-06 01:33  ---  Trump leading by 3,486 votes    Remaining (est.): 15,490   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 61.25% [2.994%]   Biden recent trend 65.20%
2020-11-06 01:27  ---  Trump leading by 3,486 votes    Remaining (est.): 21,106   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 58.26% [-0.007%]  Biden recent trend 65.20%
2020-11-06 01:26  ---  Trump leading by 3,486 votes    Remaining (est.): 21,087   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 58.27% [-1.175%]  Biden recent trend 65.20%
2020-11-06 01:26  ---  Trump leading by 3,486 votes    Remaining (est.): 18,462   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 59.44% [1.815%]   Biden recent trend 65.20%
2020-11-06 00:35  ---  Trump leading by 3,486 votes    Remaining (est.): 22,856   Change:   2,831 (Trump 47.4% / 52.6% Biden)   Precincts: 2643/2655  Biden needs 57.63% [0.550%]   Biden recent trend 65.20%
2020-11-06 00:01  ---  Trump leading by 3,635 votes    Remaining (est.): 25,687   Change:  17,356 (Trump 33.3% / 66.7% Biden)   Precincts: 2643/2655  Biden needs 57.08% [-3.874%]  Biden recent trend 67.26%
2020-11-05 23:07  ---  Trump leading by 9,426 votes    Remaining (est.): 43,043   Change:     342 (Trump 34.2% / 65.8% Biden)   Precincts: 2643/2655  Biden needs 60.95% [-0.038%]  Biden recent trend 69.74%
2020-11-05 22:45  ---  Trump leading by 9,534 votes    Remaining (est.): 43,385   Change:      19 (Trump 47.4% / 52.6% Biden)   Precincts: 2643/2655  Biden needs 60.99% [0.004%]   Biden recent trend 69.77%
2020-11-05 22:21  ---  Trump leading by 9,535 votes    Remaining (est.): 43,404   Change:     241 (Trump 52.1% / 47.9% Biden)   Precincts: 2643/2655  Biden needs 60.98% [0.072%]   Biden recent trend 69.78%
2020-11-05 22:01  ---  Trump leading by 9,525 votes    Remaining (est.): 43,645   Change:   6,305 (Trump 24.2% / 75.8% Biden)   Precincts: 2643/2655  Biden needs 60.91% [-1.874%]  Biden recent trend 69.94%
2020-11-05 21:53  ---  Trump leading by 12,773 votes   Remaining (est.): 49,950   Change:       5 (Trump 70.0% / 30.0% Biden)   Precincts: 2643/2655  Biden needs 62.79% [0.003%]   Biden recent trend 71.14%
2020-11-05 21:37  ---  Trump leading by 12,771 votes   Remaining (est.): 49,955   Change:       6 (Trump 100.0% /  0.0% Biden)  Precincts: 2643/2655  Biden needs 62.78% [0.008%]   Biden recent trend 71.15%
2020-11-05 21:13  ---  Trump leading by 12,765 votes   Remaining (est.): 49,961   Change:       8 (Trump 50.0% / 50.0% Biden)   Precincts: 2643/2655  Biden needs 62.77% [0.002%]   Biden recent trend 71.16%
2020-11-05 21:05  ---  Trump leading by 12,765 votes   Remaining (est.): 49,969   Change:       3 (n/a)                         Precincts: 2643/2655  Biden needs 62.77% [-0.002%]  Biden recent trend 71.16%
2020-11-05 20:55  ---  Trump leading by 12,768 votes   Remaining (est.): 49,972   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 62.78% [-2.159%]  Biden recent trend 71.16%
2020-11-05 20:40  ---  Trump leading by 12,768 votes   Remaining (est.): 42,747   Change:      48 (Trump 54.2% / 45.8% Biden)   Precincts: 2643/2655  Biden needs 64.93% [1.149%]   Biden recent trend 71.16%
2020-11-05 20:32  ---  Trump leading by 12,764 votes   Remaining (est.): 46,294   Change:   3,590 (Trump 49.0% / 51.0% Biden)   Precincts: 2643/2655  Biden needs 63.79% [-0.674%]  Biden recent trend 71.19%
2020-11-05 20:23  ---  Trump leading by 12,835 votes   Remaining (est.): 44,383   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 64.46% [-9.274%]  Biden recent trend 73.31%
2020-11-05 20:15  ---  Trump leading by 12,835 votes   Remaining (est.): 27,040   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 73.73% [-1.308%]  Biden recent trend 73.31%
2020-11-05 20:14  ---  Trump leading by 12,835 votes   Remaining (est.): 25,628   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 75.04% [3.185%]   Biden recent trend 73.31%
2020-11-05 20:05  ---  Trump leading by 12,835 votes   Remaining (est.): 29,363   Change:      19 (Trump 68.4% / 31.6% Biden)   Precincts: 2643/2655  Biden needs 71.86% [-3.366%]  Biden recent trend 73.31%
2020-11-05 20:01  ---  Trump leading by 12,828 votes   Remaining (est.): 25,430   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 75.22% [10.635%]  Biden recent trend 73.33%
2020-11-05 19:45  ---  Trump leading by 12,828 votes   Remaining (est.): 43,970   Change:       3 (Trump 100.0% /  0.0% Biden)  Precincts: 2643/2655  Biden needs 64.59% [0.004%]   Biden recent trend 73.33%
2020-11-05 19:37  ---  Trump leading by 12,825 votes   Remaining (est.): 43,973   Change:     586 (Trump 16.3% / 83.7% Biden)   Precincts: 2643/2655  Biden needs 64.58% [-0.251%]  Biden recent trend 73.34%
2020-11-05 19:21  ---  Trump leading by 13,220 votes   Remaining (est.): 44,559   Change:       1 (Trump 100.0% /  0.0% Biden)  Precincts: 2643/2655  Biden needs 64.83% [0.001%]   Biden recent trend 73.16%
2020-11-05 19:12  ---  Trump leading by 13,219 votes   Remaining (est.): 44,560   Change:     939 (Trump 33.2% / 66.8% Biden)   Precincts: 2643/2655  Biden needs 64.83% [-0.040%]  Biden recent trend 73.16%
2020-11-05 18:50  ---  Trump leading by 13,534 votes   Remaining (est.): 45,499   Change:      46 (Trump 43.5% / 56.5% Biden)   Precincts: 2643/2655  Biden needs 64.87% [0.008%]   Biden recent trend 73.34%
2020-11-05 17:52  ---  Trump leading by 13,540 votes   Remaining (est.): 45,545   Change:       1 (Trump 100.0% /  0.0% Biden)  Precincts: 2643/2655  Biden needs 64.86% [0.001%]   Biden recent trend 73.37%
2020-11-05 17:48  ---  Trump leading by 13,539 votes   Remaining (est.): 45,546   Change:     796 (Trump 14.8% / 85.2% Biden)   Precincts: 2643/2655  Biden needs 64.86% [-0.350%]  Biden recent trend 73.37%
2020-11-05 17:16  ---  Trump leading by 14,100 votes   Remaining (est.): 46,342   Change:   3,805 (Trump 41.3% / 58.7% Biden)   Precincts: 2643/2655  Biden needs 65.21% [0.491%]   Biden recent trend 73.07%
2020-11-05 16:46  ---  Trump leading by 14,765 votes   Remaining (est.): 50,147   Change:     219 (Trump 29.0% / 71.0% Biden)   Precincts: 2643/2655  Biden needs 64.72% [-0.027%]  Biden recent trend 76.16%
2020-11-05 16:43  ---  Trump leading by 14,857 votes   Remaining (est.): 50,366   Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 64.75% [7.870%]   Biden recent trend 76.20%
2020-11-05 16:37  ---  Trump leading by 14,857 votes   Remaining (est.): 107,985  Change:   6,130 (Trump 23.2% / 76.8% Biden)   Precincts: 2643/2655  Biden needs 56.88% [-1.706%]  Biden recent trend 76.20%
2020-11-05 16:26  ---  Trump leading by 18,146 votes   Remaining (est.): 105,678  Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 58.59% [2.237%]   Biden recent trend 76.69%
2020-11-05 15:53  ---  Trump leading by 18,146 votes   Remaining (est.): 142,914  Change:       4 (Trump 75.0% / 25.0% Biden)   Precincts: 2643/2655  Biden needs 56.35% [0.001%]   Biden recent trend 76.69%
2020-11-05 15:46  ---  Trump leading by 18,144 votes   Remaining (est.): 142,918  Change:   1,199 (Trump 31.6% / 68.4% Biden)   Precincts: 2643/2655  Biden needs 56.35% [-0.101%]  Biden recent trend 76.70%
2020-11-05 15:11  ---  Trump leading by 18,586 votes   Remaining (est.): 144,117  Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 56.45% [0.607%]   Biden recent trend 76.99%
2020-11-05 14:50  ---  Trump leading by 18,586 votes   Remaining (est.): 159,096  Change:     264 (Trump 49.2% / 50.8% Biden)   Precincts: 2643/2655  Biden needs 55.84% [0.008%]   Biden recent trend 76.99%
2020-11-05 14:10  ---  Trump leading by 18,590 votes   Remaining (est.): 159,360  Change:       2 (Trump 100.0% /  0.0% Biden)  Precincts: 2643/2655  Biden needs 55.83% [0.001%]   Biden recent trend 77.19%
2020-11-05 13:45  ---  Trump leading by 18,588 votes   Remaining (est.): 159,362  Change:   1,364 (Trump 51.8% / 48.2% Biden)   Precincts: 2643/2655  Biden needs 55.83% [0.064%]   Biden recent trend 77.19%
2020-11-05 10:55  ---  Trump leading by 18,540 votes   Remaining (est.): 160,726  Change:   8,468 (Trump 23.6% / 76.4% Biden)   Precincts: 2643/2655  Biden needs 55.77% [-1.032%]  Biden recent trend 78.38%
2020-11-05 06:12  ---  Trump leading by 23,009 votes   Remaining (est.): 169,194  Change:  10,461 (Trump 22.2% / 77.8% Biden)   Precincts: 2643/2655  Biden needs 56.80% [-1.089%]  Biden recent trend 76.00%
2020-11-05 05:24  ---  Trump leading by 28,827 votes   Remaining (est.): 182,707  Change:   4,344 (Trump 16.4% / 83.6% Biden)   Precincts: 2643/2655  Biden needs 57.89% [-0.598%]  Biden recent trend 75.38%
2020-11-05 04:22  ---  Trump leading by 31,748 votes   Remaining (est.): 187,051  Change:       0 (n/a)                         Precincts: 2643/2655  Biden needs 58.49% [0.000%]   Biden recent trend 68.73%
2020-11-05 03:57  ---  Trump leading by 31,748 votes   Remaining (est.): 187,051  Change:       0 (n/a)                         Precincts: 2605/2655  Biden needs 58.49% [0.000%]   Biden recent trend 68.73%
2020-11-05 03:53  ---  Trump leading by 31,748 votes   Remaining (est.): 187,051  Change:       0 (n/a)                         Precincts: 2575/2655  Biden needs 58.49% [0.000%]   Biden recent trend 68.73%
2020-11-05 03:53  ---  Trump leading by 31,748 votes   Remaining (est.): 187,051  Change:       0 (n/a)                         Precincts: 2572/2655  Biden needs 58.49% [0.000%]   Biden recent trend 68.73%
2020-11-05 03:49  ---  Trump leading by 31,748 votes   Remaining (est.): 187,051  Change:       0 (n/a)                         Precincts: 2557/2655  Biden needs 58.49% [0.000%]   Biden recent trend 68.73%
2020-11-05 03:41  ---  Trump leading by 31,748 votes   Remaining (est.): 187,051  Change:   3,128 (Trump 25.2% / 74.8% Biden)   Precincts: 2555/2655  Biden needs 58.49% [-0.268%]  Biden recent trend 68.73%
2020-11-05 03:11  ---  Trump leading by 33,300 votes   Remaining (est.): 190,179  Change:   6,682 (Trump 19.9% / 80.1% Biden)   Precincts: 2555/2655  Biden needs 58.75% [-0.724%]  Biden recent trend 68.16%
2020-11-05 02:58  ---  Trump leading by 37,322 votes   Remaining (est.): 196,861  Change:  -1,366 (Trump 51.8% / 48.2% Biden)   Precincts: 2555/2655  Biden needs 59.48% [-0.079%]  Biden recent trend 71.15%
2020-11-05 02:51  ---  Trump leading by 37,370 votes   Remaining (est.): 195,495  Change:  -3,242 (Trump 75.1% / 24.9% Biden)   Precincts: 2555/2655  Biden needs 59.56% [-0.584%]  Biden recent trend 74.10%
2020-11-05 02:19  ---  Trump leading by 38,996 votes   Remaining (est.): 192,253  Change:   1,414 (Trump 17.3% / 82.7% Biden)   Precincts: 2555/2655  Biden needs 60.14% [-0.165%]  Biden recent trend 70.43%
2020-11-05 02:03  ---  Trump leading by 39,921 votes   Remaining (est.): 193,667  Change:     423 (Trump 63.0% / 37.0% Biden)   Precincts: 2554/2655  Biden needs 60.31% [1.749%]   Biden recent trend 70.01%
2020-11-05 01:59  ---  Trump leading by 39,811 votes   Remaining (est.): 232,602  Change:       0 (n/a)                         Precincts: 2554/2655  Biden needs 58.56% [0.974%]   Biden recent trend 70.35%
2020-11-05 01:34  ---  Trump leading by 39,811 votes   Remaining (est.): 262,491  Change:      10 (Trump 20.0% / 80.0% Biden)   Precincts: 2554/2655  Biden needs 57.58% [-0.001%]  Biden recent trend 70.35%
2020-11-05 01:28  ---  Trump leading by 39,817 votes   Remaining (est.): 262,501  Change:  14,511 (Trump 28.7% / 71.3% Biden)   Precincts: 2554/2655  Biden needs 57.58% [-0.717%]  Biden recent trend 70.35%
2020-11-05 01:16  ---  Trump leading by 45,991 votes   Remaining (est.): 277,012  Change:     261 (Trump 59.6% / 40.4% Biden)   Precincts: 2554/2655  Biden needs 58.30% [0.017%]   Biden recent trend 69.99%
2020-11-05 01:07  ---  Trump leading by 45,941 votes   Remaining (est.): 277,273  Change:       0 (n/a)                         Precincts: 2554/2655  Biden needs 58.28% [-0.142%]  Biden recent trend 70.20%
2020-11-05 01:00  ---  Trump leading by 45,941 votes   Remaining (est.): 272,599  Change:     752 (Trump 40.2% / 59.8% Biden)   Precincts: 2554/2655  Biden needs 58.43% [-0.004%]  Biden recent trend 70.20%
2020-11-05 00:55  ---  Trump leading by 46,088 votes   Remaining (est.): 273,351  Change:   2,796 (Trump 18.9% / 81.1% Biden)   Precincts: 2554/2655  Biden needs 58.43% [-0.230%]  Biden recent trend 70.42%
2020-11-05 00:44  ---  Trump leading by 47,827 votes   Remaining (est.): 276,147  Change:   6,580 (Trump 55.4% / 44.6% Biden)   Precincts: 2553/2655  Biden needs 58.66% [0.328%]   Biden recent trend 69.51%
2020-11-05 00:20  ---  Trump leading by 47,111 votes   Remaining (est.): 282,727  Change:  16,690 (Trump 21.2% / 78.8% Biden)   Precincts: 2553/2655  Biden needs 58.33% [-1.143%]  Biden recent trend 70.46%
2020-11-05 00:16  ---  Trump leading by 56,739 votes   Remaining (est.): 299,417  Change:   1,610 (Trump 33.6% / 66.4% Biden)   Precincts: 2540/2655  Biden needs 59.47% [-1.793%]  Biden recent trend 66.28%
2020-11-05 00:01  ---  Trump leading by 57,266 votes   Remaining (est.): 254,108  Change:       0 (n/a)                         Precincts: 2540/2655  Biden needs 61.27% [-0.301%]  Biden recent trend 66.28%
2020-11-04 23:59  ---  Trump leading by 57,266 votes   Remaining (est.): 247,495  Change:       0 (n/a)                         Precincts: 2540/2655  Biden needs 61.57% [-4.703%]  Biden recent trend 66.28%
2020-11-04 23:14  ---  Trump leading by 57,266 votes   Remaining (est.): 175,963  Change:     518 (Trump 37.4% / 62.6% Biden)   Precincts: 2540/2655  Biden needs 66.27% [0.011%]   Biden recent trend 66.28%
2020-11-04 23:01  ---  Trump leading by 57,397 votes   Remaining (est.): 176,481  Change:   7,324 (Trump 28.1% / 71.9% Biden)   Precincts: 2540/2655  Biden needs 66.26% [-0.223%]  Biden recent trend 66.34%
2020-11-04 22:54  ---  Trump leading by 60,598 votes   Remaining (est.): 183,805  Change:  24,030 (Trump 35.3% / 64.7% Biden)   Precincts: 2540/2655  Biden needs 66.48% [0.211%]   Biden recent trend 68.88%
2020-11-04 22:44  ---  Trump leading by 67,644 votes   Remaining (est.): 207,835  Change:   1,128 (Trump 18.1% / 81.9% Biden)   Precincts: 2515/2655  Biden needs 66.27% [-0.084%]  Biden recent trend 66.50%
2020-11-04 22:27  ---  Trump leading by 68,363 votes   Remaining (est.): 208,963  Change:       4 (Trump 50.0% / 50.0% Biden)   Precincts: 2514/2655  Biden needs 66.36% [0.000%]   Biden recent trend 65.95%
2020-11-04 22:12  ---  Trump leading by 68,363 votes   Remaining (est.): 208,967  Change:   1,366 (Trump 51.8% / 48.2% Biden)   Precincts: 2514/2655  Biden needs 66.36% [0.118%]   Biden recent trend 65.95%
2020-11-04 21:55  ---  Trump leading by 68,315 votes   Remaining (est.): 210,333  Change:   1,638 (Trump 61.2% / 38.8% Biden)   Precincts: 2514/2655  Biden needs 66.24% [0.212%]   Biden recent trend 66.75%
2020-11-04 21:45  ---  Trump leading by 67,948 votes   Remaining (est.): 211,971  Change:  16,952 (Trump 21.4% / 78.6% Biden)   Precincts: 2513/2655  Biden needs 66.03% [-0.929%]  Biden recent trend 67.27%
2020-11-04 21:37  ---  Trump leading by 77,636 votes   Remaining (est.): 228,923  Change:     575 (Trump 57.8% / 42.2% Biden)   Precincts: 2513/2655  Biden needs 66.96% [0.062%]   Biden recent trend 65.90%
2020-11-04 21:29  ---  Trump leading by 77,546 votes   Remaining (est.): 229,498  Change:     691 (Trump 18.0% / 82.0% Biden)   Precincts: 2513/2655  Biden needs 66.89% [-0.045%]  Biden recent trend 66.36%
2020-11-04 21:12  ---  Trump leading by 77,988 votes   Remaining (est.): 230,189  Change:     343 (Trump 53.1% / 46.9% Biden)   Precincts: 2513/2655  Biden needs 66.94% [0.030%]   Biden recent trend 65.14%
2020-11-04 20:53  ---  Trump leading by 77,967 votes   Remaining (est.): 230,532  Change:   9,984 (Trump 47.7% / 52.3% Biden)   Precincts: 2513/2655  Biden needs 66.91% [0.606%]   Biden recent trend 65.34%
2020-11-04 20:48  ---  Trump leading by 78,429 votes   Remaining (est.): 240,516  Change:      11 (Trump 77.3% / 22.7% Biden)   Precincts: 2510/2655  Biden needs 66.30% [0.002%]   Biden recent trend 78.58%
2020-11-04 20:37  ---  Trump leading by 78,423 votes   Remaining (est.): 240,527  Change:       1 (n/a)                         Precincts: 2510/2655  Biden needs 66.30% [-0.000%]  Biden recent trend 78.59%
2020-11-04 20:14  ---  Trump leading by 78,424 votes   Remaining (est.): 240,528  Change:   4,905 (Trump 38.9% / 61.1% Biden)   Precincts: 2510/2655  Biden needs 66.30% [0.105%]   Biden recent trend 78.59%
2020-11-04 19:44  ---  Trump leading by 79,509 votes   Remaining (est.): 245,433  Change:   5,924 (Trump 17.7% / 82.3% Biden)   Precincts: 2508/2655  Biden needs 66.20% [-0.379%]  Biden recent trend 80.98%
2020-11-04 19:36  ---  Trump leading by 83,331 votes   Remaining (est.): 251,357  Change:     811 (Trump 65.7% / 34.3% Biden)   Precincts: 2506/2655  Biden needs 66.58% [0.104%]   Biden recent trend 80.73%
2020-11-04 19:33  ---  Trump leading by 83,076 votes   Remaining (est.): 252,168  Change:      80 (Trump 50.6% / 49.4% Biden)   Precincts: 2506/2655  Biden needs 66.47% [0.005%]   Biden recent trend 81.51%
2020-11-04 19:28  ---  Trump leading by 83,075 votes   Remaining (est.): 252,248  Change:     291 (Trump 45.4% / 54.6% Biden)   Precincts: 2506/2655  Biden needs 66.47% [0.014%]   Biden recent trend 81.59%
2020-11-04 19:13  ---  Trump leading by 83,102 votes   Remaining (est.): 252,539  Change:     499 (Trump 61.5% / 38.5% Biden)   Precincts: 2506/2655  Biden needs 66.45% [0.055%]   Biden recent trend 81.83%
2020-11-04 18:40  ---  Trump leading by 82,987 votes   Remaining (est.): 253,038  Change:   4,131 (Trump 15.5% / 84.5% Biden)   Precincts: 2506/2655  Biden needs 66.40% [-0.290%]  Biden recent trend 82.52%
2020-11-04 17:57  ---  Trump leading by 85,835 votes   Remaining (est.): 257,169  Change:   2,376 (Trump 17.6% / 82.4% Biden)   Precincts: 2504/2655  Biden needs 66.69% [-0.144%]  Biden recent trend 81.81%
2020-11-04 17:44  ---  Trump leading by 87,376 votes   Remaining (est.): 259,545  Change:   2,438 (Trump 45.1% / 54.9% Biden)   Precincts: 2503/2655  Biden needs 66.83% [0.111%]   Biden recent trend 81.75%
2020-11-04 17:41  ---  Trump leading by 87,617 votes   Remaining (est.): 261,983  Change:     247 (Trump 51.8% / 48.2% Biden)   Precincts: 2503/2655  Biden needs 66.72% [0.017%]   Biden recent trend 84.55%
2020-11-04 17:16  ---  Trump leading by 87,608 votes   Remaining (est.): 262,230  Change:  19,244 (Trump 13.0% / 87.0% Biden)   Precincts: 2503/2655  Biden needs 66.70% [-1.386%]  Biden recent trend 84.94%
2020-11-04 16:37  ---  Trump leading by 101,837 votes  Remaining (est.): 281,474  Change:   3,290 (Trump 22.9% / 77.1% Biden)   Precincts: 2493/2655  Biden needs 68.09% [-0.104%]  Biden recent trend 74.83%
2020-11-04 16:35  ---  Trump leading by 103,620 votes  Remaining (est.): 284,764  Change:     502 (Trump 41.5% / 58.5% Biden)   Precincts: 2493/2655  Biden needs 68.19% [0.017%]   Biden recent trend 62.07%
2020-11-04 14:15  ---  Trump leading by 103,705 votes  Remaining (est.): 285,266  Change:      78 (Trump 14.7% / 85.3% Biden)   Precincts: 2493/2655  Biden needs 68.18% [-0.005%]  Biden recent trend 84.62%
2020-11-04 13:28  ---  Trump leading by 103,760 votes  Remaining (est.): 285,344  Change:       0 (n/a)                         Precincts: 2493/2655  Biden needs 68.18% [68.182%]  Biden recent trend n/a
----------------  ---  ------------------------------  -------------------------  --------------------------------------------  --------------------  ----------------------------  -------------------------

North Carolina (EV: 15) Total Votes: (Trump: 2,732,120, Biden: 2,655,383)
----------------  ---  -----------------------------  -------------------------  -------------------------------------------  --------------------  ----------------------------  -------------------------
2020-11-05 18:03  ---  Trump leading by 76,737 votes  Remaining (est.): 190,621  Change:   2,233 (Trump 50.0% / 50.0% Biden)  Precincts: 2425/2662  Biden needs 70.13% [0.233%]   Biden recent trend 49.98%
2020-11-05 16:46  ---  Trump leading by 76,737 votes  Remaining (est.): 192,854  Change:   1,989 (Trump 50.0% / 50.0% Biden)  Precincts: 2424/2662  Biden needs 69.90% [0.203%]   Biden recent trend 49.98%
2020-11-05 01:02  ---  Trump leading by 76,737 votes  Remaining (est.): 194,843  Change:       0 (n/a)                        Precincts: 2423/2662  Biden needs 69.69% [-0.087%]  Biden recent trend 50.00%
2020-11-04 21:41  ---  Trump leading by 76,737 votes  Remaining (est.): 193,984  Change:     220 (Trump 50.0% / 50.0% Biden)  Precincts: 2423/2662  Biden needs 69.78% [0.022%]   Biden recent trend 50.00%
2020-11-04 13:28  ---  Trump leading by 76,737 votes  Remaining (est.): 194,204  Change:       0 (n/a)                        Precincts: 2423/2662  Biden needs 69.76% [69.757%]  Biden recent trend n/a
----------------  ---  -----------------------------  -------------------------  -------------------------------------------  --------------------  ----------------------------  -------------------------

Nevada (EV: 6) Total Votes: (Biden: 604,251, Trump: 592,813)
----------------  ---  -----------------------------  -------------------------  -------------------------------------------  --------------------  ----------------------------  -------------------------
2020-11-05 18:24  ---  Biden leading by 11,438 votes  Remaining (est.): 146,757  Change:      15 (Biden 50.0% / 50.0% Trump)  Precincts: 1764/1993  Trump needs 53.90% [0.592%]   Trump recent trend 43.56%
2020-11-05 18:03  ---  Biden leading by 11,438 votes  Remaining (est.): 173,045  Change:     913 (Biden 30.9% / 69.1% Trump)  Precincts: 1619/1993  Trump needs 53.30% [-0.083%]  Trump recent trend 43.55%
2020-11-05 17:32  ---  Biden leading by 11,787 votes  Remaining (est.): 173,958  Change:  12,189 (Biden 51.4% / 48.6% Trump)  Precincts: 1614/1993  Trump needs 53.39% [0.311%]   Trump recent trend 42.73%
2020-11-05 17:21  ---  Biden leading by 11,454 votes  Remaining (est.): 186,147  Change:   1,019 (Biden 33.5% / 66.5% Trump)  Precincts: 1614/1993  Trump needs 53.08% [-0.073%]  Trump recent trend 38.32%
2020-11-05 17:16  ---  Biden leading by 11,791 votes  Remaining (est.): 187,166  Change:     995 (Biden 37.4% / 62.6% Trump)  Precincts: 1610/1993  Trump needs 53.15% [-0.050%]  Trump recent trend 36.44%
2020-11-05 17:12  ---  Biden leading by 12,042 votes  Remaining (est.): 188,161  Change:       0 (n/a)                        Precincts: 1610/1993  Trump needs 53.20% [0.000%]   Trump recent trend 34.62%
2020-11-05 17:09  ---  Biden leading by 12,042 votes  Remaining (est.): 188,161  Change:  14,285 (Biden 65.4% / 34.6% Trump)  Precincts: 1729/1993  Trump needs 53.20% [1.311%]   Trump recent trend 34.62%
2020-11-04 16:35  ---  Biden leading by 7,647 votes   Remaining (est.): 202,446  Change:       0 (n/a)                        Precincts: 1610/1993  Trump needs 51.89% [0.001%]   Trump recent trend n/a
2020-11-04 13:28  ---  Biden leading by 7,647 votes   Remaining (est.): 202,563  Change:       0 (n/a)                        Precincts: 1610/1993  Trump needs 51.89% [51.888%]  Trump recent trend n/a
----------------  ---  -----------------------------  -------------------------  -------------------------------------------  --------------------  ----------------------------  -------------------------

Pennsylvania (EV: 20) Total Votes: (Trump: 3,280,308, Biden: 3,243,736)
----------------  ---  ------------------------------  ---------------------------  --------------------------------------------  --------------------  ----------------------------  -------------------------
2020-11-06 03:45  ***  Trump leading by 26,319 votes   Remaining (est.): 291,894    Change:  13,909 (Trump 13.1% / 86.9% Biden)   Precincts: 8366/9128  Biden needs 54.51% [-1.471%]  Biden recent trend 78.50%
2020-11-06 03:33  ***  Trump leading by 36,572 votes   Remaining (est.): 305,803    Change:   8,046 (Trump 20.6% / 79.4% Biden)   Precincts: 8338/9128  Biden needs 55.98% [-0.601%]  Biden recent trend 72.42%
2020-11-06 03:16  ---  Trump leading by 41,305 votes   Remaining (est.): 313,849    Change:   2,431 (Trump 32.8% / 67.2% Biden)   Precincts: 8328/9128  Biden needs 56.58% [-0.082%]  Biden recent trend 71.73%
2020-11-06 02:37  ---  Trump leading by 42,142 votes   Remaining (est.): 316,280    Change:  15,150 (Trump 27.8% / 72.2% Biden)   Precincts: 8326/9128  Biden needs 56.66% [-0.708%]  Biden recent trend 72.03%
2020-11-06 02:07  ---  Trump leading by 48,854 votes   Remaining (est.): 331,430    Change:   4,035 (Trump 27.5% / 72.5% Biden)   Precincts: 8316/9128  Biden needs 57.37% [-0.182%]  Biden recent trend 71.56%
2020-11-06 01:52  ---  Trump leading by 50,671 votes   Remaining (est.): 335,465    Change:   7,474 (Trump 32.9% / 67.1% Biden)   Precincts: 8309/9128  Biden needs 57.55% [-0.207%]  Biden recent trend 71.44%
2020-11-06 01:44  ---  Trump leading by 53,221 votes   Remaining (est.): 342,939    Change:       0 (n/a)                         Precincts: 8298/9128  Biden needs 57.76% [0.112%]   Biden recent trend 77.60%
2020-11-06 01:33  ---  Trump leading by 53,221 votes   Remaining (est.): 347,939    Change:   9,975 (Trump 24.6% / 75.4% Biden)   Precincts: 8298/9128  Biden needs 57.65% [-0.494%]  Biden recent trend 77.60%
2020-11-06 01:26  ---  Trump leading by 58,286 votes   Remaining (est.): 357,914    Change:   4,130 (Trump 20.2% / 79.8% Biden)   Precincts: 8298/9128  Biden needs 58.14% [-0.248%]  Biden recent trend 76.95%
2020-11-06 01:22  ---  Trump leading by 60,751 votes   Remaining (est.): 362,044    Change:   8,867 (Trump 33.2% / 66.8% Biden)   Precincts: 8293/9128  Biden needs 58.39% [-0.200%]  Biden recent trend 77.37%
2020-11-06 00:55  ---  Trump leading by 63,725 votes   Remaining (est.): 370,911    Change:     936 (Trump 21.1% / 78.9% Biden)   Precincts: 8287/9128  Biden needs 58.59% [-0.051%]  Biden recent trend 72.11%
2020-11-06 00:35  ---  Trump leading by 64,266 votes   Remaining (est.): 371,847    Change:     545 (Trump 52.7% / 47.3% Biden)   Precincts: 8286/9128  Biden needs 58.64% [0.017%]   Biden recent trend 72.00%
2020-11-06 00:11  ---  Trump leading by 64,237 votes   Remaining (est.): 372,392    Change:  12,480 (Trump 12.5% / 87.5% Biden)   Precincts: 8284/9128  Biden needs 58.62% [-0.938%]  Biden recent trend 72.23%
2020-11-05 23:59  ---  Trump leading by 73,609 votes   Remaining (est.): 384,872    Change:     205 (Trump 60.5% / 39.5% Biden)   Precincts: 8258/9128  Biden needs 59.56% [0.011%]   Biden recent trend 68.03%
2020-11-05 23:51  ---  Trump leading by 73,566 votes   Remaining (est.): 385,077    Change:   4,634 (Trump 29.9% / 70.1% Biden)   Precincts: 8258/9128  Biden needs 59.55% [-0.125%]  Biden recent trend 68.16%
2020-11-05 23:30  ---  Trump leading by 75,427 votes   Remaining (est.): 389,711    Change:   4,416 (Trump 17.3% / 82.7% Biden)   Precincts: 8254/9128  Biden needs 59.68% [-0.258%]  Biden recent trend 67.95%
2020-11-05 23:25  ---  Trump leading by 78,314 votes   Remaining (est.): 394,127    Change:     405 (Trump 50.0% / 50.0% Biden)   Precincts: 8230/9128  Biden needs 59.94% [0.010%]   Biden recent trend 66.15%
2020-11-05 22:53  ---  Trump leading by 78,314 votes   Remaining (est.): 394,532    Change:  35,858 (Trump 33.7% / 66.3% Biden)   Precincts: 8229/9128  Biden needs 59.92% [-0.534%]  Biden recent trend 66.33%
2020-11-05 22:45  ---  Trump leading by 90,027 votes   Remaining (est.): 430,390    Change:   1,333 (Trump 30.7% / 69.3% Biden)   Precincts: 8196/9128  Biden needs 60.46% [-0.027%]  Biden recent trend 72.54%
2020-11-05 22:15  ---  Trump leading by 90,542 votes   Remaining (est.): 431,723    Change:      12 (Trump 50.0% / 50.0% Biden)   Precincts: 8195/9128  Biden needs 60.49% [0.000%]   Biden recent trend 72.66%
2020-11-05 21:53  ---  Trump leading by 90,542 votes   Remaining (est.): 431,735    Change:  13,957 (Trump 31.6% / 68.4% Biden)   Precincts: 8195/9128  Biden needs 60.49% [-0.247%]  Biden recent trend 72.67%
2020-11-05 21:45  ---  Trump leading by 95,670 votes   Remaining (est.): 445,692    Change:   6,636 (Trump 38.3% / 61.7% Biden)   Precincts: 8186/9128  Biden needs 60.73% [-0.014%]  Biden recent trend 70.71%
2020-11-05 21:29  ---  Trump leading by 97,224 votes   Remaining (est.): 452,328    Change:   1,830 (Trump 31.3% / 68.7% Biden)   Precincts: 8177/9128  Biden needs 60.75% [-0.032%]  Biden recent trend 71.72%
2020-11-05 21:21  ---  Trump leading by 97,908 votes   Remaining (est.): 454,158    Change:      86 (Trump 54.7% / 45.3% Biden)   Precincts: 8175/9128  Biden needs 60.78% [0.003%]   Biden recent trend 71.90%
2020-11-05 21:13  ---  Trump leading by 97,900 votes   Remaining (est.): 454,244    Change:       0 (n/a)                         Precincts: 8175/9128  Biden needs 60.78% [0.562%]   Biden recent trend 71.97%
2020-11-05 21:05  ---  Trump leading by 97,900 votes   Remaining (est.): 479,244    Change:  13,746 (Trump 17.0% / 83.0% Biden)   Precincts: 8175/9128  Biden needs 60.21% [-0.636%]  Biden recent trend 71.97%
2020-11-05 20:48  ---  Trump leading by 106,980 votes  Remaining (est.): 492,990    Change:   4,486 (Trump 34.5% / 65.5% Biden)   Precincts: 8168/9128  Biden needs 60.85% [-0.042%]  Biden recent trend 69.50%
2020-11-05 20:40  ---  Trump leading by 108,367 votes  Remaining (est.): 497,476    Change:     942 (Trump 37.5% / 62.5% Biden)   Precincts: 8160/9128  Biden needs 60.89% [-0.003%]  Biden recent trend 70.04%
2020-11-05 20:15  ---  Trump leading by 108,602 votes  Remaining (est.): 498,418    Change:      34 (Trump 64.7% / 35.3% Biden)   Precincts: 8159/9128  Biden needs 60.89% [0.002%]   Biden recent trend 70.26%
2020-11-05 20:01  ---  Trump leading by 108,592 votes  Remaining (est.): 498,452    Change:   3,862 (Trump 47.7% / 52.3% Biden)   Precincts: 8159/9128  Biden needs 60.89% [0.066%]   Biden recent trend 70.29%
2020-11-05 19:37  ---  Trump leading by 108,772 votes  Remaining (est.): 502,314    Change:   3,349 (Trump 35.3% / 64.7% Biden)   Precincts: 8155/9128  Biden needs 60.83% [-0.025%]  Biden recent trend 78.38%
2020-11-05 19:32  ---  Trump leading by 109,754 votes  Remaining (est.): 505,663    Change:   4,226 (Trump 30.9% / 69.1% Biden)   Precincts: 8153/9128  Biden needs 60.85% [-0.068%]  Biden recent trend 79.57%
2020-11-05 19:29  ---  Trump leading by 111,369 votes  Remaining (est.): 509,889    Change:       0 (n/a)                         Precincts: 8147/9128  Biden needs 60.92% [-0.008%]  Biden recent trend 80.85%
2020-11-05 19:12  ---  Trump leading by 111,369 votes  Remaining (est.): 509,512    Change:   4,515 (Trump 20.7% / 79.3% Biden)   Precincts: 8147/9128  Biden needs 60.93% [-0.161%]  Biden recent trend 80.85%
2020-11-05 18:24  ---  Trump leading by 114,011 votes  Remaining (est.): 514,027    Change:   2,451 (Trump 43.3% / 56.7% Biden)   Precincts: 8147/9128  Biden needs 61.09% [0.021%]   Biden recent trend 81.40%
2020-11-05 18:15  ---  Trump leading by 114,341 votes  Remaining (est.): 516,478    Change:   3,076 (Trump 38.2% / 61.8% Biden)   Precincts: 8144/9128  Biden needs 61.07% [-0.005%]  Biden recent trend 83.30%
2020-11-05 17:21  ---  Trump leading by 115,069 votes  Remaining (est.): 519,554    Change:   3,003 (Trump 30.8% / 69.2% Biden)   Precincts: 8139/9128  Biden needs 61.07% [-0.047%]  Biden recent trend 81.61%
2020-11-05 17:02  ---  Trump leading by 116,224 votes  Remaining (est.): 522,557    Change:   8,193 (Trump 15.6% / 84.4% Biden)   Precincts: 8135/9128  Biden needs 61.12% [-0.359%]  Biden recent trend 82.53%
2020-11-05 16:59  ---  Trump leading by 121,856 votes  Remaining (est.): 530,750    Change:      67 (Trump 49.3% / 50.7% Biden)   Precincts: 8126/9128  Biden needs 61.48% [0.001%]   Biden recent trend 82.06%
2020-11-05 16:49  ---  Trump leading by 121,857 votes  Remaining (est.): 530,817    Change:  13,114 (Trump  9.0% / 91.0% Biden)   Precincts: 8126/9128  Biden needs 61.48% [-0.712%]  Biden recent trend 82.13%
2020-11-05 16:43  ---  Trump leading by 132,611 votes  Remaining (est.): 543,931    Change:   4,507 (Trump 16.6% / 83.4% Biden)   Precincts: 8098/9128  Biden needs 62.19% [-0.175%]  Biden recent trend 80.82%
2020-11-05 16:08  ---  Trump leading by 135,626 votes  Remaining (est.): 548,438    Change:      73 (Trump 19.2% / 80.8% Biden)   Precincts: 8095/9128  Biden needs 62.36% [-0.002%]  Biden recent trend 80.57%
2020-11-05 15:26  ---  Trump leading by 135,671 votes  Remaining (est.): 548,511    Change:     554 (Trump 51.4% / 48.6% Biden)   Precincts: 8095/9128  Biden needs 62.37% [0.014%]   Biden recent trend 80.57%
2020-11-05 15:12  ---  Trump leading by 135,655 votes  Remaining (est.): 549,065    Change:     103 (Trump 27.2% / 72.8% Biden)   Precincts: 8095/9128  Biden needs 62.35% [-0.002%]  Biden recent trend 80.95%
2020-11-05 15:11  ---  Trump leading by 135,702 votes  Remaining (est.): 549,168    Change:       0 (n/a)                         Precincts: 8095/9128  Biden needs 62.36% [0.822%]   Biden recent trend 80.97%
2020-11-05 14:36  ---  Trump leading by 135,702 votes  Remaining (est.): 588,300    Change:  13,724 (Trump 25.4% / 74.6% Biden)   Precincts: 8095/9128  Biden needs 61.53% [-0.299%]  Biden recent trend 80.97%
2020-11-05 14:20  ---  Trump leading by 142,466 votes  Remaining (est.): 602,024    Change:       0 (n/a)                         Precincts: 8084/9128  Biden needs 61.83% [1.401%]   Biden recent trend 83.63%
2020-11-05 14:10  ---  Trump leading by 142,466 votes  Remaining (est.): 682,876    Change:  32,633 (Trump 16.4% / 83.6% Biden)   Precincts: 8084/9128  Biden needs 60.43% [-1.058%]  Biden recent trend 83.63%
2020-11-05 04:22  ---  Trump leading by 164,414 votes  Remaining (est.): 715,509    Change:  38,652 (Trump 26.3% / 73.7% Biden)   Precincts: 8031/9128  Biden needs 61.49% [-0.626%]  Biden recent trend 73.71%
2020-11-05 03:57  ---  Trump leading by 182,743 votes  Remaining (est.): 754,161    Change:   1,318 (Trump 56.9% / 43.1% Biden)   Precincts: 8000/9128  Biden needs 62.12% [0.033%]   Biden recent trend 70.06%
2020-11-05 03:53  ---  Trump leading by 182,561 votes  Remaining (est.): 755,479    Change:   5,743 (Trump 34.0% / 66.0% Biden)   Precincts: 7998/9128  Biden needs 62.08% [-0.029%]  Biden recent trend 71.19%
2020-11-05 03:16  ---  Trump leading by 184,397 votes  Remaining (est.): 761,222    Change:   4,549 (Trump 24.1% / 75.9% Biden)   Precincts: 7990/9128  Biden needs 62.11% [-0.082%]  Biden recent trend 71.90%
2020-11-05 02:47  ---  Trump leading by 186,755 votes  Remaining (est.): 765,771    Change:  10,173 (Trump 29.2% / 70.8% Biden)   Precincts: 7982/9128  Biden needs 62.19% [-0.113%]  Biden recent trend 71.32%
2020-11-05 02:26  ---  Trump leading by 190,984 votes  Remaining (est.): 775,944    Change:   1,285 (Trump -4.3% / 104.3% Biden)  Precincts: 7973/9128  Biden needs 62.31% [-0.069%]  Biden recent trend 70.41%
2020-11-05 02:19  ---  Trump leading by 192,380 votes  Remaining (est.): 777,229    Change:     440 (Trump 50.0% / 50.0% Biden)   Precincts: 7972/9128  Biden needs 62.38% [0.007%]   Biden recent trend 69.55%
2020-11-05 02:03  ---  Trump leading by 192,380 votes  Remaining (est.): 777,669    Change:   9,413 (Trump 31.0% / 69.0% Biden)   Precincts: 7971/9128  Biden needs 62.37% [-0.079%]  Biden recent trend 69.72%
2020-11-05 01:21  ---  Trump leading by 195,953 votes  Remaining (est.): 787,082    Change:  10,590 (Trump 29.2% / 70.8% Biden)   Precincts: 7958/9128  Biden needs 62.45% [-0.111%]  Biden recent trend 69.89%
2020-11-05 01:16  ---  Trump leading by 200,360 votes  Remaining (est.): 797,672    Change:       0 (n/a)                         Precincts: 7948/9128  Biden needs 62.56% [-0.013%]  Biden recent trend 69.57%
2020-11-05 01:07  ---  Trump leading by 200,360 votes  Remaining (est.): 796,837    Change:   5,612 (Trump 20.2% / 79.8% Biden)   Precincts: 7948/9128  Biden needs 62.57% [-0.120%]  Biden recent trend 69.57%
2020-11-05 01:02  ---  Trump leading by 203,700 votes  Remaining (est.): 802,449    Change:     217 (Trump 31.3% / 68.7% Biden)   Precincts: 7938/9128  Biden needs 62.69% [-0.002%]  Biden recent trend 70.54%
2020-11-05 00:55  ---  Trump leading by 203,781 votes  Remaining (est.): 802,666    Change:  24,511 (Trump 32.8% / 67.2% Biden)   Precincts: 7937/9128  Biden needs 62.69% [-0.135%]  Biden recent trend 70.54%
2020-11-05 00:44  ---  Trump leading by 212,236 votes  Remaining (est.): 827,177    Change:  77,985 (Trump 28.4% / 71.6% Biden)   Precincts: 7904/9128  Biden needs 62.83% [-0.754%]  Biden recent trend 71.58%
2020-11-05 00:16  ---  Trump leading by 245,891 votes  Remaining (est.): 905,162    Change:  19,381 (Trump 23.8% / 76.2% Biden)   Precincts: 7818/9128  Biden needs 63.58% [-0.265%]  Biden recent trend 75.13%
2020-11-04 23:49  ---  Trump leading by 256,058 votes  Remaining (est.): 924,543    Change:  21,097 (Trump 25.9% / 74.1% Biden)   Precincts: 7805/9128  Biden needs 63.85% [-0.229%]  Biden recent trend 68.57%
2020-11-04 23:41  ---  Trump leading by 266,234 votes  Remaining (est.): 945,640    Change:   3,360 (Trump 58.3% / 41.7% Biden)   Precincts: 7787/9128  Biden needs 64.08% [0.079%]   Biden recent trend 67.35%
2020-11-04 23:28  ---  Trump leading by 265,676 votes  Remaining (est.): 949,000    Change:   3,197 (Trump 33.1% / 66.9% Biden)   Precincts: 7781/9128  Biden needs 64.00% [-0.010%]  Biden recent trend 73.96%
2020-11-04 23:24  ---  Trump leading by 266,754 votes  Remaining (est.): 952,197    Change:   2,190 (Trump 52.3% / 47.7% Biden)   Precincts: 7777/9128  Biden needs 64.01% [0.037%]   Biden recent trend 74.40%
2020-11-04 23:16  ---  Trump leading by 266,652 votes  Remaining (est.): 954,387    Change:   4,579 (Trump 26.0% / 74.0% Biden)   Precincts: 7776/9128  Biden needs 63.97% [-0.048%]  Biden recent trend 75.57%
2020-11-04 23:14  ---  Trump leading by 268,846 votes  Remaining (est.): 958,966    Change:  17,642 (Trump 26.9% / 73.1% Biden)   Precincts: 7771/9128  Biden needs 64.02% [-0.163%]  Biden recent trend 75.74%
2020-11-04 23:01  ---  Trump leading by 276,979 votes  Remaining (est.): 976,608    Change:  27,630 (Trump 22.5% / 77.5% Biden)   Precincts: 7744/9128  Biden needs 64.18% [-0.365%]  Biden recent trend 76.91%
2020-11-04 22:54  ---  Trump leading by 292,148 votes  Remaining (est.): 1,004,238  Change:   1,365 (Trump 40.3% / 59.7% Biden)   Precincts: 7713/9128  Biden needs 64.55% [0.007%]   Biden recent trend 68.24%
2020-11-04 22:44  ---  Trump leading by 292,412 votes  Remaining (est.): 1,005,603  Change:  16,361 (Trump 22.6% / 77.4% Biden)   Precincts: 7712/9128  Biden needs 64.54% [-0.206%]  Biden recent trend 71.72%
2020-11-04 22:38  ---  Trump leading by 301,386 votes  Remaining (est.): 1,021,964  Change:  10,527 (Trump 40.9% / 59.1% Biden)   Precincts: 7684/9128  Biden needs 64.75% [0.058%]   Biden recent trend 70.40%
2020-11-04 22:21  ---  Trump leading by 303,293 votes  Remaining (est.): 1,032,491  Change:   2,306 (Trump 50.0% / 50.0% Biden)   Precincts: 7675/9128  Biden needs 64.69% [0.033%]   Biden recent trend 73.96%
2020-11-04 22:19  ---  Trump leading by 303,293 votes  Remaining (est.): 1,034,797  Change:   7,256 (Trump 15.9% / 84.1% Biden)   Precincts: 7671/9128  Biden needs 64.65% [-0.135%]  Biden recent trend 75.73%
2020-11-04 22:12  ---  Trump leading by 308,243 votes  Remaining (est.): 1,042,053  Change:   2,056 (Trump 51.8% / 48.2% Biden)   Precincts: 7659/9128  Biden needs 64.79% [0.033%]   Biden recent trend 76.34%
2020-11-04 22:04  ---  Trump leading by 308,167 votes  Remaining (est.): 1,044,109  Change:  21,949 (Trump 24.5% / 75.5% Biden)   Precincts: 7657/9128  Biden needs 64.76% [-0.222%]  Biden recent trend 77.41%
2020-11-04 22:01  ---  Trump leading by 319,377 votes  Remaining (est.): 1,066,058  Change:       0 (n/a)                         Precincts: 7639/9128  Biden needs 64.98% [-0.051%]  Biden recent trend 78.71%
2020-11-04 21:55  ---  Trump leading by 319,377 votes  Remaining (est.): 1,062,475  Change:       0 (n/a)                         Precincts: 7639/9128  Biden needs 65.03% [-0.031%]  Biden recent trend 78.71%
2020-11-04 21:54  ---  Trump leading by 319,377 votes  Remaining (est.): 1,060,286  Change:       0 (n/a)                         Precincts: 7639/9128  Biden needs 65.06% [-0.096%]  Biden recent trend 78.71%
2020-11-04 21:45  ---  Trump leading by 319,377 votes  Remaining (est.): 1,053,561  Change:   3,278 (Trump 42.4% / 57.6% Biden)   Precincts: 7639/9128  Biden needs 65.16% [0.023%]   Biden recent trend 78.71%
2020-11-04 21:41  ---  Trump leading by 319,876 votes  Remaining (est.): 1,056,839  Change:     887 (Trump 46.9% / 53.1% Biden)   Precincts: 7634/9128  Biden needs 65.13% [0.010%]   Biden recent trend 72.73%
2020-11-04 21:29  ---  Trump leading by 319,931 votes  Remaining (est.): 1,057,726  Change:  27,598 (Trump 18.0% / 82.0% Biden)   Precincts: 7633/9128  Biden needs 65.12% [-0.430%]  Biden recent trend 73.07%
2020-11-04 21:21  ---  Trump leading by 337,616 votes  Remaining (est.): 1,085,324  Change:  23,395 (Trump 37.5% / 62.5% Biden)   Precincts: 7577/9128  Biden needs 65.55% [0.065%]   Biden recent trend 66.84%
2020-11-04 21:04  ---  Trump leading by 343,462 votes  Remaining (est.): 1,108,719  Change:  12,427 (Trump 25.0% / 75.0% Biden)   Precincts: 7557/9128  Biden needs 65.49% [-0.106%]  Biden recent trend 85.27%
2020-11-04 20:58  ---  Trump leading by 349,685 votes  Remaining (est.): 1,121,146  Change:   8,760 (Trump 26.3% / 73.7% Biden)   Precincts: 7547/9128  Biden needs 65.59% [-0.063%]  Biden recent trend 88.54%
2020-11-04 20:53  ---  Trump leading by 353,834 votes  Remaining (est.): 1,129,906  Change:   5,527 (Trump 26.6% / 73.4% Biden)   Precincts: 7538/9128  Biden needs 65.66% [-0.038%]  Biden recent trend 92.85%
2020-11-04 20:48  ---  Trump leading by 356,423 votes  Remaining (est.): 1,135,433  Change:  24,654 (Trump  2.8% / 97.2% Biden)   Precincts: 7531/9128  Biden needs 65.70% [-0.670%]  Biden recent trend 90.39%
2020-11-04 20:37  ---  Trump leading by 379,700 votes  Remaining (est.): 1,160,087  Change:  15,540 (Trump 20.4% / 79.6% Biden)   Precincts: 7479/9128  Biden needs 66.37% [-0.174%]  Biden recent trend 79.85%
2020-11-04 20:23  ---  Trump leading by 388,889 votes  Remaining (est.): 1,175,627  Change:  16,214 (Trump 19.9% / 80.1% Biden)   Precincts: 7451/9128  Biden needs 66.54% [-0.185%]  Biden recent trend 86.93%
2020-11-04 20:14  ---  Trump leading by 398,656 votes  Remaining (est.): 1,191,841  Change:  31,489 (Trump  9.6% / 90.4% Biden)   Precincts: 7423/9128  Biden needs 66.72% [-0.610%]  Biden recent trend 90.44%
2020-11-04 20:10  ---  Trump leading by 424,125 votes  Remaining (est.): 1,223,330  Change:  28,498 (Trump 30.6% / 69.4% Biden)   Precincts: 7368/9128  Biden needs 67.33% [-0.048%]  Biden recent trend 72.91%
2020-11-04 20:06  ---  Trump leading by 435,204 votes  Remaining (est.): 1,251,828  Change:     197 (Trump 43.1% / 56.9% Biden)   Precincts: 7343/9128  Biden needs 67.38% [0.002%]   Biden recent trend 77.29%
2020-11-04 20:03  ---  Trump leading by 435,231 votes  Remaining (est.): 1,252,025  Change:  23,686 (Trump 22.8% / 77.2% Biden)   Precincts: 7343/9128  Biden needs 67.38% [-0.183%]  Biden recent trend 77.42%
2020-11-04 19:47  ---  Trump leading by 448,129 votes  Remaining (est.): 1,275,711  Change:   6,879 (Trump 21.9% / 78.1% Biden)   Precincts: 7310/9128  Biden needs 67.56% [-0.056%]  Biden recent trend 71.52%
2020-11-04 19:44  ---  Trump leading by 451,993 votes  Remaining (est.): 1,282,590  Change:  14,117 (Trump 25.2% / 74.8% Biden)   Precincts: 7299/9128  Biden needs 67.62% [-0.078%]  Biden recent trend 69.06%
2020-11-04 19:33  ---  Trump leading by 459,000 votes  Remaining (est.): 1,296,707  Change:     134 (Trump 81.7% / 18.3% Biden)   Precincts: 7286/9128  Biden needs 67.70% [0.005%]   Biden recent trend 68.04%
2020-11-04 19:17  ---  Trump leading by 458,915 votes  Remaining (est.): 1,296,841  Change:  10,555 (Trump 36.5% / 63.5% Biden)   Precincts: 7286/9128  Biden needs 67.69% [0.034%]   Biden recent trend 68.22%
2020-11-04 19:13  ---  Trump leading by 461,765 votes  Remaining (est.): 1,307,396  Change:   4,397 (Trump 29.4% / 70.6% Biden)   Precincts: 7270/9128  Biden needs 67.66% [-0.010%]  Biden recent trend 74.26%
2020-11-04 18:56  ---  Trump leading by 463,576 votes  Remaining (est.): 1,311,793  Change:     494 (Trump 36.4% / 63.6% Biden)   Precincts: 7264/9128  Biden needs 67.67% [0.002%]   Biden recent trend 74.53%
2020-11-04 18:40  ---  Trump leading by 463,710 votes  Remaining (est.): 1,312,287  Change:   1,700 (Trump 42.5% / 57.5% Biden)   Precincts: 7263/9128  Biden needs 67.67% [0.013%]   Biden recent trend 74.63%
2020-11-04 18:23  ---  Trump leading by 463,965 votes  Remaining (est.): 1,313,987  Change:   4,754 (Trump 22.0% / 78.0% Biden)   Precincts: 7261/9128  Biden needs 67.65% [-0.037%]  Biden recent trend 75.14%
2020-11-04 18:15  ---  Trump leading by 466,625 votes  Remaining (est.): 1,318,741  Change:   4,780 (Trump 20.7% / 79.3% Biden)   Precincts: 7255/9128  Biden needs 67.69% [-0.042%]  Biden recent trend 74.88%
2020-11-04 17:44  ---  Trump leading by 469,423 votes  Remaining (est.): 1,323,521  Change:  10,854 (Trump 35.4% / 64.6% Biden)   Precincts: 7246/9128  Biden needs 67.73% [0.025%]   Biden recent trend 74.44%
2020-11-04 17:41  ---  Trump leading by 472,593 votes  Remaining (est.): 1,334,375  Change:     155 (Trump 40.3% / 59.7% Biden)   Precincts: 7237/9128  Biden needs 67.71% [0.001%]   Biden recent trend 77.40%
2020-11-04 17:24  ---  Trump leading by 472,623 votes  Remaining (est.): 1,334,530  Change:     219 (Trump 50.2% / 49.8% Biden)   Precincts: 7237/9128  Biden needs 67.71% [0.003%]   Biden recent trend 77.48%
2020-11-04 17:16  ---  Trump leading by 472,622 votes  Remaining (est.): 1,334,749  Change:  35,680 (Trump 22.4% / 77.6% Biden)   Precincts: 7237/9128  Biden needs 67.70% [-0.259%]  Biden recent trend 77.65%
2020-11-04 17:04  ---  Trump leading by 492,350 votes  Remaining (est.): 1,370,429  Change:  15,217 (Trump 27.8% / 72.2% Biden)   Precincts: 7193/9128  Biden needs 67.96% [-0.047%]  Biden recent trend 70.67%
2020-11-04 16:52  ---  Trump leading by 499,120 votes  Remaining (est.): 1,385,646  Change:      30 (Trump 48.3% / 51.7% Biden)   Precincts: 7181/9128  Biden needs 68.01% [0.000%]   Biden recent trend 69.96%
2020-11-04 16:45  ---  Trump leading by 499,121 votes  Remaining (est.): 1,385,676  Change:  13,349 (Trump 28.4% / 71.6% Biden)   Precincts: 7181/9128  Biden needs 68.01% [-0.034%]  Biden recent trend 69.97%
2020-11-04 16:41  ---  Trump leading by 504,890 votes  Remaining (est.): 1,399,025  Change:  20,482 (Trump 31.1% / 68.9% Biden)   Precincts: 7169/9128  Biden needs 68.04% [-0.012%]  Biden recent trend 74.13%
2020-11-04 16:32  ---  Trump leading by 512,636 votes  Remaining (est.): 1,419,507  Change:   6,321 (Trump 47.7% / 52.3% Biden)   Precincts: 7151/9128  Biden needs 68.06% [0.070%]   Biden recent trend 76.78%
2020-11-04 16:25  ---  Trump leading by 512,933 votes  Remaining (est.): 1,425,828  Change:  33,898 (Trump 18.7% / 81.3% Biden)   Precincts: 7140/9128  Biden needs 67.99% [-0.310%]  Biden recent trend 81.34%
2020-11-04 16:12  ---  Trump leading by 534,180 votes  Remaining (est.): 1,459,726  Change:  15,030 (Trump 28.7% / 71.3% Biden)   Precincts: 7115/9128  Biden needs 68.30% [-0.031%]  Biden recent trend 57.51%
2020-11-04 16:03  ---  Trump leading by 540,587 votes  Remaining (est.): 1,474,756  Change:     556 (Trump 49.4% / 50.6% Biden)   Precincts: 7102/9128  Biden needs 68.33% [0.007%]   Biden recent trend 55.21%
2020-11-04 16:01  ---  Trump leading by 540,594 votes  Remaining (est.): 1,475,312  Change:   4,467 (Trump 39.9% / 60.1% Biden)   Precincts: 7102/9128  Biden needs 68.32% [0.025%]   Biden recent trend 55.24%
2020-11-04 15:52  ---  Trump leading by 541,492 votes  Remaining (est.): 1,479,779  Change:   2,761 (Trump 26.8% / 73.2% Biden)   Precincts: 7097/9128  Biden needs 68.30% [-0.009%]  Biden recent trend 54.99%
2020-11-04 15:41  ---  Trump leading by 542,774 votes  Remaining (est.): 1,482,540  Change:  82,554 (Trump 45.6% / 54.4% Biden)   Precincts: 7097/9128  Biden needs 68.31% [0.735%]   Biden recent trend 54.38%
2020-11-04 15:25  ---  Trump leading by 550,001 votes  Remaining (est.): 1,565,094  Change:  44,605 (Trump  7.8% / 92.2% Biden)   Precincts: 7035/9128  Biden needs 67.57% [-0.683%]  Biden recent trend 92.21%
2020-11-04 15:17  ---  Trump leading by 587,656 votes  Remaining (est.): 1,609,699  Change:   4,275 (Trump 30.0% / 70.0% Biden)   Precincts: 6943/9128  Biden needs 68.25% [-0.005%]  Biden recent trend 82.36%
2020-11-04 14:56  ---  Trump leading by 589,367 votes  Remaining (est.): 1,613,974  Change:      88 (Trump 41.5% / 58.5% Biden)   Precincts: 6938/9128  Biden needs 68.26% [0.001%]   Biden recent trend 83.69%
2020-11-04 14:32  ---  Trump leading by 589,382 votes  Remaining (est.): 1,614,062  Change:  39,561 (Trump 16.3% / 83.7% Biden)   Precincts: 6938/9128  Biden needs 68.26% [-0.370%]  Biden recent trend 83.74%
2020-11-04 14:23  ---  Trump leading by 616,079 votes  Remaining (est.): 1,653,623  Change:   4,580 (Trump 19.9% / 80.1% Biden)   Precincts: 6905/9128  Biden needs 68.63% [-0.032%]  Biden recent trend 80.13%
2020-11-04 13:28  ---  Trump leading by 618,840 votes  Remaining (est.): 1,658,203  Change:       0 (n/a)                         Precincts: 6905/9128  Biden needs 68.66% [68.660%]  Biden recent trend n/a
----------------  ---  ------------------------------  ---------------------------  --------------------------------------------  --------------------  ----------------------------  -------------------------`

func main() {
	var re = regexp.MustCompile(`(?m)\S* \(EV: \d*\) Total Votes: .*`)
	res := []string{}
	for _, b := range re.FindAllStringSubmatch(text, -1){
		for _, c := range b{
			res = append(res, c)
		}
	}
	for _, s := range res{
		fmt.Println(s)
	}
}

func diff(text1, text2 string) []string {
	a := strings.Split(text1, "\n")
	amap := make(map[string]bool)
	b := strings.Split(text2, "\n")
	bmap := make(map[string]bool)
	for _, a2 := range a {
		amap[a2] = true
	}
	for _, b2 := range b {
		bmap[b2] = true
	}
	final := []string{}
	if len(bmap) < len(amap) {
	 bmap, amap = amap, bmap
	}
	for key, _ := range bmap {
		if amap[key] == false {
			final = append(final, key)
		}
	}
	return final
}
