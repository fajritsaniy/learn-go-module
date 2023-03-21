/*
GOPATH
- Ketergantungan dengan satu folder
- Pengelolaan dependency kurang efisien ketika ada update
*/

/*
Go Modules
- Management Dependency Golang -> seperti maven pada java
- Untuk get dari target package
	-  go get -u github.com/google/uuid
*/

package learngomodule

import "fmt"

func PrintSliceOfString(slc []string) {
	for _, v := range slc {
		fmt.Println(v)
	}
}
