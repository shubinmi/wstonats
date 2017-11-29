package wstonats

import (
	"log"
	"net/http"
)

var id = 0

var proxySetting = new(ProxySetting)

func writeLog(msg []byte, level int) {
	if proxySetting.LogWriter == nil || proxySetting.DebugLevel == DebugNon {
		return
	}
	if level == DebugErr {
		proxySetting.LogWriter.Write(append([]byte("[ERR] "), msg...))
		return
	}
	if proxySetting.DebugLevel == DebugInfo {
		proxySetting.LogWriter.Write(append([]byte("[INFO] "), msg...))
		return
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	c := newConnector(w, r)
	go c.init()
}

//Start noinspection GoUnusedExportedFunction
func Start(s *ProxySetting) {
	proxySetting = s
	if s.WsTls {
		log.Fatal(http.ListenAndServeTLS(proxySetting.WsAddr, s.WsTlsCert, s.WsTlsKey, http.HandlerFunc(proxyHandler)))
	} else {
		log.Fatal(http.ListenAndServe(proxySetting.WsAddr, http.HandlerFunc(proxyHandler)))
	}
}
