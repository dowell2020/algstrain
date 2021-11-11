/*
 * @Author: dowell87
 * @Date: 2021-10-15 17:21:50
 * @Descripttion: Algorithm training
 * @LastEditTime: 2021-11-07 23:35:47
 */
package main

import (
	"context"
	"os"

	"golang.design/x/code2img"
)

func main() {
	f, err := os.ReadFile("./algs14/4.go")
	if err != nil {
		panic(err)
	}
	code := string(f)

	// render it!
	b, err := code2img.Render(context.TODO(), code2img.LangGo, code)
	if err != nil {
		panic(err)
	}

	os.WriteFile("code.png", b, os.ModePerm)
}
