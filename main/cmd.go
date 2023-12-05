package main

/*
	在GO语言中，API是以包(package)的形式提供的，包可以提供常量、变量、结构体以及函数等
		flag包：提供了命令行参数解析的功能
		fmt包：
		os包：定义了一个Args变量，存放了传递给命令行的全部参数
*/
import (
	"flag"
	"fmt"
	"os"
)

// 定义Cmd结构体
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	xJreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	// Usage: 用于定义把命令的用法打印到控制台的函数，如果flag.Parse()函数解析失败该函数就会被调用
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.xJreOption, "Xjre", "", "path to jre")
	flag.Parse()

	// 默认第一个参数为主类名，后续所有参数为传递给主类的参数
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage:%s [-options] class [args...]\n", os.Args[0])
}
