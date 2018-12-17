package mtproto

import (
	"fmt"
	"reflect"
)

func (m *MTProto) Help_GetConfig() (*TL_config, error) {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_help_getConfig{},
		resp,
	}
	x := <-resp
	switch f := x.(type) {
	case TL_config:
		config, _ := x.(TL_config)
		return &config, nil
	case TL_rpc_error:
		// log.Println(reflect.TypeOf(f).String(), f.error_code, f.error_message)
		if f.error_code == 303 {
			// Migrate Code
		}
		return nil, fmt.Errorf("TL_help_getConfig error [%d] %s", f.error_code, f.error_message)
	default:
		return nil, fmt.Errorf("TL_help_getConfig unknown answer %s", reflect.TypeOf(f).String())
	}
}
