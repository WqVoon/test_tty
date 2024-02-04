package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	// 隐藏/展示光标
	hideCursor = "\x1b[?25l"
	showCursor = "\x1b[?25h"

	// 进入/退出备用缓冲区
	enterSubTerminal = "\x1b[?1049h"
	exitSubTerminal  = "\x1b[?1049l"

	// 清屏
	clearTerminal = "\x1b[1J\x1b[H\x1b[0J" // 1J 是清空光标前的所有内容，H 是让光标回到 (0, 0) 的位置，0J 是清空光标后的所有内容
)

// 输出一批终端转义序列
func printEscapeSeq(s ...string) { fmt.Printf(strings.Join(s, "")) }

// 运行在 MacOS 的 iterm2 终端模拟器，shell 使用 oh-my-zsh
func main() {
	selectOneOption()
}

// 查看某个按键的键码
func checkKeyNumber() {
	printEscapeSeq(hideCursor, enterSubTerminal)
	defer func() {
		printEscapeSeq(showCursor, exitSubTerminal)
	}()

	buf := make([]byte, 128)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		printEscapeSeq(clearTerminal)
		fmt.Println("按键的键码为:", buf[:n], "\n\n按 esc 或 enter 退出")

		// 按回车或 Esc 退出
		if n == 1 && buf[0] == 10 || n == 1 && buf[0] == 27 {
			return
		}
	}
}

// 按任意键继续
func pressAnyKey() {
	printEscapeSeq(hideCursor)
	defer func() {
		printEscapeSeq(showCursor)
	}()

	fmt.Println("按任意键继续...")
	buf := make([]byte, 128)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		panic(err)
	}
	if n > 0 {
		return
	}
}

// 选择一个选项
func selectOneOption() {
	options := []string{"一二三四五", "上山打老虎", "老虎没打着", "打到小松鼠"}
	choice := 0

	printEscapeSeq(hideCursor, enterSubTerminal)
	defer func() {
		printEscapeSeq(showCursor, exitSubTerminal)
		fmt.Println("你的选择是:", options[choice])
	}()

	buf := make([]byte, 128)
	for {
		// 输出选择
		printEscapeSeq(clearTerminal)
		fmt.Println("请选择:")
		for idx, opText := range options {
			if idx == choice {
				fmt.Printf("-> ")
			} else {
				fmt.Printf("   ")
			}
			fmt.Println(opText)
		}

		// 获取输入
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		// 按了回车
		if n == 1 && buf[0] == 10 {
			return
		}
		// 按了上键
		if n == 3 && buf[0] == 27 && buf[1] == 91 && buf[2] == 65 {
			choice = (choice + len(options) - 1) % len(options)
		}
		// 按了下键
		if n == 3 && buf[0] == 27 && buf[1] == 91 && buf[2] == 66 {
			choice = (choice + len(options) + 1) % len(options)
		}
	}
}
