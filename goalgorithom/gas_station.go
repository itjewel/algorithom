package main
import "fmt"


func canCompleteCircuit(gas []int, cost []int) int {
	totalCost :=0
	totalGas :=0
	tank :=0
	start :=0
	for i:=0; i<len(gas); i++{
		tank += gas[i] - cost[i]
		totalCost += totalCost
		totalGas += totalGas
		
		if tank < 0 {
			start = i+1
			tank = 0

		}
	}
	if totalCost > totalGas{
		return -1
	}
	return start
}

func main() {
	gas := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	result := canCompleteCircuit(gas,cost)
	fmt.Println(result)
}
