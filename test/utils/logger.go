package utils

type MockLogger struct {
	info map[Severity][]string
}

type Severity uint

const (
	Debug Severity = iota
	Info
	Warning
	Error
)

func NewMockLogger() *MockLogger {
	return &MockLogger{
		info: map[Severity][]string{},
	}
}

func (m *MockLogger) helper(severity Severity, text string) {
	val, ok := m.info[severity]

	if ok && val != nil {
		m.info[severity] = append(val, text)
	} else {
		m.info[severity] = []string{text}
	}
}

func (m *MockLogger) Debug(text string) {
	m.helper(Debug, text)
}

func (m *MockLogger) Info(text string) {
	m.helper(Info, text)
}

func (m *MockLogger) Warning(text string) {
	m.helper(Warning, text)
}

func (m *MockLogger) Error(text string) {
	m.helper(Error, text)
}

func (m *MockLogger) GetLog(severity Severity) ([]string, bool) {
	data, ok := m.info[severity]
	return data, ok
}

func (m *MockLogger) GetCount(severity Severity) int {
	val, ok := m.GetLog(severity)

	if !ok {
		return 0
	} else {
		return len(val)
	}
}

func (m *MockLogger) CleanLog() {
	m.info = map[Severity][]string{}
}
