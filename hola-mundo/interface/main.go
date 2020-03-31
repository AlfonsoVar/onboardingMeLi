package main

import "fmt"

/// interfaz /////
type User interface {
	Permisos() bool
	Usuario() string
}

/// estructura para Administrador : Alfonso

type Administrador struct {
	usuario string
}

func (this Administrador) Permisos() bool {
	return true
}
func (this Administrador) Usuario() string {
	return this.usuario
}

// estructura para usuario sin permisos

type logDown struct {
	usuario string
}

func (this logDown) Permisos() bool {
	return false
}
func (this logDown) Usuario() string {
	return this.usuario
}

///////_________________

func logeo(user User) string {
	if user.Permisos() == true {
		return user.Usuario() + " puede ingresar al portal"
	} else {
		return user.Usuario() + " no tiene permisos"
	}
}

func main() {
	logdown := logDown{"Brayan"}
	admin := Administrador{"Alfonso"}
	fmt.Println(logeo(logdown))
	fmt.Println(logeo(admin))
}
