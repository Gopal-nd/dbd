package project

import (
	"errors"
	"fmt"
)

func Taskers() {
	fmt.Print("Welcome to CLI Task Manager")
	fmt.Println("1. to add new task")
	fmt.Println("2. to remove a task")
	fmt.Println("3. complete a task")
	fmt.Println("4. Print tasks")
	fmt.Println("5. Exit")

	tasks := &Tasks{}

	for {
		fmt.Println("1. to add new task")
		fmt.Println("2. to remove a task")
		fmt.Println("3. complete a task")
		fmt.Println("4. Print tasks")
		fmt.Println("5. Print tasks")
		var input int
		fmt.Scan(&input)
		if input == 5 {
			break
		}
		switch input {
		case 1:
			{
				newTask := &Task{}
				fmt.Println("Enetr the Title :")
				fmt.Scan(&newTask.name)
				fmt.Println("Enetr the Discription :")
				fmt.Scan(&newTask.discription)
				fmt.Println("Is it completd :")
				fmt.Scan(&newTask.completd)
				newTask.id = len(tasks.tasks)

				tasks.tasks = append(tasks.tasks, *newTask)
				return

			}
		case 2:
			{
				var id int
				fmt.Println("Enetr the Id To remove :")
				fmt.Scan(&id)
				err := tasks.Remove(id)
				if err != nil {
					fmt.Println(err)
				}
			}
		case 3:
			{
				var id int
				fmt.Println("Enetr the Id To Mark Completed :")
				fmt.Scan(&id)
				err := tasks.Complete(id)
				if err != nil {
					fmt.Println(err)
				}
			}
		case 4:
			{
				tasks.PrintAll()
			}
		case 5:
			{
				return
			}
		default:
			{
				continue
			}
		}
	}

}

type Task struct {
	id          int
	name        string
	discription string
	completd    bool
}

type Tasks struct {
	tasks []Task
}

func (t *Tasks) Add(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *Tasks) Remove(id int) error {

	if id == len(t.tasks)-1 {
		t.tasks = t.tasks[:len(t.tasks)-1]
		return nil
	}

	for _, v := range t.tasks {
		if v.id == id {
			t.tasks = append(t.tasks[:id], t.tasks[id+1:]...)
			return nil
		}
	}
	return errors.New("Id not found")
}

func (t *Tasks) Complete(id int) error {

	for _, v := range t.tasks {
		if v.id == id {
			if v.completd {
				v.completd = false
			} else {
				v.completd = true
			}
			return nil
		}
	}
	return errors.New("Id Not found")
}

func (t *Tasks) PrintAll() {
	fmt.Println("**************")
	for i, v := range t.tasks {
		fmt.Println("***Taks ", i, " :***")
		fmt.Println("ID :", v.id)
		fmt.Println("Title :", v.name)
		fmt.Println("Desciption :", v.discription)
		fmt.Println("Completd :", v.completd)

	}
}
