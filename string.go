package main

import (
	"fmt"
	f "fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	f.Println("c:\\Users\\app")
	f.Println(`c:\users\app`)

	str := "love you " + "vibha " + "and go " + "love love love"
	fmt.Printf("%s \n", str)
	fmt.Printf("%q", str)

	var1 := 'a'
	fmt.Printf("type %T %d \n", var1, var1)

	// str = "love"
	f.Println(len(str))
	n := utf8.RuneCountInString(str)
	f.Println(n)
	s := []rune(str)
	fmt.Println(s)
	fmt.Println(str[0])
	aa := strings.Contains(str, "love")
	fmt.Println(aa)
	ab := strings.ContainsRune("golang", 'g')
	ac := strings.ContainsAny("golang", "logs")
	fmt.Println(ac)
	ad := strings.Count(str, "e")

	fmt.Println(ab, ad)
	ae := strings.ToLower(str)
	ai := strings.ToUpper(str)
	fmt.Println(ae, ai)

	// strings repeatation

	fmt.Println(strings.Repeat("love you \n", 10))

	// replace

	f.Println(strings.Replace(str, "love", "hate", 3))
	// strings.Replace(str, "love", "hate", 1)

	// slippting

	aj := strings.Split(str, " ")
	f.Printf(" %#v \t%d\t\n", aj, len(aj))

	//joiner
	ak := strings.Join(aj, " ")
	fmt.Println(ak)

	al := strings.Fields(str)
	f.Printf("%#v \n", al)

	//remove white space and backshlash
	am := strings.TrimSpace(str)
	f.Printf("%q \n", am)

	s2 := "你好 Go!"
	s1 := []rune(s2)
	for i, v := range s1 {
		f.Println(i, v, string(s1[i]))
	}

}
func p(s string) {
	f.Println(s)
}
