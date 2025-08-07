package typed

type LimitOptions struct {
	ShrinkThresholdCap int     // Minimum cap to consider shrinking
	ShrinkUsageRatio   float64 // Usage ratio (e.g., 0.25 = shrink if using <25%)
	EnableAutoShrink   bool    // Control whether auto-shrink is enabled
}

func DefaultLimitOptions() LimitOptions {
	return LimitOptions{
		ShrinkThresholdCap: 1024,
		ShrinkUsageRatio:   0.25,
		EnableAutoShrink:   true,
	}
}

type LimitOption func(*LimitOptions)

func WithShrinkThresholdCap(cap int) LimitOption {
	if cap <= 0 {
		panic("Shrink threshold capacity must be greater than 0")
	}
	return func(o *LimitOptions) {
		o.ShrinkThresholdCap = cap
	}
}

func WithShrinkUsageRatio(ratio float64) LimitOption {
	if ratio < 0 || ratio > 1 {
		panic("Shrink usage ratio must be between 0 and 1")
	}
	return func(o *LimitOptions) {
		o.ShrinkUsageRatio = ratio
	}
}

func WithAutoShrink(enabled bool) LimitOption {
	return func(o *LimitOptions) {
		o.EnableAutoShrink = enabled
	}
}
