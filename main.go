package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Idnum = 0

type Jsn struct {
	KullanciList []Kullanici
}
type Kullanici struct {
	Id       int
	Username string
	Mail     string
}

func NewJsn() Jsn {
	jsn := Jsn{}
	return jsn
}
func NewKullanici() Kullanici {
	kl := Kullanici{}
	return kl
}
func KullaniciEkle(jsnp *Jsn, kul *Kullanici) {
	jsnp.KullanciList = append(jsnp.KullanciList, *kul)

}

var (
	User     Kullanici
	username string
	email    string
)

func main() {
	Scanr := bufio.NewScanner(os.Stdin)
	Jsonfile := NewJsn()
	for i := 0; i < 10; i++ {
		fmt.Println("Lutfen UserName Girin")
		Scanr.Scan()
		User = NewKullanici()
		username = Scanr.Text()
		User.Username = username
		fmt.Println("Lutfen Email Girin")
		Scanr.Scan()
		email = Scanr.Text()
		User.Mail = email
		User.Id = Idnum
		KullaniciEkle(&Jsonfile, &User)
		Idnum++
	}

	Doc, _ := os.Create("DOC.json")
	defer Doc.Close()
	writer := io.Writer(Doc)
	encoder := json.NewEncoder(writer)

	encoder.Encode(Jsonfile)
}
