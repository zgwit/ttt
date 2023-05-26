package ttt

import (
	"embed"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

//go:embed all:test
var File embed.FS

func TTT() {
	log.Println(File.ReadFile("test/a.txt"))
}
