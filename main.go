package main

import (
	"github.com/cgxarrie-go/signin/cmd/signin"
)

func main() {

	// config.Instance().Bearer = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9." +
	// 	"eyJleHAiOjE3MTU4Mzg5NDYsImlzcyI6Imh0dHBzOlwvXC9iYWNrZW5kLnNp" +
	// 	"Z25pbmFwcC5jb21cLyIsImF1ZCI6Im1vYmlsZS1kZXZpY2UiLCJzdWIiOjkx" +
	// 	"NTQ5MSwiaWF0IjoxNjg0MzAyOTQ2LCJuYmYiOjE2ODQzMDI5NDZ9.OV8jldk" +
	// 	"3v3-Ib0N_lg5A14Q-FxMeD5L6MDtQvKnsxOs"

	signin.Execute()

}
