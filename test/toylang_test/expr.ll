@.str.double = private unnamed_addr constant [4 x i8] c"%f\0A\00", align 1
@.str.int = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
@.str = private unnamed_addr constant [4 x i8] c"%c\0A\00", align 1

@.scan.double = private unnamed_addr constant [4 x i8] c"%lf\00", align 1
@.scan.int = private unnamed_addr constant [3 x i8] c"%d\00", align 1
@.scan = private unnamed_addr constant [3 x i8] c"%c\00", align 1

define i32 @main() #0 {
%a_0 = alloca double
call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.scan.double, i32 0, i32 0), double* %a_0)
%b_0 = alloca double
%temp_2 = load double, double* %a_0
%temp_3 = fsub double %temp_2,                               1.00000000000000000000000000000000
store double %temp_3, double* %b_0
%temp_5 = load double, double* %a_0
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.double, i32 0, i32 0), double %temp_5)
%temp_6 = load double, double* %b_0
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.double, i32 0, i32 0), double %temp_6)
%temp_7 = load double, double* %a_0
%temp_8 = load double, double* %b_0
%temp_9 = fadd double %temp_7, %temp_8
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.double, i32 0, i32 0), double %temp_9)
%temp_10 = load double, double* %a_0
%temp_11 = load double, double* %b_0
%temp_12 = fsub double %temp_10, %temp_11
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.double, i32 0, i32 0), double %temp_12)
%temp_13 = load double, double* %a_0
%temp_14 = load double, double* %b_0
%temp_15 = fmul double %temp_13, %temp_14
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.double, i32 0, i32 0), double %temp_15)
%temp_16 = load double, double* %a_0
%temp_17 = load double, double* %b_0
%temp_18 = fdiv double %temp_16, %temp_17
call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.double, i32 0, i32 0), double %temp_18)
	ret i32 0
}

declare i32 @printf(i8*, ...) #1
declare i32 @scanf(i8*, ...) #1