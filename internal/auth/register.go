package auth

import (
	"fmt"

	"github.com/fulsep/koda-b4-playground/internal/lib"
)

type User struct {
	email string
	password string
	profile Profile
}

type Profile struct {
	FullName string
}

func Register(){
	loop := true
	for loop {
		lib.CLS()
		fmt.Printf("Masukkan nama lengkap anda: ")
		name,_ := lib.Input("string")
		fmt.Printf("Masukkan email anda: ")
		email,_ := lib.Input("string")
		fmt.Printf("Masukkan password anda: ")
		password,_ := lib.Input("string")

		user := User{
			email: email.(string),
			password: password.(string),
			profile: Profile{
				FullName: name.(string),
			},
		}
		lib.CLS()
		fmt.Println(user)

		fmt.Print("Apakah sudah benar? (y/n): ")
		confirm,_ := lib.Input("bool")
		if(confirm.(bool)){
			loop = false
		}
	}
}