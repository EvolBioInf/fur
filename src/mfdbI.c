/***** interface.c **********************************************************
 * Description: Find unique genomic regions
 * Author: Bernhard Haubold
 * Email: haubold@evolbio.mpg.de
 * License: pending.
 * Date: Mon Jun 24 10:50:06 2019
 ****************************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>
#include <bsd/stdlib.h>
#include "mfdbI.h"
#include "error.h"

Args *newArgs() {
  Args *args = (Args *)emalloc(sizeof(Args));
  args->h   = 0;
  args->v   = 0;
  args->err = 0;
  args->o   = 0;
  args->t   = NULL;
  args->n   = NULL;
  args->d   = NULL;
  args->r   = NULL;
  return args;
}

void freeArgs(Args *args) {
  free(args->t);
  free(args->n);
  free(args->d);
  if (args->r)
    free(args->r);
  free(args);
}

Args *getArgs(int argc, char *argv[]) {
  int c;
  char *optString = "hvot:n:d:r:";
  Args *args = newArgs();

  while ((c = getopt(argc, argv, optString)) != -1) {
    switch(c) {
    case 't': /* targets */
      args->t = estrdup(optarg);
      break;
    case 'n': /* neighborhood */
      args->n = estrdup(optarg);
      break;
    case 'd': /* database name */
      args->d = estrdup(optarg);
      break;
    case 'r': /* target representative to */
      args->r = estrdup(optarg);
      break;
    case 'o': /* overwrite database directory? */
      args->o = 1;
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
  if ((!args->t || !args->n || !args->d) && !args->h && !args->v) {
    printf("ERROR[makeFurDb]: Please specify target, neighborhood, and database.\n");
    args->err = 1;
  }
  return args;
}

void printUsage() {
  printf("Usage: %s [options]\n", getprogname());
  printf("Find uniqe genomic regions\n");
  printf("Example: %s -t target -n neighborhood -d furDb\n", getprogname());
  printf("Options:\n");
  printf("\t-t <STR> target directory\n");
  printf("\t-n <STR> neighborhood directory\n");
  printf("\t-d <STR> database\n");
  printf("\t[-r <STR> target representative; default: longest target sequence]\n");
  printf("\t[-o overwrite existing database directory; default: don't overwrite]\n");
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
  printf("License: Pending\n");
  printf("Bugs: haubold@evolbio.mpg.de\n");
  exit(0);
}
