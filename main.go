package main

import (
	cron1 "github.com/jakecoffman/cron"
	"os/exec"
	//cron2 "github.com/robfig/cron"
	"fmt"
)
//全局变量标识
var check_task_mark = 1
var check_week_mark = 1
var check_day_mark = 1
var php_exe = "E:\\Visual-NMP-x64\\Bin\\PHP\\php-5.6.40-nts-x86\\php"
var code_path = "E:\\szm\\www\\agent\\trunk\\"

func check_week() {
	if check_week_mark == 1 {
		check_week_mark = 2
		fmt.Println("this is check_week Test")

		cmd := exec.Command(php_exe, code_path+"check_week.php")
		buf, err := cmd.Output()
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println(string(buf))

		check_week_mark = 1
	}
}


func check_day() {
	if check_day_mark == 1 {
		check_day_mark = 2
		fmt.Println("this is check_day Test")

		cmd := exec.Command(php_exe, code_path+"check_day.php")
		buf, err := cmd.Output()
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println(string(buf))

		check_day_mark = 1
	}
}
func check_task() {
	if check_task_mark == 1 {
		check_task_mark = 2
		fmt.Println("this is conFun Test")

		cmd := exec.Command(php_exe, code_path+"check_task.php")
		buf, err := cmd.Output()
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println(string(buf))

		check_task_mark = 1
	}
}
func main() {
	//fetchweb()
	c := cron1.New()
	c.AddFunc("*/5 * * * * ?", check_task, "check_task")
	c.AddFunc("*/5 * * * * ?", check_day, "check_day")
	c.AddFunc("*/5 * * * * ?", check_week, "check_week")
	c.Start()

	select{}
}