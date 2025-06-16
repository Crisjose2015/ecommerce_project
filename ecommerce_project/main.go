// Archivo: main.go
package main

import (
	"errors"
	"fmt"
)

// Estructura Producto
type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ReducirStock(cantidad int) error {
	if cantidad > p.Stock {
		return errors.New("no hay suficiente stock disponible")
	}
	p.Stock -= cantidad
	return nil
}

// Estructura Usuario
type Usuario struct {
	Nombre string
	Email  string
}

func (u *Usuario) MostrarDatos() {
	fmt.Println("Nombre:", u.Nombre)
	fmt.Println("Email:", u.Email)
}

// Estructura Pedido
type Pedido struct {
	Productos []Producto
	Total     float64
}

func (p *Pedido) CalcularTotal() {
	total := 0.0
	for _, producto := range p.Productos {
		total += producto.Precio
	}
	p.Total = total
}

// Estructura Carrito
type Carrito struct {
	Items []Producto
}

func (c *Carrito) AgregarProducto(p Producto) {
	c.Items = append(c.Items, p)
	fmt.Println("Producto agregado al carrito:", p.Nombre)
}

func (c *Carrito) VaciarCarrito() {
	c.Items = []Producto{}
}

// Interfaz de Pago
type Pago interface {
	Pagar(monto float64) error
}

// Implementación: Pago con Tarjeta
type PagoTarjeta struct {
	NumeroTarjeta string
}

func (pt PagoTarjeta) Pagar(monto float64) error {
	fmt.Println("Procesando pago con tarjeta por:", monto)
	return nil
}

// Implementación: Pago en Efectivo
type PagoEfectivo struct{}

func (pe PagoEfectivo) Pagar(monto float64) error {
	fmt.Println("Pago en efectivo por:", monto)
	return nil
}

func main() {
	// Crear productos
	p1 := Producto{"Laptop", 1200.0, 10}
	p2 := Producto{"Mouse", 20.0, 50}

	// Crear usuario
	usuario := Usuario{"Cristian", "cristian@example.com"}
	usuario.MostrarDatos()

	// Carrito
	carrito := Carrito{}
	carrito.AgregarProducto(p1)
	carrito.AgregarProducto(p2)

	// Crear pedido
	pedido := Pedido{Productos: carrito.Items}
	pedido.CalcularTotal()
	fmt.Println("Total del pedido:", pedido.Total)

	// Pago
	var metodoPago Pago
	metodoPago = PagoTarjeta{"1234-5678-9012-3456"}
	err := metodoPago.Pagar(pedido.Total)
	if err != nil {
		fmt.Println("Error en el pago:", err)
	} else {
		fmt.Println("Pago exitoso")
	}
}
