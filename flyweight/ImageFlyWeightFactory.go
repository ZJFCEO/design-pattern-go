package main

var imgflyfactory *ImageFlyWeightFactory

type ImageFlyWeightFactory struct {
	maps map[string]*ImageFlyWeight
}

// 对象初始化，不会反复初始化，
func GetImageFlyWeightFactory() *ImageFlyWeightFactory {
	if imgflyfactory == nil {
		imgflyfactory = &ImageFlyWeightFactory{make(map[string]*ImageFlyWeight)}
	}
	return imgflyfactory
}

func (imf *ImageFlyWeightFactory) Get(filename string) *ImageFlyWeight {
	image := imf.maps[filename] //取出，map去掉重复
	if image == nil {
		image = NewImageFlyweight(filename) //存储进入
		imf.maps[filename] = image
	}
	return image
}
