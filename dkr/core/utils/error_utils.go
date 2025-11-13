package utils

// 错误中断
func Must(err error) {
	if err != nil {
		panic(err)
	}
}