# LFYLang Compiler Report

[toc]

## 1. 编译器整体架构设计及总体情况
目前LLVM比较火，因此就直接用了LLVM，还能顺便学一波LLVM Pass
 
### 前端
整体采用go语言实现，前端使用lex+goyacc，生成ast，然后手工将ast翻译成LLVM IR。

### 后端
后端直接使用LLVM套件llc将LLVM IR编译为机器码。

### 总代码规模
编译器代码总共2716行，前端LLVM Pass优化代码未计入。
![](https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/07/13/15946411975425.jpg)

### 开发环境及工具使用
IDE使用GoLand+XCode，主要采用Go+C语言实现
工具使用lex+goyacc

### 优化
- [x] 前端Pass
	- [x] SSA
	- [x] 死代码删除
	- [x] inline转换
- [x] 尾调用优化(依赖本身的Pass)

数据流分析、基本块获取使用LLVM API，

## 2. 文法定义
```
<program>::={<statement>}

// 语句
<statement>::={<returnStatement>|<assignStatement>|<printStatement>|<declareStatement>|<block>|<whileStatement>|<funcDefineStmt>|<funCallStmt>|<assignFunCallStmt>|<ifStatement>|<scanStatement>}

// 函数定义
funcDefineStmt::=<TYPE><IDENTIFY><LPARENTHESIS><defargs><RPARENTHESIS><block>
defargs::=<TYPE><IDENTIFY>{<COMMA><TYPE><IDENTIFY>}
block::=<LBRACE><stmtList><RBRACE>
stmtList::={<statement>}
returnStatement::=<RETURN><expression><EOS>

// 函数调用
funCallStmt::=<IDENTIFY><LPARENTHESIS><callargs><RPARENTHESIS>[<EOS>]
callargs::=<expression>{<COMMA><expression>}
assignFunCallStmt::=[<refExpression>|<TYPE><IDENTIFY>]<ASSIGN><funCallStmt><EOS>

// 表达式
expression::=<binaryOrAtomExpression>
binaryOrAtomExpression::=<unaryExpression>{<ALG><unaryExpression>}
unaryExpression::=<INT_LITERAL>|<DOUBLE_LITERAL>|<refExpression>|<LPARENTHESIS><expression><RPARENTHESIS>|<CHAR_LITERAL>
refExpression::=<IDENTIFY>[<LBRACKET><expression><RBRACKET>]

// 变量声明&定义
declareStatement::<TYPE><IDENTIFY>[<assign>|<LBRACKET><INT_LITERAL><RBRACKET>]<EOS>
assign::=<ASSIGN><expression>
assignStatement::=<refExpression><assign><EOS>

// 循环语句
whileStatement::=<WHILE><expression><block>

// if
ifStatement::=<IF><expression><block>[<ELSE><block>]

// 输入&输出
scanStatement::=<SCAN><LPARENTHESIS><expression><RPARENTHESIS><EOS>
printStatement::=<PRINT><LPARENTHESIS><expression><RPARENTHESIS><EOS> 
```
数据类型：int double char bool array
支持语句：while if return scan print call/define etc.

## 3. lex词法分析
词法分析部分使用lex。lex通过定义正则表达式对输入流进行匹配，将数据分隔成一个个的标记token，从而得到输入对应的具体类型。

具体的关键字列表如下：
```
// 变量
IDENTIFY::=[a-zA-Z][0-9a-zA-Z_]*

// 字面量
INT_LITERAL::=([-])?[0-9]+ // 匹配0/1次-号以及多次0-9数字
DOUBLE_LITERAL::=([-])?[0-9]+\.[0-9]*
CHAR_LITERAL::=\"(\\.|[^"])\" // 匹配双引号中的除双引号或者有\转义的符号

// 关键字
void char int double bool if else while print scan return
```

在go语言中使用lex还需要用到cgo模块，需要用go来调用c语言api，在文件lex.go中可以看到Lex方法对c中yylex等api的调用。通过对c api返回数据的分析可以返回在go语言中变量对应的类型，从而使接下来的语法分析可以使用对应的结构体等进行存储或解析。

