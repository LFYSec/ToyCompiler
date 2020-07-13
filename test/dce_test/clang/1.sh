clang -S -emit-llvm -Xclang -disable-O0-optnone mydce.c
opt -load /Users/lfy/LLVM/build/lib/LLVMMydce.dylib -mydce mydce.ll -o mydce.bc
llvm-dis mydce.bc -o -
