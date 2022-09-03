package main

import "fmt"

type Downloader interface {
	Download(url string)
}
type template struct {
	implement
	url string
}
type implement interface {
	download()
	save()
}

func Newtemplate(impl implement) *template {
	return &template{impl, ""}
}

func (t *template) Download(url string) {
	fmt.Println("模板下载start")
	t.url = url
	t.implement.download()
	t.implement.save()
	fmt.Println("模板下载end")
}
func (t *template) save() {
	fmt.Println("模板下载保存")

}
