package main

import (
	"net/http"
)

// const keyServerAddr = "serverAddr"

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	fmt.Printf("%s: got root request\n", ctx.Value(keyServerAddr))
// 	http.FileServer(http.Dir("src/html"))
// 	io.WriteString(w, "This is my website!\n")
// }

// func getHello(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

func main() {
	// main
	http.ListenAndServe(":8080", http.FileServer(http.Dir("src/html")))
}

// }
// 	fmt.Printf("starting server...\n")
// 	err := http.ListenAndServe(":3333", mux)

// 	if errors.Is(err, http.ErrServerClosed) {
// 		fmt.Printf("server closed\n")
// 	} else if err != nil {
// 		fmt.Printf("error starting server: %s/n;", err)
// 		//os.Exit(1)
// 	}
