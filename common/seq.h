#ifndef SEQ
#define SEQ
#include <stdio.h>
#define DEFAULT_LINE_LEN 70
#define NUC_STR "ACGTUWSMKRYBDHVNacgtuwsmkrybdhvn"
#define CMP_STR "TGCAAWSKMYRVHDBNtgcaawskmyrvhdbn"
#define DNA_STR "ACGTacgt"
#define GC_STR "GCgc"
typedef struct seq {
  char *name;
  char *data;
  int l;
} Seq;
typedef struct seqArr {
  Seq **arr;
  int n;
} SeqArr;
Seq *newSeq(char *header);
void freeSeq(Seq *s);
SeqArr *newSeqArr();
void freeSeqArr(SeqArr *sa);
void seqArrAdd(SeqArr *sa, Seq *s);
Seq *getJoinedSeq(char *file);
Seq *getSeq(FILE *fp);
void printSeq(FILE *fp, Seq *s, int l);
void printSeqData(FILE *fp, Seq *s, int l);
Seq *revCmpSeq(Seq *s);
double gcSeq(Seq *s);
Seq *catSeq(Seq *s1, Seq *s2);
#endif
