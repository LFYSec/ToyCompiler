@.str = private unnamed_addr constant [3 x i8] c"%c\00", align 1
@.str.1 = private unnamed_addr constant [4 x i8] c"%c\0A\00", align 1

; Function Attrs: noinline nounwind optnone ssp uwtable
define i32 @main() #0 {
  %1 = alloca i8, align 1
  %2 = alloca i8, align 1
  store i8 97, i8* %1, align 1
  %3 = load i8, i8* %2, align 1
  %4 = sext i8 %3 to i32
  %5 = call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.str, i32 0, i32 0), i32 %4)
  %6 = load i8, i8* %1, align 1
  %7 = sext i8 %6 to i32
  %8 = load i8, i8* %2, align 1
  %9 = sext i8 %8 to i32
  %10 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.1, i32 0, i32 0), i32 %7, i32 %9)
  ret i32 0
}

declare i32 @scanf(i8*, ...) #1

declare i32 @printf(i8*, ...) #1
/*
    int main() {
        char a = 'a';
        char d;
        scanf("%c", d);
        printf("%c, %c\n", a,d);
    }
*/