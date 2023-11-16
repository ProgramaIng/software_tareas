package main

import "fmt"

type tarea struct {
	nombre string
	estado bool
}

func main() {
	sliceListaDeTarea := []tarea{
		{
			nombre: "Desayunar",
			estado: false,
		}, {
			nombre: "Almorzar",
			estado: false,
		}, {
			nombre: "Cenar",
			estado: false,
		}}

	seleccionDeTarea(sliceListaDeTarea)
	fmt.Printf("Lista de Tareas: %v", sliceListaDeTarea)
}

// seleccionDeTareas ésta funcion permite al usuario seleccionar la tarea que desea ejecutar.
func seleccionDeTarea(perrito []tarea) {
	var nombreTareaRealizar string
	fmt.Println(" Ingresa la tarea que desea realizar: Agregar, Completar, Imprimir ")
	fmt.Scanln(&nombreTareaRealizar)
	fmt.Println(" Ingresaste: ", nombreTareaRealizar)

	switch nombreTareaRealizar {
	case "Agregar":

	case "Completar":
		completarTarea(perrito)
	case "Imprimir":

	}
}

func completarTarea(listaDeTarea []tarea) {

	var tareaCompletar int
	fmt.Println(" Ingrese el número de la tarea ha completar?: ")
	fmt.Scanln(&tareaCompletar)
	fmt.Println(" Ingresaste: ", tareaCompletar)

	for indice := range listaDeTarea {
		if tareaCompletar == indice {
			listaDeTarea[indice].estado = true
			fmt.Println(" indice: Completada", indice, listaDeTarea[indice].estado)

			return
		}
	}

}
