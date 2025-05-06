package main

import (
	"strings"
)

// 为了避免函数未使用的警告，添加一个简单的 main 函数来调用 mergeAlternately 函数
func main() {
	mergeAlternately("hello", "world")
}

func mergeAlternately(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	var builder strings.Builder
	builder.Grow(n + m) // 预分配内存

	for i := 0; i < n || i < m; i++ {
		if i < n {
			builder.WriteByte(word1[i]) // 直接操作字节
		}
		if i < m {
			builder.WriteByte(word2[i])
		}
	}
	return builder.String()
}
