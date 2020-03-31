package tarea

import (
	"fmt"

	"github.com/bxcodec/faker"
)

/// interfaz /////
type User interface {
	Permisos() bool
	Name() string
}

/// estructura para Administrador : Alfonso

type Administrador struct {
	Usuario string
}

func (this Administrador) Permisos() bool {
	return true
}
func (this Administrador) Name() string {
	return this.Usuario
}

// estructura para usuario sin permisos

type logDown struct {
	Usuario string
}

func (this logDown) Permisos() bool {
	return false
}
func (this logDown) Name() string {
	return this.Usuario
}

///////_________________

func logeo(user User) string {
	if user.Permisos() == true {
		return user.Name() + " puede ingresar al portal"
	} else {
		return user.Name() + " no tiene permisos"
	}
}

func main() {
	fakerEj := Administrador{}
	err := faker.FakeData(&fakerEj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%+v", fakerEj)

	logdown := logDown{"Brayan"}
	admin := Administrador{"Alfonso"}
	fmt.Println(logeo(logdown))
	fmt.Println(logeo(admin))
}
