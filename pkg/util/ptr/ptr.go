package ptr

import "time"

// Bool to the provided bool.
func Bool(v bool) *bool { return &v }

// Time to the provided time.Time.
func Time(v time.Time) *time.Time { return &v }

func Float64(v float64) *float64 { return &v }

func Int8(v int8) *int8 { return &v }

// Str to the provided string.
func Str(v string) *string { return &v }
func Any[T any](v T) *T    { return &v }
