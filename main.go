package main

import (
	IoT "CasaInteligente/IoT"
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	/**
	// Sección de las acciónes aleatoreas
	lampara := IoT.Dispositivo{
		Nombre: "Lampara",
	}
	r := rand.New(rand.NewSource(99))
	for {
		if err := randAction(&lampara, r.Float32()); err != nil {
			color.HiRed("Error: " + err.Error())
		}
		time.Sleep(2 * time.Second)
	}
	**/
	var listaDispositios []IoT.Dispositivo
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Seleccióne una opción")
		fmt.Println("1) Agregar Dispositivo")
		fmt.Println("2) Seleccionar Dispositivo")
		fmt.Println("3) Mostrar Dispositivos")
		fmt.Println("4) Salir")
		var opc int
		_, err := fmt.Scan(&opc)
		fmt.Print("\033[H\033[2J")
		if err == nil {
			switch opc {
			case 1:
				fmt.Println("Agregando Dispositivo")
				fmt.Println("Ingrese el nombre del dispositivo")
				var nombre string
				_, err := fmt.Scan(&nombre)
				if err != nil {
					color.HiRed("Error: " + err.Error())
					break
				}
				listaDispositios = append(listaDispositios, IoT.Dispositivo{Nombre: nombre})
				color.Green("Dispositivo agregado con éxito!")
			case 2:
				if len(listaDispositios) != 0 {
					fmt.Println("Seleccione un Dispositivo eligiendo el numero correspondiente")
					mostrarDispositivos(listaDispositios)
					var selected int
					_, err := fmt.Scan(&selected)
					if err != nil {
						color.HiRed("Error: " + err.Error())
						break
					}
					if selected > 0 && selected <= len(listaDispositios) {
						selected-- // Se reduce en 1 para que coincida con el indice de la lista
						fmt.Println("Nombre: " + listaDispositios[selected].Nombre)
						fmt.Println("Estado: " + listaDispositios[selected].EstadoActual())
						fmt.Println("¿Qué desea hacer?")
						fmt.Printf("1) Encender\n2) Apagar\n")
						var accion int
						fmt.Scan(&accion)
						for accion != 1 && accion != 2 {
							color.Yellow("Debe elegir entre 1) Encender y 2) Apagar")
							fmt.Scan(&accion)
						}
						if accion == 1 {
							err := listaDispositios[selected].Encender()
							if err != nil {
								color.HiRed("Error: " + err.Error())
								break
							}
							color.Green("Encendido")
						} else {
							err := listaDispositios[selected].Apagar()
							if err != nil {
								color.HiRed("Error: " + err.Error())
								break
							}
							color.Blue("Apagado")
						}
					}
				} else {
					color.Yellow("No hay dispositivos registrados!")
				}
			case 3:
				if len(listaDispositios) != 0 {
					fmt.Println("Dispositivos registrados")
					mostrarDispositivos(listaDispositios)
				} else {
					color.Yellow("No hay dispositivos registrados!")
				}
			case 4:
				fmt.Println("Adiós!")
				os.Exit(0)
			default:
				fmt.Println("Las opciones validas son entre 1 y 4")
			}
		} else {
			color.HiRed("Error: " + err.Error())
		}
		fmt.Println("Presione ENTER para continuar...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
func mostrarDispositivos(lista []IoT.Dispositivo) {
	if len(lista) != 0 {
		for i, dis := range lista {
			color.Cyan("%d: %s\n", i+1, dis.Nombre) // i se aumenta en 1 para que el conteo no comience desde 0
		}
	}
}
func randAction(d *IoT.Dispositivo, action float32) error {
	fmt.Printf("Intentando ")
	if action > 0.5 {
		fmt.Println("encender", (*d).Nombre)
		if err := (*d).Encender(); err != nil {
			return err
		} else {
			color.Green((*d).EstadoActual())
		}
	} else {
		fmt.Println("apagar", (*d).Nombre)
		if err := (*d).Apagar(); err != nil {
			return err
		} else {
			color.Blue((*d).EstadoActual())
		}
	}
	return nil
}
