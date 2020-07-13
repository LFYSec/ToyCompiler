int main() {
    int array[10];
    int i=0;
    int num = 10;
    while i<10 {
        array[i] = i;
        i = i+1;
    }
    i=0;
    while i<10 {
        print(array[i]);
        i = i+1;
    }
    return 0;
}
