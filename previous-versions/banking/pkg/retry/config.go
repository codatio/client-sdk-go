// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package retry

type BackoffStrategy struct {
	InitialInterval int
	MaxInterval     int
	Exponent        float64
	MaxElapsedTime  int
}

type Config struct {
	Strategy              string
	Backoff               *BackoffStrategy
	RetryConnectionErrors bool
}
