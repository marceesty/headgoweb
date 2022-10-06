package main

// import (
// 	"log"
// 	"os"
// 	"text/template"
// )

// type Guestbook struct {
// 	SignatureCount int
// 	Signatures []string
// }

// func check(err error)  {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	text := "Here's my template!\n"
// 	tmpl, err := template.New("test").Parse(text)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, nil)
// 	check(err)
// }