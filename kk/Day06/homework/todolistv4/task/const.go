package task

const (
	JsonTaskFile = "db/tasks.json"
	TaskFile     = "db/tasks.txt"
	GobTaskFile  = "db/tasks.gob"
	CsvTaskFile  = "db/tasks.csv"
	RetainCsvNum = 3
)

const (
	statusNew     = "未执行"
	statusBegin   = "开始执行"
	statusPause   = "暂停"
	statusCompele = "完成"
)

var (
	TimeLayout = "2006/01/02 15:04:05"
	Header     = []string{"ID", "Name", "StartTime", "EndTime", "Status", "User"}
	StatusMap  = map[string]string{
		"1": statusNew,
		"2": statusBegin,
		"3": statusPause,
		"4": statusCompele,
	}
)
