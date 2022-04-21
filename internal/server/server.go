package server

import (
	"net/http"

	"gitee.com/autom-studio/webdavd/internal/config"
	"gitee.com/autom-studio/webdavd/internal/handler"
)

func StartServer(webdavdConfig config.WebdavdConfig) {
	http.Handle("/", handler.WebdavHandler(webdavdConfig))

	if webdavdConfig.Server.TLS.Enabled {
		http.ListenAndServeTLS(webdavdConfig.Server.Addr, webdavdConfig.Server.TLS.Cert, webdavdConfig.Server.TLS.Key, nil)
	} else {
		http.ListenAndServe(webdavdConfig.Server.Addr, nil)
	}
}
