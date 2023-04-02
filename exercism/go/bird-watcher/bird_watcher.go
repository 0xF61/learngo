package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int

	for _, v := range birdsPerDay {
		sum += v
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var choicedWeekFirstDay int
	var choicedWeek []int

	choicedWeekFirstDay = 7 * (week - 1)
	choicedWeek = birdsPerDay[choicedWeekFirstDay : choicedWeekFirstDay+7]

	return TotalBirdCount(choicedWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
