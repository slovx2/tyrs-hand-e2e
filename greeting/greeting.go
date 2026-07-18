package greeting

import "strings"

// Hello 返回给指定用户的问候语。
func Hello(name string) string {
	return "Hello, " + name + "!"
}

// FriendlyHello 返回清理姓名后的友好问候语。
func FriendlyHello(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "friend"
	}

	return Hello(name)
}