## 4. goyacc语法分析
具体的文法定义开头已经给出。
goyacc是yacc的go实现，使用LALR(1)对经过lex解析过的输入流进行语法分析。因此前段解析的整体逻辑是lex.l -> lex.c -> lex.go -> grammer.y，lex.c是根据lex.l生成的，lex.go通过调用.c来返回相应逻辑到grammer.y中。grammer.y通过goyacc编译成y.go。y.go的作用其实就是语法分析。

**yacc语法说明**
yacc通过%%将整个语法定义分为三部分，整体结构如下：
```
%{
import文件
全局变量声明
%}

%%

%token/%union声明

%%

语法定义
```

%union中定义各种type，供之后的type、token节点指定类型。
%token定义各种词法分析中的token，比如可以通过\<int\>来指定类型为int。
%type指定各种语句对应的类型，用于$$返回。这里都指定为父结构体Node类型，然后在各个语句中再通过断言向下转型回之前的具体类型。

**注意：**goyacc测试过程中，具体语法匹配时不能出现比如","这种，比如把","在lex中返回具体的token值，在grammer中定义具体token才行。

## 5. 具体设计

### 5.1 语法树设计
整体语法树设计如下:
```
Node
    Expression
        RValue
            IntLiteral
            DoubleLiteral
            CharLiteral
        LValue
            VariableReference
            ArrayReference
    Statement
        CompoundStmt
        WhileStmt
        ScanStmt
        PrintStmt
        AssignStmt
        ...
```
语法树的根节点是Node interface，Node的子节点为表达式和语句。在grammer.y中，语法分析匹配到的语句规则的返回值都是Node类型结构体，这就涉及到了go中interface和struct的特性。

在c语言中，多态可以通过函数指针+结构体来实现，通过结构体包含，并将每个语句的具体函数绑定到对应结构体的函数指针来实现Node向下转型，从而实现语法树的生成。而在go语言中，并没有函数指针这一概念，因此需要使用go中的struct和interface来实现多态，并通过interface的断言和反射功能来实现向下转型。

以Statement接口举例，Stmt interface定义了一个GeneCode方法，只要实现了GeneCode方法的结构体，就可以算作实现了Stmt接口，这也是go语言中通过interface实现多态的具体应用。
```go
type Stmt interface {
	ast.Node
	GeneCode()
}
```

以CompoundStmt的具体结构体定义举例，CompoundStmt结构体继承了Stmt也就是语句结构体，并实现了绑定到结构体的GeneCode方法，这样也就实现了Stmt接口。当生成代码时，只需要调用GeneCode方法就可以通过interface的多态调用具体语句结构体的代码生成方法。
```go
type CompoundStmt struct {
	Stmt
	ChildCount 	int
	Stmts 		[]Stmt
	NamespaceId	int
}

func (c *CompoundStmt) GeneCode() {
	for _, s := range c.Stmts {
		s.GeneCode()
	}
}
```

再举例Node结构体的向下转型。如下语句在grammer.y中，首先通过空匹配将CompoundStmt结构体转型为Node向上返回，然后通过go语言的断言，将$1匹配到的Node结构体通过go的断言恢复为原来的CompoundStmt结构体，这样也就实现了语法树中类型的转换。
```go
program
    : {
        result = stmt.CreateCompoundStmt()
        $$ = result
    }
    | program statement {
        v1 := ($1).(*stmt.CompoundStmt)
        v2 := ($2).(stmt.Stmt)
        stmt.AddStmt(v1, v2)
    }
    ;
```

### 5.2 符号表设计
符号表分为变量和函数符号表，具体定义如下。
```go
type Symbol struct {
	Type            global.SymbolType // 变量类型
	Name			    string            // 变量名
	NamespaceId		 int               // 作用域
	Size			    int               // 数组size
}

type FuncSymbol struct {
	Name 			   string              // 函数名
	ReturnType	 	global.SymbolType   // 函数返回类型
	ArgsType 		map[string]global.SymbolType // 函数参数列表，用map存
}
```
作用域整体采用栈实现。当语法分析检测到进入block时，就调用PushFrame()函数，将当前的FrameId加一，并创建新的Frame，这样就可以实现具体的函数或变量作用域。Frame存在Stack结构体中，同时Stack符号栈中还存有函数符号表。
```go
type Stack struct {
	Length 		 int
	Funcs		    []*FuncSymbol
	Frames	 	 []*Frame
}
```

