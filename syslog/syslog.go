/*
Package syslog defines functionality for integration with syslog.
*/
package syslog // import "go.pedge.io/protolog/syslog"
import (
	"log/syslog"

	"go.pedge.io/protolog"
)

var (
	// DefaultTextMarshaller is the default text Marshaller for syslog.
	DefaultTextMarshaller = protolog.NewTextMarshaller(
		protolog.MarshallerOptions{
			DisableTimestamp: true,
			DisableLevel:     true,
		},
	)
)

// PusherOptions defines options for constructing a new syslog protolog.Pusher.
type PusherOptions struct {
	Marshaller protolog.Marshaller
}

// NewPusher creates a new protolog.Pusher that logs using syslog.
func NewPusher(writer *syslog.Writer, options PusherOptions) protolog.Pusher {
	return newPusher(writer, options)
}

// NewStandardPusher creates a new protolog.Pusher that logs using syslog and the default text Marshaller.
func NewStandardPusher(writer *syslog.Writer) protolog.Pusher {
	return newPusher(
		writer,
		PusherOptions{
			Marshaller: DefaultTextMarshaller,
		},
	)
}
