package main

import "fmt"

type HttpDownLoader struct {
	*template
}

func NewHttpDownLoader() Downloader {
	downloader := &HttpDownLoader{}
	template := Newtemplate(downloader)
	downloader.template = template
	return downloader
}

func (hd *HttpDownLoader) download() {
	fmt.Printf("http down %s \n", hd.url)
}

func (hd *HttpDownLoader) save() {
	fmt.Printf("http save \n")
}
