package web

import (
	"net/http"

	"github.com/bingoohuang/pkger"
)

func Serve() {
	pkger.Stat("github.com/bingoohuang/pkger:/README.md")
	dir := http.FileServer(pkger.Dir("/public"))
	http.ListenAndServe(":3000", dir)
}
