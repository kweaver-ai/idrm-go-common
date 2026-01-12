package audit

// Discard returns a Logger that discards all messages logged to it. It can be
// used whenever the caller is not interested in the logs. Logger instance
// produced by this function always compare as equal.
func Discard() Logger { return Logger{sink: nil} }
