package rate

type IDRateLimiter interface {
	Allow(id string) bool
}
