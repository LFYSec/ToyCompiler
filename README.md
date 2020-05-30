# Toy Compiler

## 整体设计

### 前端
整体采用go语言实现，前端使用lex+goyacc，生成ast，然后从ast翻译成LLVM IR。

### 后端
后端直接LLVM将LLVM IR编译为机器码。

## 进度

### 理论
看了一部分哈工大网课，熟悉基本词法语法分析知识。优化参考中科大网课，CFG学习参考龙书(TODO)

### 实验
已实现词法语法分析，能够生成LLVM IR，大体满足C0文法(太麻烦了

符号表、AST设计、interface多态等具体实现等参考代码

具体的测试用例在test/toylang_test文件夹中

- [x] variable
    - [x] int
    - [x] double
    - [x] char
    - [x] array
        - [x] one dimension
        - [ ] more
- [x] declare
   - [x] with init
   - [x] without init 
- [x] if
- [x] while
- [x] function
    - [x] define
    - [x] call
    - [x] recurse
- [x] print
- [x] scan
- [ ] const(global)

### 优化
参考LLVM CookBook，测试了dead code emit, early-cse等前端优化，未来加入寄存器相关使用llc的后端优化。

优化的测试用例在test/ssa_test, 还有很多TODO