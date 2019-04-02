package login

import (
	"sync"

	"github.com/gfandada/gserver/util"
)

var (
	_id       int32
	_mutex    sync.Mutex
	_accountp map[string]string
)

func init() {
	_id = 1000
	_accountp = make(map[string]string)
}

func NewTourist() (int32, string, string) {
	_mutex.Lock()
	defer _mutex.Unlock()
	_id++
	account := util.NewV4().String()
	passwd := util.NewV4().String()
	_accountp[account] = passwd
	return _id, account, passwd
}

func CheckTourist(account, passwd string) bool {
	_mutex.Lock()
	tmp, ok := _accountp[account]
	_mutex.Unlock()
	if !ok {
		return false
	}
	if passwd == tmp {
		return true
	}
	return false
}
