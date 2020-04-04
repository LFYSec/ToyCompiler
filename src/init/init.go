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
	fmt.Print("@double_fmt_str = private unnamed_addr constant [4 x i8] c\"%g\\0A\\00\", align 1\n@int_fmt_str = private unnamed_addr constant [4 x i8] c\"%d\\0A\\00\", align 1\n")
	fmt.Print("define i32 @main() #0 {\n")
	symbolTable.PushFrame()
	br := bufio.NewReader(os.Stdin)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			// ^D: Input end of file on Unix/Linux
			// ^Z: Input end of file on Windows
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}

		if s := strings.TrimSpace(line); s == "q" || s == "quit" || s == "exit" {
			break
		}

		// flex + goyacc


		yyParse(newLexer([]byte(line)))
		result.GeneCode()
	}
	fmt.Print("ret i32 0\n")
	fmt.Print("}\ndeclare i32 @printf(i8*, ...) #1\n")
}