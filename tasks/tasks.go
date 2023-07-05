package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type status bool

type Task struct {
	ID       int    `json:"id`
	Name     string `jason: "name"`
	Complete status `json: "complete"`
}

func (s status) String() string {
	if s {
		return fmt.Sprintf("%s Completa","✔")
	}
	return fmt.Sprintf("%s Pendiente","✘")
}

func ListTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas...")
		return
	}

	for i, task := range tasks {
		fmt.Printf("%v)  ID: %v, Nombre Tarea:' %v ',  Estado: %v \n", i+1, task.ID, task.Name, task.Complete)
	}
}

func AddTaks(tasks *[]Task, name string)  {
	task := Task{
		ID:       generateID(*tasks),
		Name:     name,
		Complete: false,
	}

	*tasks = append(*tasks, task)

}

func generateID(tasks []Task)int{
	if len(tasks) ==0{
		return 1
	}
	return tasks[len(tasks) - 1].ID + 1
}

func SaveTasks(file *os.File, tasks []Task){
	bytes, err:= json.Marshal(tasks)

	if err != nil{
		log.Fatal(err)
	}

	_,err = file.Seek(0,0)
	if err != nil{
		log.Fatal(err)
	}
	err = file.Truncate(0)
	if err != nil{
		log.Fatal(err)
	}

	writer:= bufio.NewWriter(file)

	_,err = writer.Write(bytes)

	if err != nil{
		log.Fatal(err)
	}

	if err:= writer.Flush(); err != nil{
		log.Fatal(err)
	}
}


func DeleteTaks(tasks *[]Task, id int) {
	id = id -1
	if len(*tasks) == 0 || len(*tasks) <= id  {
		fmt.Println("No hay tareas para eliminar")
		return
	}

	*tasks = append((*tasks)[:id], (*tasks)[id+1:]...)

}


func CompleteTask(id int, tasks *[]Task){
	id = id -1
	if len(*tasks) == 0 || len(*tasks) <= id  {
		fmt.Println("No hay tarea con ese indice")
		return
	}

	(*tasks)[id].Complete = true

}