# GopherTask: Go CLI Task Manager

---

GopherTask is a simple and easy-to-use command-line interface (CLI) based task manager written in Go. It allows you to manage your tasks efficiently, providing basic functionalities such as adding tasks, deleting tasks, listing tasks, and saving tasks to a file for future use.

### Installation
To install GopherTask, you need to have Go installed on your machine. Clone the repository and build the executable using the following commands:

```
git clone https://github.com/your-username/gopherTask.git
cd gopherTask
go build
```

This will generate an executable file. You can then move the executable to a directory included in your system's PATH for easy access.

### Usage

Adding a Task
To add a task, use the add command followed by the title, description, and due date (in the format yyyy-mm-dd):

```
./gopherTask add "Task Title" "Task Description" 2024-01-15
```

Deleting a Task
To delete a task, use the delete command followed by the task ID:
```
./gopherTask delete 1
```
Listing Tasks
To list all tasks, use the list command:
```
./gopherTask list
```
Quitting the Application
To exit the application and save your tasks, use the quit or exit command:
```
./gopherTask quit
```

---

### Example

```
Welcome to the Task Manager!
--------------------------------------------
Commands: add, delete, list, quit (or exit)

> add "Buy groceries" "Milk, eggs, bread" 2024-01-20
Task added successfully!

> list
ID      Title           Description             Due Date    Completed
1       Buy groceries   Milk, eggs, bread        2024-01-20  false

> delete 1
Task deleted successfully!

> list
ID      Title           Description             Due Date    Completed
```

---

### Contribution
If you find any issues or have suggestions for improvement, feel free to open an issue or create a pull request. Your contributions are always welcome!
