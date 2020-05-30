int factorial(int val) {
  if val==1 {
	return val;
  }
  int x;
  x = factorial(val - 1);
  return x * val;
}


int main(){
    int a = 2;
    a = a - 1;
    a = factorial(3);
    print(a);
    return 0;
}
