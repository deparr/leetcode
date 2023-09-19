func arraySign(nums []int) int {
   necount := 0

   for _, val := range nums {
       if val < 0 {
           necount++
       } else if val == 0 {
           return 0
       }
   }

       if necount % 2 == 0 {
           return 1
       }

       return -1
}
