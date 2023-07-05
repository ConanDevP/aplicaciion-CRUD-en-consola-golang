package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aplicacon_lista_de_tareas_GOLANG/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var task []tasks.Task

	info, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(bytes, &task)

		if err != nil {
			log.Fatal(err)
		}
		//	fmt.Println(task)

	} else {
		task = []tasks.Task{}
	}
	//fmt.Println(task)

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Opciones: list|add|complete|delete")
		return
	}

	switch os.Args[1] {
	case "list":
		tasks.ListTask(task)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Nombre de la tarea")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks.AddTaks(&task, name)
	

		tasks.SaveTasks(file,task)
	case "complete":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ingrese el ID de la tarea completada")
		idTaskString ,_:= reader.ReadString('\n')
		
		idTaskString = strings.TrimSpace(idTaskString)

		idInteger,err := strconv.Atoi(idTaskString)

		if err != nil{
			log.Fatal("El campo no es valido: ", err)
		}
		tasks.CompleteTask(idInteger,&task)
		tasks.SaveTasks(file, task)
	case "delete":
		reader:= bufio.NewReader(os.Stdin)
		fmt.Println("Ingresa el ID de la tarea")
		idString,_:= reader.ReadString('\n')
		idInteger,err := strconv.Atoi(idString)

		if err != nil{
			log.Fatal("Campo no valido: ", err)
		}

		tasks.DeleteTaks(&task,idInteger)
		tasks.SaveTasks(file,task)
	default:
		fmt.Println("Opciones: list|add|complete|delete")
	}

}
