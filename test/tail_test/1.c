int id(x) {
    return x;
}
int f(a) {
    const b = a + 1;
    return b * id(b);     //是尾调用
    // return b * id(b); 不是尾调用
}
int main() {
    f(1);
}
