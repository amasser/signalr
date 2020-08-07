package signalr

import (
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"time"
)

// TimeoutInterval is the interval one party will consider the other party disconnected
// if it hasn't received a message (including keep-alive) in it.
// The recommended value is double the KeepAliveInterval value.
// Default is 30 seconds.
func TimeoutInterval(timeout time.Duration) func(party) error {
	return func(p party) error {
		p.setTimeout(timeout)
		return nil
	}
}

// HandshakeTimeout is the interval if the other party doesn't send an initial handshake message within,
// the connection is closed. This is an advanced setting that should only be modified
// if handshake timeout errors are occurring due to severe network latency.
// For more detail on the handshake process,
// see https://github.com/dotnet/aspnetcore/blob/master/src/SignalR/docs/specs/HubProtocol.md
func HandshakeTimeout(timeout time.Duration) func(party) error {
	return func(p party) error {
		p.setHandshakeTimeout(timeout)
		return nil
	}
}

// KeepAliveInterval is the interval if the party hasn't sent a message within,
// a ping message is sent automatically to keep the connection open.
// When changing KeepAliveInterval, change the Timeout setting on the other party.
// The recommended Timeout value is double the KeepAliveInterval value.
// Default is 15 seconds.
func KeepAliveInterval(interval time.Duration) func(party) error {
	return func(p party) error {
		p.setKeepAliveInterval(interval)
		return nil
	}
}

// StreamBufferCapacity is the maximum number of items that can be buffered for client upload streams.
// If this limit is reached, the processing of invocations is blocked until the the server processes stream items.
// Default is 10.
func StreamBufferCapacity(capacity uint) func(party) error {
	return func(p party) error {
		if capacity == 0 {
			return errors.New("unsupported StreamBufferCapacity 0")
		}
		p.setStreamBufferCapacity(capacity)
		return nil
	}
}

// MaximumReceiveMessageSize is the maximum size of a single incoming hub message.
// Default is 32KB
func MaximumReceiveMessageSize(size uint) func(party) error {
	return func(p party) error {
		if size == 0 {
			return errors.New("unsupported maximumReceiveMessageSize 0")
		}
		p.setMaximumReceiveMessageSize(size)
		return nil
	}
}

// ChanReceiveTimeout is the timeout for processing stream items from the client, after StreamBufferCapacity was reached
// If the hub method is not able to process a stream item during the timeout duration,
// the server will send a completion with error.
// Default is 5 seconds.
func ChanReceiveTimeout(timeout time.Duration) func(party) error {
	return func(p party) error {
		p.setChanReceiveTimeout(timeout)
		return nil
	}
}

// EnableDetailedErrors - if true, detailed exception messages are returned to the other party when an exception is thrown in a Hub method.
// The default is false, as these exception messages can contain sensitive information.
func EnableDetailedErrors(enable bool) func(party) error {
	return func(p party) error {
		p.setEnableDetailedErrors(enable)
		return nil
	}
}

// StructuredLogger is the simplest logging interface for structured logging.
// See github.com/go-kit/kit/log
type StructuredLogger interface {
	Log(keyVals ...interface{}) error
}

// Logger stets the logger used by the party to log info events.
// If debug is true, debug log event are generated, too
func Logger(logger StructuredLogger, debug bool) func(party) error {
	return func(p party) error {
		i, d := buildInfoDebugLogger(logger, debug)
		p.setLoggers(i, d)
		return nil
	}
}

func buildInfoDebugLogger(logger log.Logger, debug bool) (log.Logger, log.Logger) {
	if debug {
		logger = level.NewFilter(logger, level.AllowDebug())
	} else {
		logger = level.NewFilter(logger, level.AllowInfo())
	}
	return level.Info(logger), log.With(level.Debug(logger), "caller", log.DefaultCaller)
}