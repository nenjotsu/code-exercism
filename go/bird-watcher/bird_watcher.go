package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	t := 0
	for i := 0; i < len(birdsPerDay); i++ {
		t += birdsPerDay[i]
	}
	return t
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	limit := 7
	counter := 0
	for i := 0; i < len(birdsPerDay); i += limit {
		perWeek := birdsPerDay[i:min(i+limit, len(birdsPerDay))]
		counter++
		if counter == week {
			sum := 0
			for j := 0; j < len(perWeek); j++ {
				sum += perWeek[j]
			}
			return sum
		}
	}
	return 0
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	bit := 1
	for i := 0; i < len(birdsPerDay); i++ {
		if bit == 1 {
			birdsPerDay[i]++
			bit = 0
		} else {
			bit = 1
		}
	}
	return birdsPerDay
}
