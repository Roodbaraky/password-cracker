package internal

import (
	"crypto/md5"
	"fmt"
	"password_cracker/internal/myMd5"
)

func Md5Func() {
	fmt.Printf("%x\n", md5.Sum([]byte("adf")))

}

func MyMd5Func() {
	fmt.Printf("%x\n", myMd5.Sum([]byte("adf")))
}
