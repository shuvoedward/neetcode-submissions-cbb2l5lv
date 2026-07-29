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

/*

Your broader question — is understanding the truth more important than exploring possibilities?

Yes, for building working solutions efficiently. But the possibilities-first thinking isn't wasted — it's usually how you discover the truth in the first place. You often need to mentally simulate a few brute-force cases to notice the pattern ("oh, if I fail here, everything before the fail point is also doomed") before you can state the invariant cleanly and trust it.

So a decent workflow is:

Brute-force a few examples by hand to build intuition
Look for what's always true across cases (invariant/property)
Prove to yourself why that property must hold (like you just did)
Only then write code that exploits the invariant directly, skipping the exhaustive search

You basically just did steps 2–3 out loud. That's a solid habit to keep leaning on.

*/
