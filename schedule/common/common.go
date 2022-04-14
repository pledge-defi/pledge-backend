package common

import (
	"os"
	"strings"
)

var PlgrAdminPrivateKeyTestNet string
var PlgrAdminPrivateKeyMainNet string

func Getenv() {

	PlgrAdminPrivateKeyTestNet = os.Getenv("plgr_admin_private_key_test_net")
	PlgrAdminPrivateKeyTestNet = strings.Replace(PlgrAdminPrivateKeyTestNet, " ", "", -1)

	PlgrAdminPrivateKeyMainNet = os.Getenv("plgr_admin_private_key_main_net")
	PlgrAdminPrivateKeyMainNet = strings.Replace(PlgrAdminPrivateKeyMainNet, " ", "", -1)

	if PlgrAdminPrivateKeyTestNet == "" || PlgrAdminPrivateKeyMainNet == "" {
		panic("environment variable is not set")
	}

}
