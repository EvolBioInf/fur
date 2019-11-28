/***** interface.c **********************************************************
 * Description: Find unique genomic regions
 * Author: Bernhard Haubold
 * Email: haubold@evolbio.mpg.de
 * License: Pending
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
  args->d   = NULL;
  args->u   = 0;
  args->w   = DEFAULT_W;
  args->p   = DEFAULT_P;
  args->t   = DEFAULT_T;
  args->i   = DEFAULT_I;
  args->l   = DEFAULT_L;
  args->u   = 0;
  return args;
}

void freeArgs(Args *args) {
  free(args->d);
  free(args);
}

Args *getArgs(int argc, char *argv[]) {
  int c;
  char *optString = "hvud:w:p:t:i:l:";
  Args *args = newArgs();

  while ((c = getopt(argc, argv, optString)) != -1) {
    switch(c) {
    case 'd': /* database */
      args->d = estrdup(optarg);
      break;
    case 'w': /* window length */
      args->w = atoi(optarg);
      break;
    case 'p': /* p-value */
      args->p = atof(optarg);
      break;
    case 't': /* number of threads */
      args->t = atoi(optarg);
      break;
    case 'i': /* minimum percent identity */
      args->i = atof(optarg);
      break;
    case 'l': /* minimum length of alignment */
      args->l = atoi(optarg);
      break;
    case 'u': /* print unique regions */
      args->u = 1;
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
  if (!args->d && !args->h && !args->v) {
    printf("ERROR[fur]: Please specify a fur database.\n");
    args->err = 1;
  }
  return args;
}

void printUsage() {
  printf("Usage: %s [options] [inputFiles]\n", getprogname());
  printf("Find uniqe genomic regions\n");
  printf("Example: %s -d furDb\n", getprogname());
  printf("Options:\n");
  printf("\t-d <STR> database\n");
  printf("\t[-w <NUM> window length; default: %d]\n", DEFAULT_W);
  printf("\t[-p <NUM> p-value for uniqueness; default: %g]\n", DEFAULT_P);
  printf("\t[-i <NUM> minimum percent identity in target; default: %.3f]\n", DEFAULT_I);
  printf("\t[-l <NUM> minimum alignment length in target; default: query lengt]\n");
  printf("\t[-T <NUM> number of threads in BLAST search; default: %d]\n", DEFAULT_T);
  printf("\t[-u print unique regions and exit]\n");
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
