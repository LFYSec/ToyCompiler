; ModuleID = '1.c'
source_filename = "1.c"
target datalayout = "e-m:o-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-apple-macosx10.14.0"

; Function Attrs: norecurse nounwind readnone ssp uwtable
define i32 @foo(i32, i32, i32) local_unnamed_addr #0 {
  br label %4

; <label>:4:                                      ; preds = %4, %3
  %5 = phi i32 [ %1, %3 ], [ %6, %4 ]
  %6 = add nsw i32 %5, %0
  %7 = icmp eq i32 %5, 0
  br i1 %7, label %8, label %4

; <label>:8:                                      ; preds = %4
  %9 = mul nsw i32 %6, %5
  %10 = sub nsw i32 %9, %5
  ret i32 %10
}

; Function Attrs: norecurse nounwind readnone ssp uwtable
define i32 @main() local_unnamed_addr #0 {
  ret i32 0
}

attributes #0 = { norecurse nounwind readnone ssp uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.module.flags = !{!0, !1, !2}
!llvm.ident = !{!3}

!0 = !{i32 2, !"SDK Version", [2 x i32] [i32 10, i32 14]}
!1 = !{i32 1, !"wchar_size", i32 4}
!2 = !{i32 7, !"PIC Level", i32 2}
!3 = !{!"Apple LLVM version 10.0.1 (clang-1001.0.46.4)"}
