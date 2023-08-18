1. 知识整理
2. 结构体
   1. task 改成结构体
3. todolist 密码
   1. 检查密码文件.password
   2. 文件不存在的话，让用户初始化密码
   3. 如果文件存在，就让用户输入密码进行验证
   4. 命令(增删改查)  修改密码
4. 文件、文件夹
   1. cp -r dir dir
   2. 实现一个cp命令
5. todolist 存储
   1. id,name,startTime,endTime,status,user
   2. bufio

```go
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
```

 每次读操作 read文件，写操作 write文件(用第一种)

 程序启动 read文件，程序退出write文件