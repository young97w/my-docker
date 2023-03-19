package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	//cmd.SysProcAttr是一个结构体指针，用于设置新进程的系统属性
	//syscall.SysProcAttr是一个包含系统进程属性的结构体，其中Cloneflags字段用于设置进程的命名空间。
	//syscall.CLONE_NEWUTS标志表示创建一个新的UTS命名空间，该命名空间包含了主机名和域名信息，可以实现在不同进程之间隔离主机名和域名信息。
	//syscall.CLONE_NEWIPC标志表示创建一个新的IPC命名空间，该命名空间包含了进程间通信的各种IPC机制，如消息队列、信号量等，
	//可以实现不同进程之间的IPC隔离。
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
