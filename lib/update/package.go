package update

import (
	"sync"
	"fmt"
	"time"
)

var Flag uint16
var m *sync.RWMutex = new(sync.RWMutex)
//return false,the caller have to unpack the SSU,and inc Flag
func GetFlag()bool{
	m.RLock()
	defer m.RUnlock()
	if Flag == 0 {
		return false
	}else{
		return true
	}
}
//when unpack SSU done, it should call this function
func IncFlag(){
	m.Lock()
	defer m.Unlock()
	Flag++
}

//when upgrade success, it should call this function
func DecFlag(){
	m.Lock()
	defer m.Unlock()
	if Flag > 0{
		Flag--
	}
}

//相同的版本的SSU只能解压一次,在没有解压完成之前其它goroute只能等待解压完成，需要channel来通信
var once sync.Once
func (S *Session)unpackSSU(ssu string){

}

func UnpackSSU(){
	if !GetFlag(){
		IncFlag()
		//don't have to unpack SSU,because it has been unpacked
		return
	}
	//var name string
	var S Session
	once.Do(S.unpackSSU)


	IncFlag()
}


func PrepareUpgrade(S *Session,ssu string)error{
	fmt.Println("init to upgrade or restore  the package:%s",ssu)
	if S.UpdatingFlag && ((time.Now() - S.UpdateTime) < UPD_TIMEOUT){
		fmt.Errorf("now update the package:%s,begin at %t ....")
	}

}