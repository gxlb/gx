package stopwatch

// Option is the optional function.
type Option func(opts *Options)

func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

// Options contains all options which will be applied when instantiating a stopwatch.
type Options struct {
	StepCacheCap int
}

func WithStepCacheCap(stepCacheCap int) Option {
	return func(opts *Options) {
		opts.StepCacheCap = stepCacheCap
	}
}
