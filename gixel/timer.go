package gixel

type TimerCallback = func(totalElapsed float64, iteration int)

type GxlTimer struct {
	iterations    int
	currIteration int
	elapsed       float64
	totalElapsed  float64
	interval      float64
	done          bool
	running       bool
	callback      TimerCallback
	onStart       TimerCallback
}

func NewTimer(time float64) *GxlTimer {
	return NewIterationTimer(1, time)
}

func NewLoopTimer(time float64) *GxlTimer {
	return NewIterationTimer(0, time)
}

func NewIterationTimer(iterations int, interval float64) *GxlTimer {
	return &GxlTimer{
		iterations:    iterations,
		currIteration: 0,
		elapsed:       0,
		totalElapsed:  0,
		interval:      interval,
		done:          false,
		running:       false,
		callback:      nil,
		onStart:       nil,
	}
}

func (t *GxlTimer) Iteration() int {
	return t.currIteration
}

func (t *GxlTimer) TotalElapsed() float64 {
	return t.totalElapsed
}

func (t *GxlTimer) Elapsed() float64 {
	return t.elapsed
}

func (t *GxlTimer) Done() bool {
	return t.done
}

func (t *GxlTimer) Running() bool {
	return t.running
}

func (t *GxlTimer) Progress() float64 {
	if t.iterations > 1 {
		return t.totalElapsed / (float64(t.iterations) * t.interval)
	}

	return t.elapsed / t.interval
}

func (t *GxlTimer) Start() {
	if t.onStart != nil {
		t.onStart(0, 0)
	}
	t.running = true
}

func (t *GxlTimer) Stop() {
	t.running = false
}

func (t *GxlTimer) Resume() {
	t.running = true
}

func (t *GxlTimer) Reset() {
	t.running = false
	t.done = false
	t.currIteration = 0
	t.elapsed = 0
}

func (t *GxlTimer) Restart() {
	t.Reset()
	t.Start()
}

func (t *GxlTimer) SetOnStart(callback TimerCallback) *GxlTimer {
	t.onStart = callback
	return t
}

func (t *GxlTimer) SetCallback(callback TimerCallback) *GxlTimer {
	t.callback = callback
	return t
}

func (t *GxlTimer) Update(elapsed float64) {
	if !t.running || t.done {
		return
	}

	t.elapsed += elapsed
	t.totalElapsed += elapsed

	if t.elapsed >= t.interval {
		t.elapsed = 0
		t.currIteration++

		if t.iterations > 0 {
			if t.currIteration >= t.iterations {
				t.callback(t.totalElapsed, t.currIteration)
				t.currIteration = 0
				t.done = true
				t.running = false
				return
			}
		}

		t.callback(t.totalElapsed, t.currIteration)
	}
}
