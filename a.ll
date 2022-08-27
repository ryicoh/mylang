@s = unnamed_addr constant [3 x i8] c"%f\0A"

define i8 @main() {
0:
	%1 = fadd double 0x3FF3333333333333, 0x3FF199999999999A
	%hello = alloca double
	store double %1, double* %hello
	%2 = load double, double* %hello
	%3 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @s, i64 0, i64 0), double %2)
	ret i8 0
}

declare i32 @printf(i8* %0, ...)