## 6. 代码生成
代码生成也是编译器前端的最后一步了，回顾一下解析的整体流程：
`词法分析 -> 语法分析 -> AST -> GeneCode()`
因为我们已经成功生成了AST，并且在Node interface中定义了GeneCode方法，所以我们的语句只要按照规范实现GeneCode方法，然后在最后通过遍历Stmt[]中的元素，通过多态调用对应的GeneCode()方法就可以很方便的做到代码生成。
拿变量定义的代码生成举例，通过输出alloca指令来实现declare语句生成。同时，通过go语言反射来判断Stmt.Initial是否为AssignStmt{}，如果不是空的Assign结构体，就说明这是一个声明+赋初值的语句，这时就要调用assign语句的GeneCode方法。
```go
func (d DeclareStmt) GeneCode() {
	fmt.Printf("%%%s_%d = alloca %s\n", d.Variable.Name, d.Variable.NamespaceId, symbolTable.TypeName(d.Variable))
	if !reflect.DeepEqual(d.Initial, AssignStmt{}) {
		d.Initial.GeneCode()
	}
}
```

## 7. 优化
我的编译器采用的是LLVM IR，因此整体的优化是要通过LLVM Pass来做的。这里需要查阅学习英文文档，补充很多LLVM的知识。我实现了如下的优化Pass
- [x] 前端Pass
	- [x] SSA
	- [x] 死代码删除
	- [x] inline转换
- [x] 尾调用优化(依赖本身的Pass)

首先介绍LLVM Pass的概念。LLVM的架构如下，前端通过Clang等编译器生成LLVM IR，然后LLVM的优化器对IR进行前端优化，接下来通过llc将IR生成为对应平台的汇编。这里优化器就是通过Pass实现的。Pass是一种类似流水线的处理机制，通过编写一系列的Pass，使这些Pass按顺序作用于IR，来产生对IR一步步优化的效果。
![](https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/14/15921457699421.jpg)

Pass分为两类, 一类是分析(analysis)Pass, 负责收集信息供其它Pass使用, 辅助调试或使程序可视化; 另一类是变换(transform)Pass, 改变程序的dataflow / controlflow。 LLVM中实现了几十种优化pass, 其中许多pass运行不止一次。analysis pass存放在lib/Analysis下, transform pass存放在lib/Transforms下。

这里拿我实现的死代码删除(DCE) Pass为例。因为LLVM API非常复杂，因此只挑核心实现举例。  
MyDCE Pass的设计思路是遍历函数中的所有指令集，并将指令集中对后续调用有影响的指令加入set中，将无影响的指令删除即可完成dce。首先明确Pass是作用于函数内部指令的Pass，因此需要集成Function模块，runOnFunction方法会在LLVM处理到函数时进行调用。Function中的核心判断算法是I.isTerminator() ||!I.isSafeToRemove()，他判断的条件是：如果指令是basicblock的终止符或者指令对后续调用有影响，则加入Alive中。接下来通过遍历Alive中的指令，并将指令的操作数转换为上一次操作的Instruction，这样就可以通过类似use-def chain的方式实现Alive的反向传播，从而将指令中变量use chain中的变量都加入Alive白名单中，表示该变量被使用过。最后再遍历所有指令就可以识别出改指令是不是无用的代码。

核心代码如下：
```cpp
namespace {
struct MYDCE : public FunctionPass {
  ...
  bool runOnFunction(Function &F) override;
  void getAnalysisUsage(AnalysisUsage &AU) const override {
    AU.setPreservesCFG();
}};}
  
bool MYADCE::runOnFunction(Function &F) {
  for (Instruction &I : instructions(F)) {
    if (I.isTerminator() ||!I.isSafeToRemove()) {
        Instruction *MyIn = &I;
        std::string mystr;
        raw_string_ostream stream(mystr);
        stream << *MyIn;
        mystr = stream.str();
        Alive.insert(&I);
        Worklist.push_back(&I);
}}

while (!Worklist.empty()) {
Instruction *Curr = Worklist.pop_back_val();
for (Use &OI : Curr->operands()) { // 遍历当前指令的操作数
  // dyn_cast将当前指令的操作数动态转化为Instruction
  if (Instruction *Inst = dyn_cast<Instruction>(OI)) {
      Worklist.push_back(Inst);
}}}
...
// 遍历函数的指令集，如果指令不在Alive集中，则放入Worklist，并删除指令的所有引用
for (Instruction &I : instructions(F)) {
if (!Alive.count(&I)) {
  Worklist.push_back(&I);
  I.dropAllReferences();
}}
for (Instruction *I : Worklist) {
  I->eraseFromParent(); // 清除
}
```

