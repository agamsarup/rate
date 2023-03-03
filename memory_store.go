package rate

type SlidingWindowMemoryStore interface {
	Increment(id string, count int)
	CurrentWindowCount(id string) int
	PreviousWindowCount(id string) int
	MoveToNextWindow()
	Reset()
}
