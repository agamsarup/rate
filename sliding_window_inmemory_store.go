package rate

type SlidingWindowInMemoryStore struct {
	prevWindowRequestsCount map[string]int
	curWindowRequestsCount  map[string]int
}

func NewSlidingWindowInMemoryStore() *SlidingWindowInMemoryStore {
	return &SlidingWindowInMemoryStore{
		prevWindowRequestsCount: map[string]int{},
		curWindowRequestsCount:  map[string]int{},
	}
}

func (s *SlidingWindowInMemoryStore) CurrentWindowCount(id string) int {
	return s.curWindowRequestsCount[id]
}

func (s *SlidingWindowInMemoryStore) PreviousWindowCount(id string) int {
	return s.prevWindowRequestsCount[id]
}

func (s *SlidingWindowInMemoryStore) MoveToNextWindow() {
	s.prevWindowRequestsCount = s.curWindowRequestsCount
	s.curWindowRequestsCount = map[string]int{}
}

func (s *SlidingWindowInMemoryStore) Reset() {
	s.prevWindowRequestsCount = map[string]int{}
	s.curWindowRequestsCount = map[string]int{}
}

func (s *SlidingWindowInMemoryStore) Increment(id string, count int) {
	s.curWindowRequestsCount[id] += count
}
