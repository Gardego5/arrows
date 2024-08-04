package game

import (
	"net/http"

	"github.com/Gardego5/arrows/lib"
	. "github.com/Gardego5/htmdsl"
	"github.com/labstack/echo/v4"
)

type hApp struct{ *lib.State }

func (h hApp) handle(c echo.Context) error {
	const script = `
if (!WebAssembly.instantiateStreaming) {
  WebAssembly.instantiateStreaming = (resp, importObject) => resp
    .then((data) => data.arrayBuffer())
    .then((source) => WebAssembly.instantiate(source, importObject));
}

const go = new Go();
WebAssembly
  .instantiateStreaming(fetch("/static/client.wasm"), go.importObject)
  .then((result) => go.run(result.instance));
`

	return c.Stream(http.StatusOK, "text/html", Fragment{DOCTYPE, Html{
		Head{
			Script{Attrs{{"src", "/static/wasm_exec.js"}}},
			Script{PreEscaped(script)},
		},
		Body{lib.CN("overflow-hidden")},
	}}.Reader())
}
