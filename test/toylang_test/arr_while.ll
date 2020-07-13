@.str.double = private unnamed_addr constant [4 x i8] c"%g\0A\00", align 1
@.str.int = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
@.str = private unnamed_addr constant [4 x i8] c"%c\0A\00", align 1

@.scan.double = private unnamed_addr constant [3 x i8] c"%g\00", align 1
@.scan.int = private unnamed_addr constant [3 x i8] c"%d\00", align 1
@.scan = private unnamed_addr constant [3 x i8] c"%c\00", align 1

define i32 @main() #0 {
%array_0 = alloca [10 x i32]
%i_0 = alloca i32
store i32 0, i32* %i_0
%num_0 = alloca i32
store i32 10, i32* %num_0
br label %while_condition_1
while_condition_1:
%temp_3 = load i32, i32* %i_0
%temp_4 = icmp slt i32 %temp_3, 10
br i1 %temp_4, label %while_body_1,label %while_end_1
while_body_1:
%temp_5 = load i32, i32* %i_0
%array_element_7 = getelementptr inbounds [10 x i32], [10 x i32]* %array_0, i64 0, i32 %temp_5
%temp_8 = load i32, i32* %i_0
store i32 %temp_8, i32* %array_element_7
%temp_10 = load i32, i32* %i_0
%temp_11 = add nsw i32 %temp_10, 1
store i32 %temp_11, i32* %i_0
br label %while_condition_1
while_end_1:
store i32 0, i32* %i_0
br label %while_condition_2
while_condition_2:
%temp_13 = load i32, i32* %i_0
%temp_14 = icmp slt i32 %temp_13, 10
br i1 %temp_14, label %while_body_2,label %while_end_2
while_body_2:
%temp_15 = load i32, i32* %i_0
%array_element_17 = getelementptr inbounds [10 x i32], [10 x i32]* %array_0, i64 0, i32 %temp_15
%temp_16 = load i32, i32* %array_element_17
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.int, i32 0, i32 0), i32 %temp_16)
%temp_19 = load i32, i32* %i_0
%temp_20 = add nsw i32 %temp_19, 1
store i32 %temp_20, i32* %i_0
br label %while_condition_2
while_end_2:
	ret i32 0
}

declare i32 @printf(i8*, ...) #1
declare i32 @scanf(i8*, ...) #1