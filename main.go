package main

import (
	"log"
	"os"

	"github.com/james-lawrence/tinygo-wasi-bug/wasi/runtime/env"
)

func main() {
	log.Println("hello world", os.Getenv("FOO"), env.Boolean(false, "HELLO_WORLD", "HELLO_WORLD_2"))
}

// ERROR
// exported wasi/runtime/env.Boolean ( i32, i32, i32, i32 ) i32
// imported wasi/runtime/env.Boolean ( i32, i32, i32, i32, i32 ) i32
// import[4] func[env.wasi/runtime/env.Boolean]: signature mismatch: i32i32i32i32i32_i32 != i32i32i32i32_i32
