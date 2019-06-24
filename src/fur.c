/***** fur.c ****************************************************************
 * Description: Find unique genomic regions
 * Author: Bernhard Haubold
 * Email: haubold@evolbio.mpg.de
 * License: GNU General Public License, https://www.gnu.org/licenses/gpl.html
 * Date: Mon Jun 24 10:50:06 2019
 ****************************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include "interface.h"
#include "eprintf.h"

void scanFile(FILE *fp, Args *args) {
  for(int i = 0; i < args->i; i++)
    printf("Test output.\n");
}

int main(int argc, char *argv[]){
  FILE *fp;
  Args *args = getArgs(argc, argv);

  setprogname2(argv[0]);
  if(args->v)
    printSplash(args);
  if(args->h || args->err)
    printUsage();
  if(args->nf == 0) {
    fp = stdin;
    scanFile(fp, args);
  } else {
    for(int i = 0; i < args->nf; i++) {
      fp = efopen(args->fi[i], "r");
      scanFile(fp, args);
      fclose(fp);
    }
  }
  freeArgs(args);
  free(progname());
  return 0;
}

