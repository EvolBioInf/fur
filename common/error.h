#include <stdio.h>
#include <dirent.h>
void error(char *fmt, ...);
void *emalloc(size_t n);
void *erealloc(void *p, size_t n);
FILE *efopen(const char *file, const char *mode);
FILE *epopen(const char *cmd, const char *type);
DIR *eopendir(char *dir);
char *estrdup(const char *s);
char *estrndup(const char *s, size_t n);
