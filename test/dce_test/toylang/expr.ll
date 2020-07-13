@.str.double = private unnamed_addr constant [4 x i8] c"%f\0A\00", align 1
@.str.int = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
@.str = private unnamed_addr constant [4 x i8] c"%c\0A\00", align 1

@.scan.double = private unnamed_addr constant [4 x i8] c"%lf\00", align 1
@.scan.int = private unnamed_addr constant [3 x i8] c"%d\00", align 1
@.scan = private unnamed_addr constant [3 x i8] c"%c\00", align 1

define i32 @test() #0 {
%a_0 = alloca i32
%b_0 = alloca i32
%c_0 = alloca i32
store i32 1, i32* %a_0
%temp_2 = load i32, i32* %a_0
%temp_3 = icmp eq i32 %temp_2, 1
br i1 %temp_3, label %if_lable_1, label %else_lable_1
if_lable_1:
%d_0 = alloca i32
br label %if_end_lable_1
else_lable_1:
br label %if_end_lable_1
if_end_lable_1:
	ret i32 0
}
define i32 @main() #0 {
%a_0 = alloca i32
%b_0 = alloca i32
%temp_4 = load i32, i32* %a_0
store i32 %temp_4, i32* %b_0
%fvar_test_0 = call i32 @test()
	ret i32 0
}

declare i32 @printf(i8*, ...) #1
declare i32 @scanf(i8*, ...) #1