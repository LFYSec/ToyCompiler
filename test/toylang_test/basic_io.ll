@.str.double = private unnamed_addr constant [4 x i8] c"%g\0A\00", align 1
@.str.int = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
@.str = private unnamed_addr constant [4 x i8] c"%c\0A\00", align 1

@.scan.double = private unnamed_addr constant [3 x i8] c"%g\00", align 1
@.scan.int = private unnamed_addr constant [3 x i8] c"%d\00", align 1
@.scan = private unnamed_addr constant [3 x i8] c"%c\00", align 1

define i32 @main() #0 {
%a_0 = alloca i8
store i8 97, i8* %a_0
%b_0 = alloca i8
call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.scan, i32 0, i32 0), i8* %b_0)
%c_0 = alloca i32
store i32 2, i32* %c_0
%temp_5 = load i32, i32* %c_0
%temp_6 = add nsw i32 %temp_5, 1
store i32 %temp_6, i32* %c_0
%temp_7 = load i8, i8* %a_0
%temp_7_c = sext i8 %temp_7 to i32
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i32 0, i32 0), i32 %temp_7_c)
%temp_8 = load i8, i8* %b_0
%temp_8_c = sext i8 %temp_8 to i32
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i32 0, i32 0), i32 %temp_8_c)
%temp_9 = load i32, i32* %c_0
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.int, i32 0, i32 0), i32 %temp_9)
	ret i32 0
}

declare i32 @printf(i8*, ...) #1
declare i32 @scanf(i8*, ...) #1