package main

func main() {

}

func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		stones = sorts.(stones)

		if stones[0] == stones[1] {
			stones = stones[2:]
		} else {
			stones[0] = stones[0] - stones[1]
			stones = stones[1:]
		}
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
