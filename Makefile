run:
	make flex
	make goyacc

	@go fmt
	go run .

flex:
	flex --prefix=yy --header-file=lex.yy.h -o src/tools/C/lex.yy.c lex.l

goyacc:
	goyacc grammer.y
