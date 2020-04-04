run:
	make flex
	make goyacc

	@go fmt
	go run .

flex:
	flex --prefix=yy --header-file=src/init/lex.yy.h -o src/init/lex.yy.c src/init/lex.l

goyacc:
	goyacc -o src/init/y.go src/init/grammer.y
