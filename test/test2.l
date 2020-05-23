int factorial(int val, int total) {
  if val==1 {
	return total;
  }
  return factorial(val-1, val * total);
}

int main(){
    print(factorial(3,1));
    return 0;
}
