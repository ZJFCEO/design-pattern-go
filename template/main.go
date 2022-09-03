package main

func main() {
	var downloader Downloader = NewHttpDownLoader()
	downloader.Download("http://tsinghua.edu.cn/ai.zip")
	var downloader1 Downloader = NewFtpDownLoader()
	downloader1.Download("ftp://tsinghua.edu.cn/ai.zip")
}
