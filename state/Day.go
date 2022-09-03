package main

import "fmt"

type Sunday struct {
}

func (sd *Sunday) Today() {
	fmt.Println("Sunday\n")
}
func (sd *Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

type Monday struct {
}

func (sd *Monday) Today() {
	fmt.Println(" Monday\n")
}
func (sd *Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct {
}

func (sd *Tuesday) Today() {
	fmt.Println(" Monday\n")
}
func (sd *Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct {
}

func (sd *Wednesday) Today() {
	fmt.Println(" Wednesday \n")
}
func (sd *Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct {
}

func (sd *Thursday) Today() {
	fmt.Println(" Thursday\n")
}
func (sd *Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct {
}

func (sd *Friday) Today() {
	fmt.Println(" Friday\n")
}
func (sd *Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct {
}

func (sd *Saturday) Today() {
	fmt.Println(" Saturday\n")
}
func (sd *Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}
