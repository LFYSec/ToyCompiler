int factorial(int val) {
  if val==1 {
	return val;
  }
  int x = factorial(val-1);
  return x;
}


int main(){
    int a;
    a = factorial(3);
    print(a);
    return 0;
}
