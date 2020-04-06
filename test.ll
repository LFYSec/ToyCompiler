@double_fmt_str = private unnamed_addr constant [4 x i8] c"%g\0A\00", align 1
@int_fmt_str = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
define void @a(i32 %b_65,i32 %c_65) #0 {
%tmp_b_65 = alloca i32, align 4
store i32 %b_65, i32* %tmp_b_65, align 4
%tmp_c_65 = alloca i32, align 4
store i32 %c_65, i32* %tmp_c_65, align 4
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @int_fmt_str, i32 0, i32 0), i32 %b_65)
ret void
}
define i32 @main() #0 {
call void @a(i32 1,i32 2)
ret i32 0
}
declare i32 @printf(i8*, ...) #1
