func climbStairs(n int) int {
   one, two := 1, 1

   // one = is the current value, but initially it is step 1
   // two = the n-1 steps or initialially 0. because technically it takes 1 steps to reach zero
   //lets say for n = 2. i need sum of ways to reach 2-1 =1 and 2-2 = 0
   // so, one = 2-1 and two = 2-2, so one and two steps from n. 
   // but after loop, one will be the current stairs.
   // for n = 2, already have 1 and 0. so need to run one time/
   for range n - 1{
      one, two = one + two, one 
   }   
   return one
}
