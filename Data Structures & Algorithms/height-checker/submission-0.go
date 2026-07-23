func heightChecker(heights []int) int {
   count := make([]int, 101)
   for _, h := range heights{
	count[h]++
   } 

   expected := []int{}
   for h := 1; h <= 100; h++{
	c := count[h]
	for range c{
		expected = append(expected, h)
	}
   }

   res := 0
   for i, h := range heights{
	if h != expected[i]{
		res++
	}
   }

   return res
}