package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

var Maxsize int
var Backup_path string

func init() {
	Maxsize = 50000              // 日志文件最大容量
	Backup_path = "./log_backup" // 备份日志存放的文件夹
}

// 判断日志文件是否已满
func is_full(file *os.File) bool {
	file_info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if file_info.Size() >= int64(Maxsize) {
		return true
	}
	return false
}

// 将当前日志文件备份到备份文件夹中
func backup_file(file *os.File) {
	// 先检查备份文件夹是否存在，若不存在则创建备份文件夹
	_, err := os.Stat(Backup_path)
	if os.IsNotExist(err) {
		err = os.Mkdir(Backup_path, 0777)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	// 用备份时的时间作为备份日志的文件名
	now := time.Now().Format("20060102150405.000000")
	backup_name := Backup_path + "/" + now + ".log"
	new_file, err := os.OpenFile(backup_name, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	io.Copy(new_file, file)
}

// 创建四种logger,若指定文件已满则备份并清空该日志文件
func new_debugger(s string) *log.Logger {
	file, err2 := os.OpenFile(s, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err2 != nil {
		fmt.Println(err2)
		panic(err2)
	}
	if is_full(file) {
		backup_file(file)                                       // 备份日志
		file, err2 = os.OpenFile(s, os.O_RDWR|os.O_TRUNC, 0777) //清空文件
		if err2 != nil {
			fmt.Println(err2)
			panic(err2)
		}
	}
	return log.New(file, "[DEBUG]", log.LstdFlags)
}

func new_warner(s string) *log.Logger {
	file, err2 := os.OpenFile(s, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err2 != nil {
		fmt.Println(err2)
		panic(err2)
	}
	if is_full(file) {
		backup_file(file)                                       // 备份日志
		file, err2 = os.OpenFile(s, os.O_RDWR|os.O_TRUNC, 0777) //清空文件
		if err2 != nil {
			fmt.Println(err2)
			panic(err2)
		}
	}
	return log.New(file, "[WARNING]", log.LstdFlags)
}

func new_infoer(s string) *log.Logger {
	file, err2 := os.OpenFile(s, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err2 != nil {
		fmt.Println(err2)
		panic(err2)
	}
	if is_full(file) {
		backup_file(file)                                       // 备份日志
		file, err2 = os.OpenFile(s, os.O_RDWR|os.O_TRUNC, 0777) //清空文件
		if err2 != nil {
			fmt.Println(err2)
			panic(err2)
		}
	}
	return log.New(file, "[INFO]", log.LstdFlags)
}

func new_errorer(s string) *log.Logger {
	file, err2 := os.OpenFile(s, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err2 != nil {
		fmt.Println(err2)
		panic(err2)
	}
	if is_full(file) {
		backup_file(file)                                       // 备份日志
		file, err2 = os.OpenFile(s, os.O_RDWR|os.O_TRUNC, 0777) //清空文件
		if err2 != nil {
			fmt.Println(err2)
			panic(err2)
		}
	}
	return log.New(file, "[ERROR]", log.LstdFlags)
}

// 获得调用该package来记录日志的程序所在位置
func get_caller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed!")
	}
	return "file:" + file + "  line:" + strconv.Itoa(line)
}

// 记录四种日志，传入日志文件名和要记录的日志内容

// RecordDebug 向指定的日志文件中写DEBUG类型的日志
func RecordDebug(filename string, content string) {
	l := new_debugger(filename)
	caller_information := get_caller(2)
	l.Printf("[%v] %v", caller_information, content)
}

// RecordWarning 向指定的日志文件中写WARNING类型的日志
func RecordWarning(filename string, content string) {
	l := new_warner(filename)
	caller_information := get_caller(2)
	l.Printf("[%v] %v", caller_information, content)
}

// RecordInfo 向指定的日志文件中写INFO类型的日志
func RecordInfo(filename string, content string) {
	l := new_infoer(filename)
	caller_information := get_caller(2)
	l.Printf("[%v] %v", caller_information, content)
}

// RecordError 向指定的日志文件中写ERROR类型的日志
func RecordError(filename string, content string) {
	l := new_errorer(filename)
	caller_information := get_caller(2)
	l.Printf("[%v] %v", caller_information, content)
}
