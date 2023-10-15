package job

import "errors"

var (
	ErrCanNotUpdateDelayedJobProgress = errors.New("can't change status, job is marked as 'Delayed'")
	ErrProgressUnknown                = errors.New("job progress status is unknown")
)
