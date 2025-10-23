package main

import (
	"fmt"
	"os"

	"github.com/fulsep/koda-b4-playground/internal/auth"
	"github.com/fulsep/koda-b4-playground/internal/lib"
)

type MENU struct {
	title string
	process func()
}

func main(){
	MAIN_MENU := []MENU {
		{
			title: "Register",
			process: auth.Register,
		},
	}
	for{
		lib.CLS()
		fmt.Println("--- Welcome to system ---")
		for x:=range len(MAIN_MENU){
			fmt.Printf("%d. %v\n",x+1,MAIN_MENU[x].title)
		}
		fmt.Println()
		fmt.Println("0. Exit")
		rawChoice,_ := lib.Input("int")
		choice := rawChoice.(int)
		if(choice == 0){
			os.Exit(0)
		} else if (choice <= len(MAIN_MENU)){
			MAIN_MENU[choice-1].process()
		}
	}
}