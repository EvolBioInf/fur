/***** mfdbI.h **************************************************************
 * Description: Find unique genomic regions
 * Author: Bernhard Haubold
 * Email: haubold@evolbio.mpg.de
 * License: GNU General Public License, https://www.gnu.org/licenses/gpl.html
 * Date: Mon Jun 24 10:50:06 2019
 ****************************************************************************/
#ifndef INTERFACE
#define INTERFACE

/* define argument container */
typedef struct args{
  char    h; /* help message?         */
  char    v; /* version message?      */
  char    o; /* overwrite db dir?     */
  char  err; /* error?                */
  char **fi; /* input files           */
  int    nf; /* number of input files */

  char   *t; /* target directory      */
  char   *n; /* neighborhood dir.     */
  char   *d; /* database              */
  char   *r; /* target representative */
} Args;

Args *getArgs(int argc, char *argv[]);
Args *newArgs();
void freeArgs(Args *args);
void printUsage();
void printSplash(Args *args);

#endif
