package model

import (
	"fmt"
	"os"
)

// Authorization 正しいユーザー(自分)かどうかを調べる
func Authorization(name, pass string) (bool, error) {
	// 現在一人用なので環境変数で指定するようにする
	fmt.Printf("name: %v, pass: %v", name, pass)
	if name == os.Getenv("NAME") && pass == os.Getenv("PASS") {
		return true, nil
	}
	return false, fmt.Errorf("UserID or password is wrong")
}
