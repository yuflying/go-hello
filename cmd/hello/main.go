package main

func main() {

	// 初始化全局配置
	initConfig()

	// 启动服务
	s := NewServer()
	s.start()
	defer s.stop()

}
