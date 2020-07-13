clang -S -emit-llvm -Xclang -disable-O0-optnone ssa.c
opt -sroa ssa.ll -o ssa.bc
llvm-dis ssa.bc -o -
