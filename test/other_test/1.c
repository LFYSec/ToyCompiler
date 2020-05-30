#include <stdio.h>

int foo(int a, int b, int e) {
  int c, d;
  goto L1;

  L1: {
    c = a + b;
    d = c - a;
    if (d) goto L2;
    d = b * d;
    e = e + 1;
  }
  L2: {
    b = a + b;
    e = c - a;
    if (e) goto L1;
    a = b * d;
    b = a - d;
  }
  return b;
}
int main(){
  foo(1,2,3);
  return 0;
}
