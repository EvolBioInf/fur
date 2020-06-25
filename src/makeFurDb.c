#include <bsd/stdlib.h>
#include "mfdbI.h"
#include "seq.h"
#include <dirent.h>
#include <sys/types.h>
#include "error.h"
#include <string.h>
#include <stdlib.h>
void readSeq(SeqArr *sa, char *dir, char *file);
void readSeq(SeqArr *sa, char *dir, char *file) {
  char *path = emalloc(strlen(dir) + strlen(file) + 2);
  path[0] = '\0';
  strcat(path, dir);
  strcat(path, "/");
  strcat(path, file);
  seqArrAdd(sa, getJoinedSeq(path));
  free(path);
}
int main(int argc, char **argv) {
  setprogname(argv[0]);
  Args *args = getArgs(argc, argv);
  if (args->h || args->err)
    printUsage();
  if (args->v)
    printSplash(args);
  fprintf(stderr, "# Reading data...");
  SeqArr *ta, *ne;
  DIR *d;
  struct dirent *dir;
  ta = newSeqArr();
  d = eopendir(args->t);
  while ((dir = readdir(d)) != NULL)
    if (strcmp(dir->d_name, ".")  != 0 &&
          strcmp(dir->d_name, "..") != 0)
      readSeq(ta, args->t, dir->d_name);
  closedir(d);
  ne = newSeqArr();
  d = eopendir(args->n);
  while ((dir = readdir(d)) != NULL)
    if (strcmp(dir->d_name, ".")  != 0 &&
          strcmp(dir->d_name, "..") != 0)
      readSeq(ne, args->n, dir->d_name);
  closedir(d);
  fprintf(stderr, "done.\n");
  struct stat sb;
  if (stat(args->d, &sb) != -1) {
    if (!args->o)
      error("%s already exists.\n", args->d);
  } else {
    char cmd[1024];
    sprintf(cmd, "mkdir %s", args->d);
    if (system(cmd) < 0)
      error("couldn't run system command %s\n", cmd);
  }
  int r = 0;
  if (args->r) {
    r = -1;
    for (int i = 0; i < ta->n; i++)
      if (strstr(ta->arr[i]->name, args->r)) {
        if (r == -1)
            r = i;
        else
            error("%s is ambiguous.\n", args->r);
      }
    if (r == -1)
      error("couldn't find %s.\n", args->r);
  } else {
    int max = -1;
    for (int i = 0; i < ta->n; i++)
      if (max < ta->arr[i]->l) {
        max = ta->arr[i]->l;
        r = i;
      }
  }
  char *tmpl = "macle -s > %s/macle.idx", cmd[1024];
  sprintf(cmd, tmpl, args->d);
  FILE *pp = epopen(cmd, "w");
  fprintf(stderr, "# Making macle index with target representative \"%s\"...",
            ta->arr[r]->name);
  fprintf(pp, ">t%d\n%s\n", r, ta->arr[r]->data);
  for (int i = 0; i < ne->n; i++)
    fprintf(pp, ">n%d\n%s\n", i, ne->arr[i]->data);
  pclose(pp);
  fprintf(stderr, "done.\n");
  fprintf(stderr, "# Making BLAST database...");
  tmpl = "makeblastdb -parse_seqids -out %s/blastdb "
    "-dbtype nucl -title db > /dev/null";
  sprintf(cmd, tmpl, args->d);
  pp = epopen(cmd, "w");
  for (int i = 0; i < ta->n; i++)
    fprintf(pp, ">t%d\n%s\n", i, ta->arr[i]->data);
  for (int i = 0; i < ne->n; i++)
    fprintf(pp, ">n%d\n%s\n", i, ne->arr[i]->data);
  pclose(pp);
  fprintf(stderr, "done.\n");
  freeArgs(args);
  freeSeqArr(ta);
  freeSeqArr(ne);
}
