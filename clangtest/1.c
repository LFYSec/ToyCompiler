#include <stdio.h>

int b(int a){
int c = a;
printf("%d",c);
return 1;
}

int b2(int a, int b){
    int c = a;
    int d = b;
    printf("%d, %d", c, d);
    return 1;
}
int factorial(int val, int total) {
  if(val==1) return total;
  return factorial(val-1, val * total);
}

int main(){
  //  int a = b(1);
  //  int c = 2;
  //  printf("%d",a);
  //  printf("%d",b2(c,3));
    printf("%d",factorial(3,1));
    return 0;
}
/*
int main(){
    int a;
    scanf("a:%d", &a);
    printf("1");
    printf("%d",a);
}*/

