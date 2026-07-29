func canCompleteCircuit(gas []int, cost []int) int {
   totalGas, totalCost := 0, 0 
   for i := range gas{
	totalGas += gas[i]
	totalCost += cost[i]
   } 
   if totalGas < totalCost{
	return -1
   }

   total, res := 0, 0

   for i := range gas{
	total += gas[i] - cost[i]
	if total < 0{
		total = 0
		res = i + 1
	}
   }

   return res
}
