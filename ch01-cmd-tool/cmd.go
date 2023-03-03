package main

/*
	在GO语言中，API是以包(package)的形式提供的，包可以提供常量、变量、结构体以及函数等
		flag包：提供了命令行参数解析的功能
		fmt包：
		os包：
*/
import "flag"
import "fmt"
import "os"

// 定义Cmd结构体
type Cmd struct {
	helpFlag 		bool
	versionFlag 	bool
	cpOption 		string
	class			string
	args			[]string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

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