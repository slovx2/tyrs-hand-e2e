package greeting

import "strings"

// Hello 返回给指定用户的问候语。
func Hello(name string) string {
	return "Hello, " + name + "!"
}

// Welcome 返回给指定用户的欢迎语。
func Welcome(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Welcome!"
	}

	return "Welcome, " + name + "!"
}
