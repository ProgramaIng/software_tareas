package main

import "fmt"

var (
	mapListaDeTareas = make(map[string]tarea)
)

type tarea struct {
	nombre string
	estado bool
}

func main() {

	seleccionDeTarea()
	fmt.Printf("Lista de Tareas: %v", mapListaDeTareas)
}

// seleccionDeTareas ésta funcion permite al usuario seleccionar la tarea que desea ejecutar.
func seleccionDeTarea() {
	var nombreTareaRealizar string
	fmt.Println(" Ingresa la tarea que desea realizar: Agregar, Completar, Imprimir ")
	fmt.Scanln(&nombreTareaRealizar)
	fmt.Println(" Ingresaste: ", nombreTareaRealizar)

	switch nombreTareaRealizar {
	case "Agregar":
		agregarTarea()
	case "Completar":

	case "Imprimir":

	}
}

// agregarTarea ésta función solicita al usuario la nueva tarea a añadir a la lista de tareas y la agrega al map tareas.
func agregarTarea() {

	var nombreTareaAgregar string
	fmt.Println(" Ingresa el nombre de la Nueva Tarea: ")
	fmt.Scanln(&nombreTareaAgregar)
	fmt.Println(" Ingresaste: ", nombreTareaAgregar)

	var tareaNueva tarea
	tareaNueva.nombre = nombreTareaAgregar

	mapListaDeTareas[tareaNueva.nombre] = tareaNueva

	return
}
