#include "matchLen.h"
#include <float.h>
#include <math.h>
#include <gsl/gsl_sf_gamma.h>
double prob(long l, double g, int x);
float erfinv(float x);
double meanMatchLen(long l, double g) {
  double cp = 0., m = 0.;
  for (int x = 1; x < l; x++) {
    double p = prob(l, g, x + 1) - prob(l, g, x);
    m += x * p;
    cp += p;
    if (cp >= 1. - DBL_EPSILON)
        return m;
  }
  return m;
}
double prob(long l, double g, int x) {
  l *= 2;
  g /= 2.;
  double s = 0.;
  for (int k = 0; k <= x; k++) {
    double l1, l2, l3, l4, l5;
    l1 = log(pow(2, x));
    l2 = gsl_sf_lnchoose(x, k);
    l3 = log(pow(g, k));
    l4 = log(pow(0.5 - g, x - k));
    l5 = log(pow(1. - pow(g, k) * pow(0.5 - g, x - k), l));
    s += exp(l1 + l2 + l3 + l4 + l5);
  }
  return s;
}
double varMatchLen(long l, double g) {
  double cp = 0., v = 0.;
  double m = meanMatchLen(l, g);
  for (int x = 1; x < l; x++) {
    double p = prob(l, g, x + 1) - prob(l, g, x);
    v += x * x * p;
    cp += p;
    if (cp >= 1. - DBL_EPSILON)
        return v - m * m;
  }
  return v - m * m;
}
double quantCm(long l, double g, int w, double p) {
  double m = meanMatchLen(l, g);
  double v = varMatchLen(l, g);
  v /= m * w;
  double r = 1. + sqrt(2 * v) * erfinv(2. * p - 1.);
  return r;
}
