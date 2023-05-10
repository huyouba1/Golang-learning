package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	//	BASEDIR          string
	PROC_PORT        int
	h                bool
	help             bool
	RASP_PACKAGE_URL string = "https://172.32.10.16/YSUpdateV6/zhiyu-rasp.zip"
	RASP_BASEDIR     string = "/usr/local/rasp"
	JAVA_HOME        string
	JAVA_TOOLS       string = JAVA_HOME + "/lib/tools.jar"
	//	RASP_PACKAGE_URL string
	//	RASP_BASEDIR     string
)

func main() {
	flag.IntVar(&PROC_PORT, "p", 0, "Java service port requiring attach")
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("Usage: jrasp_install [-p] <JAVA_SERVICE_PORT>")
		flag.PrintDefaults()
	}

	if h || help {
		flag.Usage()
		os.Exit(1)
	}

	checkStatement := fmt.Sprintf(`ss -utpln |grep -w %d |awk '{print $NF}'  |column -t | awk -F ',' '{print $(NF-1)}' |egrep -o [0-9]+`, PROC_PORT)
	ProcPID, err := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	JavaOp := fmt.Sprintf(`readlink -f /proc/%s/exe`, strings.TrimSpace(string(ProcPID)))
	JavaCmdPath, err := exec.Command("sh", "-c", JavaOp).CombinedOutput()

	//fmt.Println(JavaOp)
	if err != nil {
		fmt.Println(err)
	}

	// 去掉换行符
	path := strings.TrimSpace(string(JavaCmdPath))
	JAVA_HOME = filepath.Dir(filepath.Dir(path))

	WriteFile("/tmp/"+string(ProcPID)+".pid", string(ProcPID))
	//fmt.Println(GetJavaHome(PROC_PORT))
	DownloadRaspPackage(RASP_PACKAGE_URL, RASP_BASEDIR)
	StartJar(PROC_PORT)
}

func StartJar(port int) {
	checkStatement := fmt.Sprintf(`ss -utpln |grep -w %d |awk '{print $NF}'  |column -t | awk -F ',' '{print $(NF-1)}' |egrep -o [0-9]+`, PROC_PORT)
	ProcPID, err := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	UserOp := fmt.Sprintf(`ps -p %s -o user |tail -1`, strings.TrimSpace(string(ProcPID)))
	StartUser, err := exec.Command("sh", "-c", UserOp).CombinedOutput()

	//fmt.Println(StartUser)
	if err != nil {
		fmt.Println(err)
	}
	WriteFile("/tmp/"+string(ProcPID)+".pid", string(ProcPID))

	if string(StartUser) == "root" {
		op := fmt.Sprintf(`%s/bin/java  -jar -Xbootclasspath/a:%s  %s/zhiyu-rasp/rasp-attach.jar -p %s`, JAVA_HOME, JAVA_TOOLS, RASP_BASEDIR, strings.TrimSpace(string(ProcPID)))
		exec.Command(op)
	} else {
		op := fmt.Sprintf(`su - %s -c "%s/bin/java  -jar -Xbootclasspath/a:%s  %s/zhiyu-rasp/rasp-attach.jar -p %s"`, string(StartUser), JAVA_HOME, JAVA_TOOLS, RASP_BASEDIR, strings.TrimSpace(string(ProcPID)))
		exec.Command(op)
	}
}

func http_download_curl(url, path string) {
	downloadCmd := exec.Command("curl", "-sLk", "-o", path, url)
	err := downloadCmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 判断文件是否存在
func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

// 判断目录是否存在
func DirIsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	} else {
		panic(err)
	}
}

// IsDir函数，检查给定路径是否为目录
func IsDir(path string) bool {
	// 获取给定路径的文件信息
	fileinfo, err := os.Stat(path)
	// 如果发生错误，则抛出异常
	if err != nil {
		println(err)
	}
	// 返回文件信息是否为目录
	return fileinfo.IsDir()
}

func DownloadRaspPackage(url, path string) {
	os.MkdirAll(path, os.ModePerm)
	if FileIsExists(path+"/zhiyu-rasp.zip") == false {
		http_download_curl(url, path)
	}
	if DirIsExists(path+"/zhiyu-rasp") == false {
		exec.Command("unzip", path+"/zhiyu-rasp.zip", "-d", path)
	}

}

func WriteFile(path, txt string) {
	// 打开给定路径的文件
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	// 如果发生错误，则抛出异常
	if err != nil {
		println(err)
	}
	// 延迟关闭文件
	defer file.Close()
	// 将给定的文本写入文件
	file.WriteString(txt)
}
