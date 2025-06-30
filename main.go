package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

type Pedido struct {
	ID        int      `json:"id"`
	UsuarioID int      `json:"usuario_id"`
	Productos []string `json:"productos"`
	Total     float64  `json:"total"`
}

func main() {
	http.HandleFunc("/crear-usuario", crearUsuario)
	http.HandleFunc("/listar-productos", listarProductos)
	http.HandleFunc("/buscar-producto", buscarProducto)
	http.HandleFunc("/agregar-carrito", agregarCarrito)
	http.HandleFunc("/realizar-pedido", realizarPedido)
	http.HandleFunc("/ver-historial", verHistorial)
	http.HandleFunc("/realizar-pago", realizarPago)
	http.HandleFunc("/cancelar-pedido", cancelarPedido)

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func crearUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := Usuario{ID: 1, Nombre: "Cristian", Correo: "cristian@ejemplo.com"}
	json.NewEncoder(w).Encode(usuario)
}

func listarProductos(w http.ResponseWriter, r *http.Request) {
	productos := []Producto{
		{ID: 1, Nombre: "Camisa", Precio: 19.99},
		{ID: 2, Nombre: "Zapatos", Precio: 39.99},
	}
	json.NewEncoder(w).Encode(productos)
}

func buscarProducto(w http.ResponseWriter, r *http.Request) {
	producto := Producto{ID: 2, Nombre: "Zapatos", Precio: 39.99}
	json.NewEncoder(w).Encode(producto)
}

func agregarCarrito(w http.ResponseWriter, r *http.Request) {
	respuesta := map[string]string{"mensaje": "Producto agregado al carrito"}
	json.NewEncoder(w).Encode(respuesta)
}

func realizarPedido(w http.ResponseWriter, r *http.Request) {
	pedido := Pedido{ID: 1, UsuarioID: 1, Productos: []string{"Camisa", "Zapatos"}, Total: 59.98}
	json.NewEncoder(w).Encode(pedido)
}

func verHistorial(w http.ResponseWriter, r *http.Request) {
	historial := []Pedido{
		{ID: 1, UsuarioID: 1, Productos: []string{"Camisa", "Zapatos"}, Total: 59.98},
	}
	json.NewEncoder(w).Encode(historial)
}

func realizarPago(w http.ResponseWriter, r *http.Request) {
	respuesta := map[string]string{"estado": "Pago realizado con éxito"}
	json.NewEncoder(w).Encode(respuesta)
}

func cancelarPedido(w http.ResponseWriter, r *http.Request) {
	respuesta := map[string]string{"estado": "Pedido cancelado"}
	json.NewEncoder(w).Encode(respuesta)
}
