package fib

func FibRec(n int) int {
   if n <=1 {
      return 1
   } else {
      return FibRec(n - 1) + FibRec(n - 2)
   }
}
