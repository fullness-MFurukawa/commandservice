package main

import (
	"commandservice/presen"

	"go.uber.org/fx"
)

func main() {
	// fxを起動する
	fx.New(
		presen.CommandDepend, // 依存性を定義する
	).Run()
}
