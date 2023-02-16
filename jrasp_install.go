package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	//	BASEDIR          string
	PROC_PORT int
	h         bool
	help      bool
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
	fmt.Println(GetPID(PROC_PORT))
}

func GetPID(port int) string {

	checkStatement := fmt.Sprintf(`ss -utpln |grep -w %d |awk '{print $NF}'  |column -t | awk -F ',' '{print $(NF-1)}' |egrep -o [0-9]+`, port)
	ProcPID, err := exec.Command("sh", "-c", checkStatement).CombinedOutput()

	if err != nil {
		panic(err)
	}

	JavaOp := fmt.Sprintf(`"readlink", "-f", "/proc/%s/exe"`, ProcPID)
	JavaCmdPath, err := exec.Command("sh", "-c", JavaOp).CombinedOutput()
	//exec.Command("readlink", "-f", "/proc/"+string(ProcPID)+"/exe")
	if err != nil {
		//fmt.Println(err)
		panic(err)
		//os.Exit(1)
	}
	// 去掉换行符
	path := strings.TrimSpace(string(JavaCmdPath))
	fmt.Println(path)
	return path

}

func GetJavaHome(path string) string {

	return ""
}
