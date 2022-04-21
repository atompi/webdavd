package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"gitee.com/autom-studio/webdavd/internal/config"
	"gitee.com/autom-studio/webdavd/internal/utils"
	"golang.org/x/net/webdav"
)

func DirListHandler(fs webdav.FileSystem, w http.ResponseWriter, r *http.Request) bool {
	ctx := context.Background()
	f, err := fs.OpenFile(ctx, r.URL.Path, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	defer f.Close()
	if fileInfo, _ := f.Stat(); fileInfo != nil && !fileInfo.IsDir() {
		return false
	}
	dirs, err := f.Readdir(-1)
	if err != nil {
		log.Print(w, "Error reading directory", http.StatusInternalServerError)
		return false
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", name, name)
	}
	fmt.Fprintf(w, "</pre>\n")
	return true
}

func WebdavHandler(webdavdConfig config.WebdavdConfig) http.HandlerFunc {
	rootPath := webdavdConfig.RootPath
	dir := webdavdConfig.Dir
	fs := &webdav.Handler{
		FileSystem: webdav.Dir(rootPath + dir.Path),
		LockSystem: webdav.NewMemLS(),
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if dir.Auth.Enabled {
			username, password, ok := r.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if username != dir.Auth.Username || !utils.PasswordVerifier(dir.Auth.Password, password) {
				http.Error(w, "WebDAV: Unauthorized!", http.StatusUnauthorized)
				return
			}
		}
		if r.Method == "GET" && DirListHandler(fs.FileSystem, w, r) {
			return
		}
		if dir.Readonly {
			switch r.Method {
			case "POST", "DELETE", "PUT", "MKCOL", "COPY", "MOVE", "LOCK", "UNLOCK", "PROPPATCH":
				http.Error(w, "WebDAV: Method Not Allowed!", http.StatusMethodNotAllowed)
				return
			}
		}
		fs.ServeHTTP(w, r)
	}
}
