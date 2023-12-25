package practice

import "fmt"

func P1() {

	// 空のマップを作成
	emptyMap := make(map[string]int)

	// マップにキーと値を追加
	emptyMap["one"] = 1
	emptyMap["two"] = 2

	// マップの値を取得
	value := emptyMap["one"]
	fmt.Println("Value for key 'one':", value)

	// 存在しないキーへのアクセスはゼロ値を返す
	nonExistentValue := emptyMap["three"]
	fmt.Println("Value for key 'three':", nonExistentValue)

	// マップの要素を削除
	delete(emptyMap, "two")

	// キーと値のペアを直接指定してマップを初期化
	initializedMap := map[string]string{
		"name":  "John",
		"email": "john@example.com",
	}

	for key, value := range initializedMap {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

}
