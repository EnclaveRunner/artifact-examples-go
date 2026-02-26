//go:generate go tool wit-bindgen-go generate --world examples --out internal ./enclave:examples.wasm

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	helloworld "github.com/EnclaveRunner/examples-go/internal/enclave/examples-go/hello-world"
	_ "github.com/ydnar/wasi-http-go/wasihttp" // enable wasi-http
)

func init() {
	helloworld.Exports.Hello = func() [2]string {
		fmt.Println("Hello, World!")
		time.Sleep(2 * time.Second)
		fmt.Print("Slept for two seconds.")
		return [2]string{"", ""}
	}

	helloworld.Exports.Google = func() [2]string {
		fmt.Println("Fetching google.com!")
		resp, err := http.Get("https://www.google.com/")
		if err != nil {
			return [2]string{"", fmt.Sprintf("Error fetching google.com: %v\n", err)}
		}
		defer resp.Body.Close()
		fmt.Printf("Fetched google.com with status code: %d\n", resp.StatusCode)
		return [2]string{"", ""}
	}

	helloworld.Exports.Greet = func(name string) (result [2]string) {
		fmt.Printf("Welcome to Enclave, %s\n", name)
		return [2]string{"", ""}
	}

	helloworld.Exports.PrintArgs = func() (result [2]string) {
		fmt.Println("Provided arguments:")
		for _, arg := range os.Args {
			fmt.Println(arg)
		}

		return [2]string{"", ""}
	}

	helloworld.Exports.PrintEnv = func() (result [2]string) {
		fmt.Println("Environment variables:")
		for _, env := range os.Environ() {
			fmt.Println(env)
		}

		return [2]string{"", ""}
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
