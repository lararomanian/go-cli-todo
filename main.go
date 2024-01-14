package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
    Text string
    Completed bool
}

func showMenu(){
    fmt.Println("\n Menu:")
    fmt.Println("1. Show Tasks")
    fmt.Println("2. Add Tasks")
    fmt.Println("3. Mark Tasks as Completed")
    fmt.Println("4. Save Tasks to File")
    fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)

    input, err := reader.ReadString('\n')

    if err != nil {
        fmt.Print(err)
        return "An error occured"
    }

    return strings.TrimSpace(input)

}

func showTasks(tasks []Task, addIndex bool){

    if len(tasks) == 0 {
        fmt.Print("No Tasks have been added yet")
        return
    }

    fmt.Println("Tasks: ")

    for i, task := range tasks {
        status :=" "

        if task.Completed {
            status = "x"
        }
        if !addIndex {
            fmt.Printf("%d. [%s] %s\n", i, status, task.Text)
        } else {
            fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
        }

    }
}

func addTasks(tasks *[]Task){
    taskText := getUserInput("Enter Text Description Here: ")
    *tasks = append(*tasks, Task{Text: taskText} )

    fmt.Println("Task Has Been Added.")
}

func markTaskComplete(tasks *[]Task){

    showTasks(*tasks, true)

    taskIndexStr := getUserInput("Enter Task Number To Mark As Completed")
    taskIndex, err := strconv.Atoi(taskIndexStr)

    if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
        fmt.Println("Invalid Task number provided")
        return
    }

    (*tasks)[taskIndex - 1].Completed = true
    fmt.Println("Task has been marked as completed")
}

func saveTasksToFile(tasks []Task){

    file, err := os.Create("./saves/tasks.txt")

    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    for _, task := range tasks {
        status := " "

        if task.Completed {
            status = "x"
        }

        file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Text))
    }
    fmt.Println("Tasks have been saved to /saves/tasks.txt")
}

func main() {
    fmt.Println("Hello, World!")

    tasks := []Task{}

    for {

        showMenu()

        option := getUserInput("Enter your choice: ")

        switch option {

        case "1":
            showTasks(tasks, false)
        case "2":
            addTasks(&tasks)
        case "3":
            markTaskComplete(&tasks)
        case "4":
            saveTasksToFile(tasks)
        case "5":
            fmt.Println("Exiting.....")
            return
        default:
            fmt.Println("Invalid option selected")
        }

    }
}