## 8. 测试
### LFYLang Test

#### while & array
<img src="https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/15/15921533208940.jpg" width="55%">

#### scan & print & char
<img src="https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/15/15921535227654.jpg" width="55%">

#### expr & double
<img src="https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/15/15921535979914.jpg" width="55%">

#### if & func with no args 
<img src="https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/15/15921536709103.jpg" width="55%">

#### func with args & recurse
<img src="https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/15/15921538951144.jpg" width="55%">

### Opt Pass Test
#### MyDCE
左中右分别为：MyDCE优化过的IR、原本的IR、source code。可以看出dce对死代码进行了消除。
![](https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/16/15923170754881.jpg)

#### MyInline
左右分别为：MyInline优化过的IR、原本的IR。可以看出MyInline对inline函数b进行了展开。
![](https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/16/15923175666107.jpg)

#### SSA
左右分别为：原本的IR、经过SSA优化过的IR。可以看出我的编译器已经能够生成满足SSA格式的IR，并且能够通过LLVM本身的SROA Pass进行phi指令的插入。
![](https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/06/16/15923177946650.jpg)

#### 尾调用
所谓为调用，是在函数调用的过程中，如果要再次调用一个函数并且返回，那么应该在这个函数的最后直接通过return调用。比如：
```c
int id(x) {
    return x;
}
int f(a) {
    const b = a + 1;
    // return id(b);     是尾调用
    // return b * id(b); 不是尾调用
}
int main() {
    f(1);
}
```
这样，在f函数要返回时，我们的函数main的栈帧中只需要存id函数的栈帧即可，这样做最大的意义是，当使用递归时，我们可以保证我们的栈帧是个可接受的常量，而不是一直递归导致爆栈。

在LLVM中，想要使用尾调用优化需要显示使用fastcc调用约定。同时还需要加上tail call关键字。使用后我们可以通过LLVM Pass对函数中的return值进行回溯然后进行优化。

## 9. 设计特色
1. 在c语言中，多态可以通过函数指针+结构体来实现，通过结构体包含，并将每个语句的具体函数绑定到对应结构体的函数指针来实现Node向下转型，从而实现语法树的生成。而在go语言中，并没有函数指针这一概念，因此需要使用go中的struct和interface来实现多态，并通过interface的断言和反射功能来实现向下转型。具体实现在5.1节语法树设计中有体现。
2. Github部分commit截图
![](https://pic-1257433408.cos.ap-chengdu.myqcloud.com/2020/07/13/15946416657645.jpg)
目前已经开源在github：
https://github.com/LFYSec/ToyCompiler

## 10. 学习过程中，遇到和解决的技术问题

1. 采用Go语言，如何实现语法树中各节点的多态
开始的设计是使用Go，但是因为并没有开发编译器的经验，因此我也不知道应该怎么实现语法树。在开发过程中发现需要用到类似函数指针的东西，通过查阅资料用Go的interface解决了问题。
2. LLVM资料很少，且大多数是英文
LLVM虽然已经很火了，但是中文资料非常少，并且很老，最靠谱的还是官方文档。初此之外，如果想开发LLVM Pass，那就得掌握一定的LLVM API，而LLVM的版本号更迭实在是太频繁了(不愧是版本号大师)，因此我只能通过阅读LLVM doxygen来学习。这里推荐一下LLVM CookBook这本书，虽然也已经很老了，但是还是十分适合入门学习。

## 11. 自我评价
通过参加编译原理试点班，我学到了很多理论课上学不到的东西，比如编译器的开发、整体设计、算法实现等等，这个过程中我既加深了对理论知识的理解，又锻炼了开发能力。
在开发过程中，我也有过几次熬夜到3点多向老师请教学习，非常敬佩老师的工作态度，也非常感谢老师的教导，让我开发和对编译原理的理解更深入更扎实。