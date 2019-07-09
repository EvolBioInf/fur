/***** interface.h **********************************************************
 * Description: Find unique genomic regions
 * Author: Bernhard Haubold
 * Email: haubold@evolbio.mpg.de
 * License: GNU General Public License, https://www.gnu.org/licenses/gpl.html
 * Date: Mon Jun 24 10:50:06 2019
 ****************************************************************************/
#ifndef INTERFACE
#define INTERFACE

#define DEFAULT_W 500
#define DEFAULT_P 0.95

/* define argument container */
typedef struct args{
  char    h; /* help message?         */
  char    v; /* version message?      */
  char  err; /* error                 */
  char **fi; /* input files           */
  int    nf; /* number of input files */

  char   *t; /* target directory      */
  char   *n; /* neighborhood dir.     */
  char   *i; /* macle index           */
  int     w; /* window length         */
  float   p; /* p-value               */

} Args;

Args *getArgs(int argc, char *argv[]);
Args *newArgs();
void freeArgs(Args *args);
void printUsage();
void printSplash(Args *args);

#endif
