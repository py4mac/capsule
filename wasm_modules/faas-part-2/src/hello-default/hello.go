package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {
	html := `
    <html>
        <head>
            <title>Wasm is fantastic 😍</title>
        </head>

        <body>
            <h1>👋 Hello World 🌍</h1>
            <h2>Served with 💚💜 with Capsule 💊</h2>
            <h1>🚀</h1>
        </body>

    </html>
    `

	headersResp = map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return html, headersResp, nil
}
