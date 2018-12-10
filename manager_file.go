package mtproto

import (
	"reflect"
	"fmt"
)

func (m *MTProto) Upload_GetFile(in TL, offset int32, limit int32) ([]byte, error) {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_upload_getFile{
			Location: in,
			Offset:   offset,
			Limit:    limit,
		},
		resp,
	}
	x := <-resp
	switch f := x.(type) {
	case TL_upload_file:
		return f.Bytes, nil
	case TL_upload_fileCdnRedirect:
	case TL_rpc_error:
		// log.Println(reflect.TypeOf(f).String(), f.error_code, f.error_message)
		if f.error_code == 303 {
			// Migrate Code
		}
		return nil, fmt.Errorf("TL_upload_getFile error [%d] %s", f.error_code, f.error_message)
	default:
		return nil, fmt.Errorf("TL_upload_getFile unknown answer %s", reflect.TypeOf(f).String())
	}
	return nil, fmt.Errorf("unknown error")
}

func (m *MTProto) Upload_GetCdnFile(fileToken []byte, offset, limit int32) []byte {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_upload_getCdnFile{
			fileToken,
			offset,
			limit,
		},
		resp,
	}
	x := <-resp
	switch f := x.(type) {
	case TL_upload_cdnFileReuploadNeeded:
		m.queueSend <- packetToSend{
			TL_upload_reuploadCdnFile{
				Request_token: f.Request_token,
				File_token:    fileToken,
			},
			resp,
		}
		z := <-resp
		switch reflect.TypeOf(z).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(z)
			for i := 0; i < s.Len(); i++ {
				//hash := s.Interface().(TL_cdnFileHash)
				//TODO:: what to do now ?!!
			}
		}
	case TL_upload_cdnFile:
		return f.Bytes
	}
	return []byte{}
}
