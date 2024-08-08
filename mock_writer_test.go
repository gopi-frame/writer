package writer

import "bytes"

type mockWriter struct {
	buffer *bytes.Buffer
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return m.buffer.Write(p)
}

func (m *mockWriter) Close() error {
	return nil
}
