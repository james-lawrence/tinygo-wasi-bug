package env

//export github.com/james-lawrence/tinygo-wasi-bug/wasi/runtime/env.Boolean
func Boolean(fallback bool, keys ...string) bool {
	return false
}
