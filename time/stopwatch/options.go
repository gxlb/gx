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
	// StepCacheCap is the initial cap of step duration list.
	StepCacheCap int

	// Unit is the min time unit of duration report.
	Unit Duration

	// Name if name the the stop watch.
	Name string
}

// WithStepCacheCap is the optional function that set the initial cap of step cache.
func WithStepCacheCap(stepCacheCap int) Option {
	return func(opts *Options) {
		opts.StepCacheCap = stepCacheCap
	}
}

// WithUnit is the optional function that set time unit of stopwatch.
func WithUnit(unit Duration) Option {
	return func(opts *Options) {
		opts.Unit = unit
	}
}

// WithName is the optional function that set name of stopwatch.
func WithName(name string) Option {
	return func(opts *Options) {
		opts.Name = name
	}
}
