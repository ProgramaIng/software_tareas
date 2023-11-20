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

}

// seleccionDeTareas ésta funcion permite al usuario seleccionar la tarea que desea ejecutar.
func seleccionDeTarea(perrito []tarea) {
	var nombreTareaRealizar string
	fmt.Println(" Ingresa la tarea que desea realizar: Agregar, Completar, Imprimir ")
	fmt.Scanln(&nombreTareaRealizar)
	fmt.Println(" Ingresaste: ", nombreTareaRealizar)

	switch nombreTareaRealizar {
	case "Agregar":
		perrito1 := agregarTarea(perrito)
		imprimirTarea(perrito1)
	case "Completar":
		completarTarea(perrito)
		imprimirTarea(perrito)
	case "Imprimir":
		imprimirTarea(perrito)
	}
}

// imprimirTarea ésta función muestra en consola cada tarea, junto a su estado (completada o no completada)
func imprimirTarea(carro []tarea) {

	for _, tarea := range carro {
		fmt.Println(fmt.Sprintf(" Tarea: %s Estado: %t", tarea.nombre, tarea.estado))

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

// agregarTarea ésta función solicita al usuario la nueva tarea a añadir a la lista de tareas y la agrega al map tareas.
func agregarTarea(listaDeTarea []tarea) []tarea {

	var nombreTareaAgregar string
	fmt.Println(" Ingresa el nombre de la Nueva Tarea: ")
	fmt.Scanln(&nombreTareaAgregar)
	fmt.Println(" Ingresaste: ", nombreTareaAgregar)

	var tareaNueva tarea
	tareaNueva.nombre = nombreTareaAgregar

	listaDeTarea = append(listaDeTarea, tareaNueva)

	return listaDeTarea
}
