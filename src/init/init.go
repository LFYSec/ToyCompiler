package init

import (
	"Compiler/src/symbolTable"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Parse() {
	str := "@.str.double = private unnamed_addr constant [4 x i8] c\"%g\\0A\\00\", align 1\n" +
		"@.str.int = private unnamed_addr constant [4 x i8] c\"%d\\0A\\00\", align 1\n"
	fmt.Println(str)
	//fmt.Print("define i32 @main() #0 {\n")
	symbolTable.PushFrame()
	file, err := os.Open("test4.l")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	//br := bufio.NewReader(os.Stdin)
	var s string
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			// ^D: Input end of file on Unix/Linux
			// ^Z: Input end of file on Windows
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if s := strings.TrimSpace(line); s == "q" || s == "quit" || s == "exit" {
			break
		}
		s += strings.Replace(line, "\n", "", -1)
	}
	// flex + goyacc
	yyParse(newLexer([]byte(s)))
	result.GeneCode()
	//fmt.Print("ret i32 0\n")
	fmt.Print("\ndeclare i32 @printf(i8*, ...) #1\ndeclare i32 @scanf(i8*, ...) #1")
}