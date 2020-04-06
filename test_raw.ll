@double_fmt_str = private unnamed_addr constant [4 x i8] c"%g\0A\00", align 1
@int_fmt_str = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
define i32 @main() #0 {
%a_65 = alloca i32
store i32 1, i32* %a_65
%temp_2 = load i32, i32* %a_65
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @int_fmt_str, i32 0, i32 0), i32 %temp_2)
br label %while_condition_1
while_condition_1:
%temp_3 = load i32, i32* %a_65
%temp_4 = icmp eq i32 %temp_3, 1
br i1 %temp_4, label %while_body_1,label %while_end_1
while_body_1:
%temp_5 = load i32, i32* %a_65
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @int_fmt_str, i32 0, i32 0), i32 %temp_5)
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @int_fmt_str, i32 0, i32 0), i32 2)
br label %while_condition_1
while_end_1:
ret i32 0
}
declare i32 @printf(i8*, ...) #1
