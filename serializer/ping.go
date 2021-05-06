package serializer

// Ping 测试序列化器
type Ping struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

