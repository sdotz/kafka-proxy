// Code generated by mocker; DO NOT EDIT
// github.com/travisjeffery/mocker
package kafkaproxy_test

import (
	"github.com/travisjeffery/jocko/protocol"
	"sync"
)

var (
	lockMockConnFetch sync.RWMutex
)

// MockConn is a mock implementation of Conn.
//
//     func TestSomethingThatUsesConn(t *testing.T) {
//
//         // make and configure a mocked Conn
//         mockedConn := &MockConn{
//             FetchFunc: func(in1 *protocol.FetchRequest) (*protocol.FetchResponses, error) {
// 	               panic("TODO: mock out the Fetch method")
//             },
//         }
//
//         // TODO: use mockedConn in code that requires Conn
//         //       and then make assertions.
//
//     }
type MockConn struct {
	// FetchFunc mocks the Fetch method.
	FetchFunc func(in1 *protocol.FetchRequest) (*protocol.FetchResponses, error)

	// calls tracks calls to the methods.
	calls struct {
		// Fetch holds details about calls to the Fetch method.
		Fetch []struct {
			// In1 is the in1 argument value.
			In1 *protocol.FetchRequest
		}
	}
}

// Reset resets the calls made to the mocked APIs.
func (mock *MockConn) Reset() {
	lockMockConnFetch.Lock()
	mock.calls.Fetch = nil
	lockMockConnFetch.Unlock()
}

// Fetch calls FetchFunc.
func (mock *MockConn) Fetch(in1 *protocol.FetchRequest) (*protocol.FetchResponses, error) {
	if mock.FetchFunc == nil {
		panic("moq: MockConn.FetchFunc is nil but Conn.Fetch was just called")
	}
	callInfo := struct {
		In1 *protocol.FetchRequest
	}{
		In1: in1,
	}
	lockMockConnFetch.Lock()
	mock.calls.Fetch = append(mock.calls.Fetch, callInfo)
	lockMockConnFetch.Unlock()
	return mock.FetchFunc(in1)
}

// FetchCalled returns true if at least one call was made to Fetch.
func (mock *MockConn) FetchCalled() bool {
	lockMockConnFetch.RLock()
	defer lockMockConnFetch.RUnlock()
	return len(mock.calls.Fetch) > 0
}

// FetchCalls gets all the calls that were made to Fetch.
// Check the length with:
//     len(mockedConn.FetchCalls())
func (mock *MockConn) FetchCalls() []struct {
	In1 *protocol.FetchRequest
} {
	var calls []struct {
		In1 *protocol.FetchRequest
	}
	lockMockConnFetch.RLock()
	calls = mock.calls.Fetch
	lockMockConnFetch.RUnlock()
	return calls
}
