package models

type User struct {
	ID    int
	Name  string `json:"name"`
	Email string `json:"email"`
	// ここに datatime 的なカラムを入れて、使い方の練習がしたい。
}
