package main

import (
	"github.com/nicle-lin/ADupgrade/lib/update"


	"os"

	//"github.com/astaxie/beego/logs"
	"fmt"
)

func main() {

	if err := update.Upgrade(os.Args[1],"51111","admin",os.Args[2]);err != nil{
		fmt.Println("err:",err)
	}else {
		fmt.Println("success")
	}

	/*
	err := update.PutFile(os.Args[1],"51111","admin",os.Args[2],os.Args[3])
	if err != nil {
		logs.Error(err)
	}
	*/
}
