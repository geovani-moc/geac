package cmd

import (
	"fmt"
)

// Menu structure
type Menu struct {
	options []string
}

//NewMenu create new menu options
func (menu *Menu) NewMenu() {
	menu.options = []string{
		"Carregar imagem base",
		"Carregar lista",
		"Visualizar certificado",
		"Enviar certificados",
		"Gerenciar meu e-mail e senha",
		"Definir local de assinatura",
		"sair",
	}
}

//Print menu options
func (menu *Menu) Print() {

	fmt.Print("Geac\n\n")

	size := len(menu.options)
	for i := 0; i < size; i++ {
		fmt.Printf("%d - %30v\n", i, menu.options[i])
	}
}

//Select selection a option in menu
func (menu *Menu) Select(option int) {
	switch option {
	case 1:
		fmt.Println("Opção", option)
	case 2:
		fmt.Println("Opção", option)
	case 3:
		fmt.Println("Opção", option)
	case 4:
		fmt.Println("Opção", option)
	case 5:
		fmt.Println("Opção", option)
	default:
		fmt.Println("Opcção", option, "inválida.")
	}
}
