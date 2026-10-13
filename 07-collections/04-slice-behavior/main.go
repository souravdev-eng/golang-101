package main

import "fmt"

// A function receives its own slice header, but refers to the caller's elements.
func addScore(s []int) []int {
	s[0] = 99
	s = append(s, 30)
	return s
}

func main() {
	// Each block starts fresh so earlier writes cannot affect later predictions.
	{
		scores := []int{10, 20, 30, 40}
		middle := scores[1:3]
		middle[0] = 99
		fmt.Println("shared:", scores, "middle:", middle, "len/cap:", len(middle), cap(middle))
	}
	{
		scores := make([]int, 2, 4) // Enough capacity: append keeps sharing.
		scores[0], scores[1] = 10, 20
		alias := scores
		scores = append(scores, 30)
		scores[0] = 99
		fmt.Println("reuse:", scores, "alias:", alias)
	}
	{
		scores := []int{10, 20} // Full: append must allocate.
		alias := scores
		scores = append(scores, 30)
		scores[0] = 99
		fmt.Println("growth:", scores, "alias:", alias)
	}
	{
		scores := []int{10, 20, 30, 40}
		firstTwo := scores[:2] // Capacity extends past this view's length.
		firstTwo = append(firstTwo, 99)
		fmt.Println("overwrite:", scores)
	}
	{
		scores := []int{10, 20, 30, 40}
		firstTwo := scores[:2:2]        // Limits capacity; does not itself copy.
		firstTwo = append(firstTwo, 99) // Now forces new storage.
		firstTwo[0] = 77
		fmt.Println("limited:", scores, "result:", firstTwo)
	}
	{
		scores := make([]int, 2, 3)
		scores[0], scores[1] = 10, 20
		updated := addScore(scores)
		fmt.Println("function:", scores, "returned:", updated)
		scores = updated // Keep the returned header when the length changes.
	}
	{
		scores := []int{10, 20, 30}
		snapshot := make([]int, len(scores))
		n := copy(snapshot, scores)
		snapshot[0] = 99
		fmt.Println("copy:", n, "original:", scores, "snapshot:", snapshot)
		empty := make([]int, 0, len(scores))
		fmt.Println("empty copy:", copy(empty, scores), "len/cap:", len(empty), cap(empty))
	}
	{
		var scores []int
		fmt.Println("nil:", scores == nil, "len/cap:", len(scores), cap(scores), "appended:", append(scores, 10))
	}
	{
		scores := []int{10, 20}
		for _, score := range scores {
			score += 5 // Only changes this local int.
		}
		for i := range scores {
			scores[i] += 5 // Changes the stored elements.
		}
		fmt.Println("range:", scores)
	}
}
