package util

import "testing"

/*
 测试单个模块
 使用 go test 执行测试代码
 go test -v 显示执行的明细
 go test -v -n 、 go test -v -5   执行测试代码五次
*/
func TestGenShortId(t *testing.T) {
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		t.Error("genshortId failed!")
	}
	t.Log("genshortId test pass")
}

// 性能测试
/*
说明:
性能测试函数名必须以 Benchmark 开头，如 BenchmarkXxx 或 Benchmark_xxx
go test 默认不会执行压力测试函数，需要通过指定参数 -test.bench 来运行压力测试函数，
-test.bench 后跟正则表达式，如 go test -test.bench=".*" 表示执行所有的压力测试函数
在压力测试中，需要在循环体中指定 testing.B.N 来循环执行压力测试代码
在 util 目录下执行命令 go test -test.bench=".*"


## 查看性能并生成函数调用图
执行命令：
$ go test -bench=".*" -cpuprofile=cpu.profile ./util
上述命令会在当前目录下生成 cpu.profile 和 util.test 文件。
执行 go tool pprof util.test cpu.profile 查看性能（进入交互界面后执行 top 指令）： top

## pprof 程序中最重要的命令就是 topN，此命令用于显示 profile 文件中的最靠前的 N 个样本（sample），它的输出格式各字段的含义依次是：
采样点落在该函数中的总时间
采样点落在该函数中的百分比
上一项的累积百分比
采样点落在该函数，以及被它调用的函数中的总时间
采样点落在该函数，以及被它调用的函数中的总次数百分比
函数名

*/
func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N ; i ++  {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B)  {
	b.StopTimer() // 调用该函数停止压力测试的时间技术
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		b.Error()
	}
	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N ; i++  {
		GenShortId()
	}
}