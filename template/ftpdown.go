package main

import "fmt"

type FtpDownLoader struct {
	*template
}

func NewFtpDownLoader() Downloader {
	downloader := &FtpDownLoader{}
	template := Newtemplate(downloader)
	downloader.template = template
	return downloader
}

func (hd *FtpDownLoader) download() {
	fmt.Printf("ftp down %s \n", hd.url)
}

func (hd *FtpDownLoader) save() {
	fmt.Printf("ftp save %s \n")
}
