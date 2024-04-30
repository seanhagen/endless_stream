// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package grpc

import (
	"fmt"
	"strings"
)

const (
	// ServerOptionKeyNetwork is a serverOptionKey of type Network.
	ServerOptionKeyNetwork serverOptionKey = iota
	// ServerOptionKeyTLS is a serverOptionKey of type TLS.
	ServerOptionKeyTLS
	// ServerOptionKeyStatsHandler is a serverOptionKey of type StatsHandler.
	ServerOptionKeyStatsHandler
)

var ErrInvalidserverOptionKey = fmt.Errorf("not a valid serverOptionKey, try [%s]", strings.Join(_serverOptionKeyNames, ", "))

const _serverOptionKeyName = "NetworkTLSStatsHandler"

var _serverOptionKeyNames = []string{
	_serverOptionKeyName[0:7],
	_serverOptionKeyName[7:10],
	_serverOptionKeyName[10:22],
}

// serverOptionKeyNames returns a list of possible string values of serverOptionKey.
func serverOptionKeyNames() []string {
	tmp := make([]string, len(_serverOptionKeyNames))
	copy(tmp, _serverOptionKeyNames)
	return tmp
}

// serverOptionKeyValues returns a list of the values for serverOptionKey
func serverOptionKeyValues() []serverOptionKey {
	return []serverOptionKey{
		ServerOptionKeyNetwork,
		ServerOptionKeyTLS,
		ServerOptionKeyStatsHandler,
	}
}

var _serverOptionKeyMap = map[serverOptionKey]string{
	ServerOptionKeyNetwork:      _serverOptionKeyName[0:7],
	ServerOptionKeyTLS:          _serverOptionKeyName[7:10],
	ServerOptionKeyStatsHandler: _serverOptionKeyName[10:22],
}

// String implements the Stringer interface.
func (x serverOptionKey) String() string {
	if str, ok := _serverOptionKeyMap[x]; ok {
		return str
	}
	return fmt.Sprintf("serverOptionKey(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x serverOptionKey) IsValid() bool {
	_, ok := _serverOptionKeyMap[x]
	return ok
}

var _serverOptionKeyValue = map[string]serverOptionKey{
	_serverOptionKeyName[0:7]:                    ServerOptionKeyNetwork,
	strings.ToLower(_serverOptionKeyName[0:7]):   ServerOptionKeyNetwork,
	_serverOptionKeyName[7:10]:                   ServerOptionKeyTLS,
	strings.ToLower(_serverOptionKeyName[7:10]):  ServerOptionKeyTLS,
	_serverOptionKeyName[10:22]:                  ServerOptionKeyStatsHandler,
	strings.ToLower(_serverOptionKeyName[10:22]): ServerOptionKeyStatsHandler,
}

// ParseserverOptionKey attempts to convert a string to a serverOptionKey.
func ParseserverOptionKey(name string) (serverOptionKey, error) {
	if x, ok := _serverOptionKeyValue[name]; ok {
		return x, nil
	}
	return serverOptionKey(0), fmt.Errorf("%s is %w", name, ErrInvalidserverOptionKey)
}

func (x serverOptionKey) Ptr() *serverOptionKey {
	return &x
}

// MarshalText implements the text marshaller method.
func (x serverOptionKey) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *serverOptionKey) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseserverOptionKey(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}