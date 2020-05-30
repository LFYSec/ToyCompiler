package init

import (
	"Compiler/src/global"
	"Compiler/src/symbolTable"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Parse() {
	printStr := "@.str.double = private unnamed_addr constant [4 x i8] c\"%g\\0A\\00\", align 1\n" +
		"@.str.int = private unnamed_addr constant [4 x i8] c\"%d\\0A\\00\", align 1\n" +
		"@.str = private unnamed_addr constant [4 x i8] c\"%c\\0A\\00\", align 1\n"
	// TODO generate format string auto
	scanStr := "@.scan.double = private unnamed_addr constant [3 x i8] c\"%g\\00\", align 1\n" +
		"@.scan.int = private unnamed_addr constant [3 x i8] c\"%d\\00\", align 1\n" +
		"@.scan = private unnamed_addr constant [3 x i8] c\"%c\\00\", align 1\n"

	fmt.Println(printStr)
	fmt.Println(scanStr)

	symbolTable.PushFrame()
	var br *bufio.Reader
	if global.DEBUG {
		file, err := os.Open("recurse.l")
		//file, err := os.Open("test/toylang_test/recurse.l")
		if err != nil {
			fmt.Println("read file err:", err)
			return
		}
		defer file.Close()
		br = bufio.NewReader(file)
	} else {
		br = bufio.NewReader(os.Stdin)
	}


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
	fmt.Print("\ndeclare i32 @printf(i8*, ...) #1\ndeclare i32 @scanf(i8*, ...) #1")
}