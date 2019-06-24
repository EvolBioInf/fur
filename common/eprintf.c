/***** eprintf.c **************************************************
 * Description: Collection of functions for error handling.
 * Reference: Kernighan, B. W. and Pike, R. (1999). The Practice
 *            of programming. Addision Wesley; chapter 4.
 * Author: Bernhard Haubold, haubold@evolbio.mpg.de
 * Licence: Public domain.
 * File created on Fri Dec 17 11:16:34 2004.
 *****************************************************************/
#include <stdio.h>
#include <stdarg.h>
#include <string.h>
#include <errno.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include "eprintf.h"

/* efopen: open file and report if error */
FILE *efopen(char *fname, char *mode){
  FILE *fp;

  if((fp = fopen(fname, mode)) == NULL)
    eprintf("efopen(%s, %s) failed:",fname,mode);
  
  return fp;
}

/* eopen: open file on system level and report if error */
int eopen(char *fname, int flag){
  int fd;

  if((fd = open(fname, flag, 0)) < 0)
    eprintf("eopen(%s, %s) failed:",fname,flag);
  return fd;
}

/* eprintf: print error message and exit */
void eprintf(char *fmt, ...){
  va_list args;
  fflush(stdout);
  if(progname() != NULL)
    fprintf(stderr, "%s: ", progname());
  
  va_start(args, fmt);
  vfprintf(stderr, fmt, args);
  va_end(args);
  
  if(fmt[0] != '\0' && fmt[strlen(fmt)-1] == ':')
    fprintf(stderr, " %s", strerror(errno));
  fprintf(stderr, "\n");
  exit(2); /* conventional value for failed execution */
}

/* estrdup: duplicate a string, report if error */
char *estrdup(char *s){
  char *t;
  
  t = (char *)malloc(strlen(s)+1);
  if(t == NULL)
    eprintf("estrdup(\"%.20s\") failed:", s);
  strcpy(t, s);
  return t;
}

/* emalloc: malloc and report if error */
void *emalloc(size_t n){
  void *p;
  
  p = malloc(n);
  if(p == NULL)
    eprintf("malloc of %u bytes failed:", n);
  return p;
}

/* erealloc: realloc and report if error */
void *erealloc(void *p, size_t n){

  p = realloc(p, n);
  if(p == NULL)
    eprintf("realloc of %u bytes failed:", n);
  return p;
}

static char *name = NULL; /* program name for messages */

/* setprogname2: set stored name of program */
void setprogname2(char *str){
  name = estrdup(str);
}

/* progname: return stured name of program */
char *progname(void){
  return name;
}
