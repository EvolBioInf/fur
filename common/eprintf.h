/***** eprintf.h **************************************************
 * Description: Header file for eprintf, which provides error-
 *              handling capabilities.
 * Author: Bernhard Haubold, haubold@evolbio.mpg.de
 * License: Public domain.
 * File created on Fri Dec 17 11:16:37 2004.
 *****************************************************************/
#ifndef EPRINTF
#define EPRINTF
#include <stdio.h>

extern FILE *efopen(char *fname, char *mode); 
extern int eopen(char *fname, int flag);
extern void eprintf(char *, ...);
extern char *estrdup (char *);
extern void *emalloc(size_t);
extern void *erealloc(void *, size_t);
extern char *progname(void);
extern void setprogname2(char *);

#endif
