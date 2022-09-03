package main

import "fmt"

type ImageFlyWeight struct {
	data string
}

// 初始化
func NewImageFlyweight(filename string) *ImageFlyWeight {
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyWeight{data}
}
func (i *ImageFlyWeight) Data() string { //截取数据
	return i.data
}
