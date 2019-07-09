#include "error.h"
#include <stdarg.h>
#include <stdio.h>
#include <bsd/stdlib.h>
#include <stdlib.h>
#include <string.h>
void error(char *fmt, ...) {
  va_list args;
  va_start(args, fmt);
  if (getprogname())
    fprintf(stderr, "error in %s: ", getprogname());
  else
    fprintf(stderr, "error in anonymous program: ");
  vfprintf(stderr, fmt, args);
  va_end(args);
  exit(1);
}
void *emalloc(size_t n) {
  void *p = malloc(n);
  if (!p)
    error("emalloc, can't allocate %ld bytes.\n", n);
  return p;
}
void *erealloc(void *p, size_t n) {
  p = realloc(p, n);
  if (!p)
    error("erealloc, can't allocate %ld bytes.\n", n);
  return p;
}
FILE *efopen(const char *name, const char *mode) {
  FILE *fp = fopen(name, mode);
  if (!fp)
    error("efopen, can't open %s\n", name);
  return fp;
}
FILE *epopen(const char *cmd, const char *type) {
  FILE *pp = popen(cmd, type);
  if (!pp)
    error("epopen, couldn't open %s\n", cmd);
  return pp;
}
DIR *eopendir(char *dir) {
  DIR *d = opendir(dir);
  if (!d)
    error("eopendir, couldn't open %s\n", dir);
  return d;
}
char *estrdup(const char *s) {
  char *p = strdup(s);
  if (!p)
    error("estrdup, couldn't duplicate a string.\n");
  return p;
}
char *estrndup(const char *s, size_t n) {
  char *p = strndup(s, n);
  if (!p)
    error("estrndup, couldn't copy a string.\n");
  return p;
}
