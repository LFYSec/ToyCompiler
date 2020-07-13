clang -S -emit-llvm -Xclang -disable-O0-optnone expr.c
opt -load /Users/lfy/LLVM/build/lib/LLVMMydce.dylib -mydce expr.ll -o expr.bc
llvm-dis expr.bc -o -
