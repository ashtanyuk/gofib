package main

func FibRec(n int) int {
   if n <=1 {
      return 1
   } 
   return FibRec(n - 1) + FibRec(n - 2)
}
