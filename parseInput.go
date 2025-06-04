package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type ItemContainer struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Index    int    `json:"index"`
}

var Globalinput string
var ItemContainers []ItemContainer
var selectedIndex int

func (a *App) ParseInput(input string) *[]ItemContainer {

	input = strings.TrimSpace(input)

	Globalinput = input

	ItemContainers = ItemContainers[:0]

	luaApi(-1)

	return &ItemContainers
}

func (a *App) ClickItemContainer(index int) {

	luaApi(index)

	fmt.Println("selected index", selectedIndex)
}

func luaApi(index int) {

	selectedIndex = index

	ItemContainers = ItemContainers[:0]
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("getInput", L.NewFunction(getInput))
	L.SetGlobal("addItem", L.NewFunction(addItem))

	scriptDir := "./scripts"

	files, err := ioutil.ReadDir(scriptDir)
	if err != nil {
		log.Fatalf("Fehler beim Lesen des Ordners: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".lua" {
			scriptPath := filepath.Join(scriptDir, file.Name())
			fmt.Printf("Führe aus: %s\n", scriptPath)
			if err := L.DoFile(scriptPath); err != nil {
				log.Printf("Fehler beim Ausführen von %s: %v", scriptPath, err)
			}
		}
	}
}

func getInput(L *lua.LState) int {
	L.Push(lua.LString(Globalinput))
	return 1
}

func addItem(L *lua.LState) int {
	Title := L.ToString(1)
	SubTitle := L.ToString(1)

	index := 0
	index = len(ItemContainers)

	ItemContainers = append(ItemContainers, ItemContainer{
		Title:    Title,
		SubTitle: SubTitle,
		Index:    index,
	})

	fmt.Println("Index Item", index)

	if selectedIndex == index {
		L.Push(lua.LBool(true))
	} else {
		L.Push(lua.LBool(false))
	}

	return 1
}
