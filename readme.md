compile
GOOS=js GOARCH=wasm go build -o gogame.wasm github.com/aofiee/gogame
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

// fs := http.FileServer(http.Dir("./static"))
// http.Handle("/static/", http.StripPrefix("/static/", fs))
// log.Fatal(http.ListenAndServe(":8080", nil))
