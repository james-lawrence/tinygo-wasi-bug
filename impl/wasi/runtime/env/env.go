package env

//export github.com/james-lawrence/tinygo-wasi-bug/wasi/runtime/env.Boolean
func Boolean(fallback bool, keys ...string) bool {
	return false
}

//export github.com/james-lawrence/tinygo-wasi-bug/wasi/runtime/env.Boolean2
func Boolean2(i bool) bool {
	return false
}

//export github.com/james-lawrence/tinygo-wasi-bug/wasi/runtime/env.Ellipsis
func Ellipsis(b ...bool) bool {
	return false
}
