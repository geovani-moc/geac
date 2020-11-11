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
		"sair",
		"Carregar imagem base",
		"Carregar lista",
		"Visualizar certificado",
		"Exportar certificados",
		"Enviar certificados",
		"Gerenciar meu e-mail e senha",
		"Definir local de assinatura",
	}
}

//Print menu options
func (menu *Menu) Print() {

	fmt.Print("Geac\n\n")

	size := len(menu.options)
	for i := 0; i < size; i++ {
		fmt.Printf("%d - %30v (%d)\n", i, menu.options[i], i)
	}
}

//Select selection a option in menu
func (menu *Menu) Select(option int) bool {
	switch option {
	case 1:
		fmt.Println("Opção", option)
		return false
	case 2:
		fmt.Println("Opção", option)
		return true
	case 3:
		fmt.Println("Opção", option)
		return true
	case 4:
		fmt.Println("Opção", option)
		return true
	case 5:
		fmt.Println("Opção", option)
		return true
	default:
		fmt.Println("Opcção", option, "inválida.")
		return true
	}
}
