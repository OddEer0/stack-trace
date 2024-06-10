package stacktrace

func ParceFunc(trace any) any {
	switch trace.(type) {
	case string:
		return Message{Message: trace.(string)}
	}
	return trace
}
