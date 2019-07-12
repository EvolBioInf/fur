/***** interface.c **********************************************************
 * Description: Find unique genomic regions
 * Author: Bernhard Haubold
 * Email: haubold@evolbio.mpg.de
 * License: GNU General Public License, https://www.gnu.org/licenses/gpl.html
 * Date: Mon Jun 24 10:50:06 2019
 ****************************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>
#include <bsd/stdlib.h>
#include "interface.h"
#include "error.h"

Args *newArgs() {
  Args *args = (Args *)emalloc(sizeof(Args));
  args->h   = 0;
  args->v   = 0;
  args->err = 0;
  args->t   = NULL;
  args->n   = NULL;
  args->i   = NULL;
  args->I   = NULL;
  args->w   = DEFAULT_W;
  args->p   = DEFAULT_P;
  args->T   = DEFAULT_TT;
  return args;
}

void freeArgs(Args *args) {
  free(args->t);
  free(args->n);
  if (args->i)
    free(args->i);
  free(args);
}

Args *getArgs(int argc, char *argv[]) {
  int c;
  char *optString = "hvt:n:i:I:w:p:T:";
  Args *args = newArgs();

  while ((c = getopt(argc, argv, optString)) != -1) {
    switch(c) {
    case 't': /* targets */
      args->t = estrdup(optarg);
      break;
    case 'n': /* neighborhood */
      args->n = estrdup(optarg);
      break;
    case 'i': /* index file */
      args->i = estrdup(optarg);
      break;
    case 'I': /* name of file index is written to */
      args->I = estrdup(optarg);
      if (access(args->I, F_OK) != -1)
	error("%s already exists, please chose another name for the index file.\n", args->I);
      break;
    case 'w': /* window length */
      args->w = atoi(optarg);
      break;
    case 'p': /* p-value */
      args->p = atof(optarg);
      break;
    case 'T': /* number of threads */
      args->T = atoi(optarg);
      break;
    case 'h': /* help       */
      args->h = 1;
      break;
    case 'v': /* version    */
      args->v = 1;
      break;
    case '?':
      args->err = 1;
      if(optopt == 'i')
	fprintf(stderr, "Error: Option `%c` requires an argument.\n", optopt);
      else if(isprint(optopt))
	fprintf(stderr, "Error: Unknown option `%c`.\n", optopt);
      else
	fprintf(stderr, "Error: Unknown option character `\\x%x`.\n", optopt);
    default:
      args->err = 1;
      return args;
    }
  }
  args->fi = argv + optind;
  args->nf = argc - optind;
  if (!args->t || !args->n) {
    printf("ERROR[fur]: Please specify a target and a neighborhood directory.\n");
    args->err = 1;
  }
  return args;
}

void printUsage() {
  printf("Usage: %s [options] [inputFiles]\n", getprogname());
  printf("Find uniqe genomic regions\n");
  printf("Example: %s -t target -n neighborhood\n", getprogname());
  printf("Options:\n");
  printf("\t-t <STR> target directory\n");
  printf("\t-n <STR> neighborhood directory\n");
  printf("\t[-i <STR> macle index; default: computed from scratch (slow)]\n");
  printf("\t[-I <STR> name of file to write macle index to; default: generated internally]\n");
  printf("\t[-w <NUM> window length; default: %d]\n", DEFAULT_W);
  printf("\t[-p <NUM> p-value for uniqueness; default: %g]\n", DEFAULT_P);
  printf("\t[-T <NUM> number of threads in BLAST search; default: %d]\n", DEFAULT_TT);
  printf("\t[-h print this help message and exit]\n");
  printf("\t[-v print version & program information and exit]\n");
  exit(0);
}

void printSplash(Args *args) {
  printf("%s ", getprogname());
  int l = strlen(VERSION);
  for(int i = 0; i < l - 1; i++)
    printf("%c", VERSION[i]);
  printf(", %s\n", DATE);
  printf("Author: Bernhard Haubold\n");
  printf("License: GNU General Public License, https://www.gnu.org/licenses/gpl.html\n");
  printf("Bugs: haubold@evolbio.mpg.de\n");
  exit(0);
}
