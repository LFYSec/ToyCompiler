opt -load /Users/lfy/LLVM/build/lib/LLVMMyinline.dylib -myinliner inline.ll -o inline.bc
llvm-dis inline.bc -o -
