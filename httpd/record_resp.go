package httpd

import (
	"net/http"
)

type writeOnceRecordResponseWriter struct {
	http.ResponseWriter

	statusCode int
	nw         int
}

func newWriteOnceRecordResponseWriter(w http.ResponseWriter) *writeOnceRecordResponseWriter {
	return &writeOnceRecordResponseWriter{
		ResponseWriter: w,
		nw:             -1,
		statusCode:     http.StatusOK,
	}
}

func (rp *writeOnceRecordResponseWriter) Write(data []byte) (int, error) {
	if rp.nw == -1 {
		rp.nw = 0
		rp.ResponseWriter.WriteHeader(rp.statusCode)
	}
	nw, err := rp.ResponseWriter.Write(data)
	rp.nw += nw
	return nw, err
}

func (rp *writeOnceRecordResponseWriter) WriteHeader(statusCode int) {
	if rp.nw != -1 {
		return
	}
	rp.statusCode = statusCode
}
