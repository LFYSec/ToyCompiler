@.str.double = private unnamed_addr constant [4 x i8] c"%f\0A\00", align 1
@.str.int = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
@.str = private unnamed_addr constant [4 x i8] c"%c\0A\00", align 1

@.scan.double = private unnamed_addr constant [4 x i8] c"%lf\00", align 1
@.scan.int = private unnamed_addr constant [3 x i8] c"%d\00", align 1
@.scan = private unnamed_addr constant [3 x i8] c"%c\00", align 1

define i32 @factorial(i32 %arg_val_0) #0 {
%val_0 = alloca i32, align 4
store i32 %arg_val_0, i32* %val_0, align 4
%temp_1 = load i32, i32* %val_0
%temp_2 = icmp eq i32 %temp_1, 1
br i1 %temp_2, label %if_lable_1, label %else_lable_1
if_lable_1:
%temp_3 = load i32, i32* %val_0
	ret i32 %temp_3
br label %if_end_lable_1
else_lable_1:
br label %if_end_lable_1
if_end_lable_1:
%x_0 = alloca i32
%temp_5 = load i32, i32* %val_0
%temp_6 = sub nsw i32 %temp_5, 1
%fvar_factorial_0 = call i32 @factorial(i32 %temp_6)
store i32 %fvar_factorial_0, i32* %x_0
%temp_8 = load i32, i32* %x_0
%temp_9 = load i32, i32* %val_0
%temp_10 = mul nsw i32 %temp_8, %temp_9
	ret i32 %temp_10
}
define i32 @main() #0 {
%a_0 = alloca i32
store i32 2, i32* %a_0
%temp_13 = load i32, i32* %a_0
%temp_14 = sub nsw i32 %temp_13, 1
store i32 %temp_14, i32* %a_0
%fvar_factorial_0 = call i32 @factorial(i32 3)
store i32 %fvar_factorial_0, i32* %a_0
%temp_17 = load i32, i32* %a_0
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.int, i32 0, i32 0), i32 %temp_17)
	ret i32 0
}

declare i32 @printf(i8*, ...) #1
declare i32 @scanf(i8*, ...) #1