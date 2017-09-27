package log

type Logger struct {
	hm map[string]int
}

func Constructor() Logger {
	return Logger{
		hm: make(map[string]int),
	}
}

func (l *Logger) ShouldPrintMessage(message string, timestamp int) bool {
	if val, ok := l.hm[message]; ok {
		if val+10 > timestamp {
			return false
		}
	}
	l.hm[message] = timestamp
	return true
}
