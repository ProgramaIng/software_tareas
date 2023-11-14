package main

import "fmt"

var (
	sliceListaDeTarea = []tarea{}
)

type tarea struct {
	nombre string
	estado bool
}

func main() {

	seleccionDeTarea()
	fmt.Printf("Lista de Tareas: %v", sliceListaDeTarea)

}

// seleccionDeTareas ésta funcion permite al usuario seleccionar la tarea que desea ejecutar.
func seleccionDeTarea() {
	var nombreTareaRealizar string
	fmt.Println(" Ingresa la tarea que desea realizar: Agregar, Completar, Imprimir ")
	fmt.Scanln(&nombreTareaRealizar)
	fmt.Println(" Ingresaste: ", nombreTareaRealizar)

	switch nombreTareaRealizar {
	case "Agregar":

	case "Completar":
		completarTarea(sliceListaDeTarea)
	case "Imprimir":

	}
}

func completarTarea(listaDeTarea []tarea) {

	var tareaCompletar int
	fmt.Println(" Ingrese el nombre de la tarea ha completar?: ")
	fmt.Scanln(&tareaCompletar)
	fmt.Println(" Ingresaste: ", tareaCompletar)

	Completada := true

	for indice, _ := range sliceListaDeTarea {
		if tareaCompletar == indice {
			fmt.Println(" indice: ", "Completada", indice, Completada)

			return
		}
	}

}
