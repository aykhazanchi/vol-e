package lib

import (
	"time"
)

type Timer struct {
	start time.Time
	stop  time.Time
}

func (t *Timer) Start() {
	t.start = time.Now()
}

func (t *Timer) Stop() {
	t.stop = time.Now()
}

func (t *Timer) SetTimeLearned() string {
	duration := t.stop.Sub(t.start)
	return string(int(duration.Seconds()/3600)) + "h " + string(int(duration.Seconds()/60)) + "m spent learning"
}
