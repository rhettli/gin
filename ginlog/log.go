package ginlog

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Init()
}

type FileName struct {
	Name string
}

func SetLogFileName(name string) *FileName {
	n := fmt.Sprintf(name+"-%s.log", time.Now().Format("20060102")) //2006-01-02 15:04:05
	fn := FileName{n}
	return &fn
}

//日志打印方法
func (n *FileName) Init() {
	flag.Parse() //解析参数付给logF

	//f:=os.Create("")

	outfile, err := os.OpenFile(n.Name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	defer func() {
		fmt.Println("io.close()")
		//outfile.Close()

	}()

	//return  nil
	//log.Printf(log_explain+"=%v \n ", log_content) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了
}
