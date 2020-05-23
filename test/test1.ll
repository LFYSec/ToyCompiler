@.str.double = private unnamed_addr constant [4 x i8] c"%g\0A\00", align 1
@.str.int = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1

define i32 @factorial(i32 %arg_val_65) #0 {
%val_65 = alloca i32, align 4
store i32 %arg_val_65, i32* %val_65, align 4
br i1 %temp_2, label %if_lable_1, label %else_lable_1
if_lable_1:
%temp_3 = load i32, i32* %val_65
	ret i32 %temp_3
br label %if_end_lable_1
else_lable_1:
br label %if_end_lable_1
if_end_lable_1:
%fvar_factorial_65 = call i32 @factorial()
store i32 %fvar_factorial_65, i32* %x_65
%temp_7 = load i32, i32* %x_65
	ret i32 %temp_7
}
define i32 @main() #0 {
%a_65 = alloca i32
%fvar_factorial_65 = call i32 @factorial()
store i32 %fvar_factorial_65, i32* %a_65
%temp_10 = load i32, i32* %a_65
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.int, i32 0, i32 0), i32 %temp_10)
	ret i32 0
}

declare i32 @printf(i8*, ...) #1
declare i32 @scanf(i8*, ...) #1