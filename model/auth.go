package model

import (
	"fmt"
	"os"
)

// User ユーザー情報管理(今は自分だけ
type User struct {
	Name string
}

// Me 自分を表す
var Me = User{}

// Authorization 正しいユーザー(自分)かどうかを調べる
func Authorization(name, pass string) (bool, error) {
	// 現在一人用なので環境変数で指定するようにする
	fmt.Printf("name: %v, pass: %v", name, pass)
	if name == os.Getenv("NAME") && pass == os.Getenv("PASS") {
		Me.Name = name
		return true, nil
	}
	return false, fmt.Errorf("UserID or password is wrong")
}
