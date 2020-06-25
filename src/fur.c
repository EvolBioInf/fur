#include "seq.h"
#include "error.h"
#include <bsd/stdlib.h>
#include "interface.h"
#include "matchLen.h"
#include <string.h>
#include <limits.h>
typedef struct intv {
  int s, e;
} Intv;
typedef struct intvArr {
  Intv **arr;
  int n;
} IntvArr;
Intv *newIntv(int s, int e);
IntvArr *newIntvArr();
void freeIntvArr(IntvArr *ia);
void intvArrAdd(IntvArr *ia, Intv *i);
Intv *newIntv(int s, int e) {
  Intv *i = (Intv *)emalloc(sizeof(Intv));
  i->s = s;
  i->e = e;
  return i;
}
IntvArr *newIntvArr() {
  IntvArr *ia = (IntvArr *)emalloc(sizeof(IntvArr));
  ia->arr = NULL;
  ia->n = 0;
  return ia;
}
void freeIntvArr(IntvArr *ia) {
  for (int i = 0; i < ia->n; i++)
    free(ia->arr[i]);
  free(ia->arr);
  free(ia);
}
void intvArrAdd(IntvArr *ia, Intv *i) {
  ia->arr = (Intv **)
    erealloc(ia->arr, (ia->n + 1) * sizeof(Intv *));
  ia->arr[ia->n++] = i;
}
int main(int argc, char **argv) {
  setprogname(argv[0]);
  Args *args = getArgs(argc, argv);
  if (args->h || args->err)
    printUsage();
  if (args->v)
    printSplash(args);
  char rn[256];
  Seq *rs = NULL;
  char *tmpl = "macle -l %s/macle.idx | "
    "head -n 6 | tail -n 1 | "
    "awk '{print $6 }'";
  char cmd[1024];
  sprintf(cmd, tmpl, args->d);
  FILE *pp = epopen(cmd, "r");
  if (fscanf(pp, "%s", rn) == EOF)
    error("couldn't run %s\n", cmd);
  pclose(pp);
  tmpl = "blastdbcmd -entry %s -db %s/blastdb";
  sprintf(cmd, tmpl, rn, args->d);
  pp = epopen(cmd, "r");
  Seq *sp;
  while ((sp = getSeq(pp)) != NULL)
    rs = sp;
  pclose(pp);
  double mc, gc = 0.;
  long len = 0;
  IntvArr *ia;
  tmpl = "macle -l %s/macle.idx | "
    "tail -n +2 | "
    "awk '{print $2}'";
  sprintf(cmd, tmpl, args->d);
  pp = epopen(cmd, "r");
  if (fscanf(pp, "%ld", &len) == EOF)
    error("couldn't run %s\n", cmd);
  if (fscanf(pp, "%lf", &gc) == EOF)
    error("couldn't run %s\n", cmd);
  mc = quantCm(len, gc, args->w, args->p);
  pclose(pp);
  tmpl =
    "macle -i %s/macle.idx -n %s -w %d -k %d | "
    "cut -f 2,3 | "
    "awk '$2 > -1'";
  sprintf(cmd, tmpl, args->d, rn, args->w, args->k);
  pp = epopen(cmd, "r");
  float m, c;
  ia = newIntvArr();
  int is, ie, in = 0;
  while (fscanf(pp, "%f %f", &m, &c) != EOF) {
    int ws = m - args->w / 2;
    int we = m + args->w / 2 - 1;
    if (in) {
      if (ws <= ie && c >= mc)
        ie = we;
      else if (ws > ie) {
        in = 0;
        intvArrAdd(ia, newIntv(is, ie));
      }
    } else {
      if (c >= mc) {
        in = 1;
        is = m - args->w / 2;
        ie = m + args->w / 2 - 1;
      }
    }
  }
  pclose(pp);
  int nn = 0, nm = 0;
  for (int i = 0; i < ia->n; i++)
    for (int j = ia->arr[i]->s; j <= ia->arr[i]->e; j++)
      if (rs->data[j] == 'N')
          nm++;
      else
          nn++;
  char *h1 = "# Step                    Sequences  Nucleotides  "
    "Mutations (N)";
  char *h2 = "# ------------------------------------------------"
    "-------------";
  fprintf(stderr, "%s\n%s\n", h1, h2);
  tmpl = "# Sliding window             %6d     %8d         %6d\n";
  fprintf(stderr, tmpl, ia->n, nn, nm);
  SeqArr *sa = newSeqArr();
  char name[1024];
  for (int i = 0; i < ia->n; i++) {
    Intv *iv = ia->arr[i];
    sprintf(name, "template_%d %d-%d\n", i + 1, iv->s + 1,
              iv->e + 1);
    Seq *s = newSeq(name);
    s->data = emalloc(iv->e - iv->s + 2);
    for (int j = iv->s; j <= iv->e; j++)
      s->data[s->l++] = rs->data[j];
    s->data[s->l] = '\0';
    seqArrAdd(sa, s);
  }
  freeIntvArr(ia);
  if (args->u) {
    for (int i = 0; i < sa->n; i++)
      printSeq(stdout, sa->arr[i], -1);
    exit(0);
  }
  tmpl = "%s/r.fasta";
  sprintf(name, tmpl, args->d);
  FILE *fp = efopen(name, "w");
  for (int i = 0; i < sa->n; i++)
    printSeq(fp, sa->arr[i], -1);
  fclose(fp);
  in = 1;
  tmpl = "blastdbcmd -entry all -db %s/blastdb | sed 's/ $//'";
  sprintf(cmd, tmpl, args->d);
  pp = epopen(cmd, "r");
  while ((sp = getSeq(pp)) != NULL) {
    if (sp->name[0] == 't' && strcmp(sp->name, rn) != 0)  {
      sprintf(name, "%s/t%d.fasta", args->d, in++);
      fp = efopen(name, "w");
      printSeq(fp , sp, -1);
      fclose(fp);
    }
    freeSeq(sp);
  }    
  pclose(pp);
  tmpl = "phylonium -p %s/p.fasta -r %s/r.fasta %s/*.fasta "
    "> /dev/null 2> /dev/null";
  sprintf(cmd, tmpl, args->d, args->d, args->d);
  if (system(cmd) < 0)
    error("couldn't run system call %s\n", cmd);
  sprintf(name, "%s/p.fasta", args->d);
  fp = efopen(name, "r");
  freeSeqArr(sa);
  sa = newSeqArr();
  while ((sp = getSeq(fp)) != NULL)
    if (sp->l >= args->n)
      seqArrAdd(sa, sp);
    else
      freeSeq(sp);
  fclose(fp);
  for (int i = 0; i < sa->n; i++) {
    char *h = strstr(sa->arr[i]->name, ")");
    h += 2;
    int j = atoi(strtok(h, " "));
    if (j == 0) {
      continue;
    } else if (args->x) {
      freeSeq(sa->arr[i]);
      sa->arr[i] = NULL;
      continue;
    }
    char *t = strtok(NULL, " ");
    while (t != NULL) {
      int p = atoi(t) - 1;
      sa->arr[i]->data[p] = 'N';
      t = strtok(NULL, " ");
    }
  }
  tmpl = "rm %s/*.fasta";
  sprintf(cmd, tmpl, args->d);
  if (system(cmd) < 0)
    error("couldn't run system call %s\n", cmd);
  int ns = 0;
  nn = nm = 0;
  for (int i = 0; i < sa->n; i++) {
    if (!sa->arr[i]) continue;
    ns++;
    for (int j = 0; j < sa->arr[i]->l; j++)
      if (sa->arr[i]->data[j] == 'N') nm++;
      else nn++;
  }
  tmpl = "# Presence in targets        %6d     %8ld         %6d\n";
  fprintf(stderr, tmpl, ns, nn, nm);
  if (args->U) {
    for (int i = 0; i < sa->n; i++)
      if (sa->arr[i])
          printSeq(stdout, sa->arr[i], -1);
    exit(0);
  }
  tmpl = "blastn -task blastn -db %s/blastdb -num_threads %d "
    "-evalue %e -outfmt \"6 sacc qacc qstart qend\" "
    "| grep '^n' > "
    "%s/o.txt";
  sprintf(cmd, tmpl, args->d, args->t, args->e, args->d);
  pp = epopen(cmd, "w");
  for (int i = 0; i < sa->n; i++)
    if (sa->arr[i])
      fprintf(pp, ">%d\n%s\n", i, sa->arr[i]->data);
  pclose(pp);
  tmpl = "%s/o.txt";
  sprintf(name, tmpl, args->d);
  fp = efopen(name, "r");
  int *start = emalloc(sa->n * sizeof(int));
  int *end   = emalloc(sa->n * sizeof(int));
  for (int i = 0; i < sa->n; i++) {
    start[i] = INT_MAX;
    end[i]   = -1;
  }
  int ii, qs, qe;
  char s[32];
  while (fscanf(fp, "%s %d %d %d", s, &ii, &qs, &qe) != EOF) {
    if (qs < start[ii])
      start[ii] = qs ;
    if (qe > end[ii])
      end[ii] = qe;
  }
  fclose(fp);
  for (int i = 0; i < sa->n; i++) {
    int l = end[i] - start[i] + 1;
    if (l > 0 && args->x) {
      freeSeq(sa->arr[i]);
      sa->arr[i] = NULL;
      continue;
    }
    for (int j = start[i] - 1; j < end[i]; j++)
      sa->arr[i]->data[j] = 'N';
  }
  tmpl = "rm %s/o.txt";
  sprintf(cmd, tmpl, args->d);
  if (system(cmd) < 0) {
    fprintf(stderr, "couldn't run system call %s\n", cmd);
    exit(0);
  }
  free(start);
  free(end);
  nn = nm = ns = 0;
  for (int i = 0; i < sa->n; i++) {
    if (!sa->arr[i]) continue;
    int cn = 0, cm = 0;
    for (int j = 0; j < sa->arr[i]->l; j++)
      if (sa->arr[i]->data[j] == 'N')
        cm++;
      else
        cn++;
    if (cn >= args->n) {
      ns++;
      nm += cm;
      nn += cn;
    } else {
      freeSeq(sa->arr[i]);
      sa->arr[i] = NULL;
    }
  }
  tmpl = "# Absence from neighbors     %6d     %8ld         %6d\n";
  fprintf(stderr, tmpl, ns, nn, nm);
  for (int i = 0; i < sa->n; i++)
    if (sa->arr[i])
      printSeq(stdout, sa->arr[i], -1);
  freeSeqArr(sa);
  freeArgs(args);
  freeSeq(rs);
}
