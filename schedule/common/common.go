package common

import (
	"os"
	"os/exec"
	"pledge-backend/log"
)

var PlgrAdminPrivateKeyTestNet string
var PlgrAdminPrivateKeyMainNet string

func GetEnv() {

	var ok bool

	exec.Command("source /etc/profile")

	PlgrAdminPrivateKeyTestNet, ok = os.LookupEnv("PLGR_PRIVATE_KEY")
	if !ok {
		log.Logger.Error("environment variable is not set")
		panic("environment variable is not set")
	}

	PlgrAdminPrivateKeyMainNet, ok = os.LookupEnv("plgr_admin_private_key_main_net")
	if !ok {
		log.Logger.Error("environment variable is not set")
		panic("environment variable is not set")
	}

}
