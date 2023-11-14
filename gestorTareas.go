package main

import "fmt"

type tarea struct {
	nombre string
	estado bool
}

func main() {
	var sliceListaDeTarea []tarea
	seleccionDeTarea(sliceListaDeTarea)
}

// seleccionDeTareas ésta funcion permite al usuario seleccionar la tarea que desea ejecutar.
func seleccionDeTarea(sliceListaDeTarea []tarea) {
	var nombreTareaRealizar string
	fmt.Println(" Ingresa la tarea que desea realizar: Agregar, Completar, Imprimir ")
	fmt.Scanln(&nombreTareaRealizar)
	fmt.Println(" Ingresaste: ", nombreTareaRealizar)

	switch nombreTareaRealizar {
	case "Agregar":

	case "Completar":
		listaDeTarea := []tarea{}
		imprimirTarea(listaDeTarea)
	case "Imprimir":
		listaDeTarea := []tarea{}
		imprimirTarea(listaDeTarea)

	}
}

// imprimirTarea ésta función muestra en consola cada tarea, junto a su estado (completada o no completada)
func imprimirTarea(listaDeTarea []tarea) {

	for _, tarea := range listaDeTarea {
		fmt.Println(fmt.Printf(" Tarea: %v Estado: %v", tarea.nombre, tarea.estado))

	}

}
