@.str.double = private unnamed_addr constant [4 x i8] c"%g\0A\00", align 1
@.str.int = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
@.str = private unnamed_addr constant [4 x i8] c"%c\0A\00", align 1

@.scan.double = private unnamed_addr constant [3 x i8] c"%g\00", align 1
@.scan.int = private unnamed_addr constant [3 x i8] c"%d\00", align 1
@.scan = private unnamed_addr constant [3 x i8] c"%c\00", align 1

define i32 @b() #0 {
%a_0 = alloca i32
store i32 2, i32* %a_0
%temp_2 = load i32, i32* %a_0
%temp_3 = icmp ne i32 %temp_2, 1
br i1 %temp_3, label %if_lable_1, label %else_lable_1
if_lable_1:
	ret i32 2
br label %if_end_lable_1
else_lable_1:
	ret i32 1
br label %if_end_lable_1
if_end_lable_1:
	ret i32 0
}
define i32 @main() #0 {
%a_0 = alloca i32
%fvar_b_0 = call i32 @b()
store i32 %fvar_b_0, i32* %a_0
%temp_6 = load i32, i32* %a_0
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.int, i32 0, i32 0), i32 %temp_6)
	ret i32 0
}

declare i32 @printf(i8*, ...) #1
declare i32 @scanf(i8*, ...) #1