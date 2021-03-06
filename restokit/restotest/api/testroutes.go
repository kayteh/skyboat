// Package api is for API routes. Everything here is ideally unexported.
package api // import "skyboat.io/x/restokit/restotest/api"

//go:generate go run $PWD/restokit/codegen/codegen.go $CWD

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// GET /test v2 default
func testGet(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello world! v2")
}

// GET /test v1
func testGetv1(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello world! v1")
}

// GET /hello/:name
// NoLogging
func hello(ctx *fasthttp.RequestCtx) {
	ctx.WriteString(fmt.Sprintf("Hello, %s!", ctx.UserValue("name")))
}

// // POST /test v1
// // Inject
// func testPost(ctx *fasthttp.RequestCtx) {
// 	ctx.WriteString("Hello world!")
// }

// // POST /test v2
// // Inject Inject2(Ole) Other
// func testPost(ctx *fasthttp.RequestCtx) {
// 	ctx.WriteString("Hello world!")
// }
