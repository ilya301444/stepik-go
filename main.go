/*
 * Этот проект не более чем очередной Hello World, однако с помощью него
 * мы рассмотрим организацию кода в проекте на Go. Подробнее об организации
 * кода вы можете почитать в официальной документации https://golang.org/doc/code.html.
 * В директории gopher содержится одноименный пакет, вы можете увидеть, что не все
 * объекты из него доступны в этом пакете. И если отдельно gopher - пакет, то вместе с
 * пакетом main это уже модуль, т.к. в этой директории вы можете найти файлы go.mod и
 * go.sum, которые содержат описание зависимостей модуля.
 */

package main

// Импортируемые пакеты
import (
	"fmt"
	"time" // <= пакеты стандартной библиотеки импортируется без указания пути, Go сам знает, где его искать

	"stepik-go/gopher" // <= это пакет не из стандартной библиотеки, к нему указан путь
	/*
	 * stepik-go - корневая директория модуля
	 * gopher - поддиректория, в которой находится пакет gopher
	 */

	log "github.com/ivahaev/go-logger" // <= этот пакет тоже не из стандартной библиотеки
	/*
	 * Мы могли воспользоваться просто fmt или log из стандартной библиотеки, но тогда не увидели бы
	 * список зависимостей в файле go.mod
	 */)

func main() {
	fmt.Printf("Program starting in %s\n", time.Now().Format("15:04:05"))

	gopher := gopher.NewGopher()
	if err := gopher.Go(); err != nil {
		log.JSON(err)
	}
}
