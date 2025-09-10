package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/chhz0/projectx-go/pkg/app"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// type Task struct {
// 	ID          int        `json:"id"`
// 	Description string     `json:"description"`
// 	Priority    string     `json:"priority"`
// 	Completed   bool       `json:"completed"`
// 	CreatedAt   time.Time  `json:"created_at"`
// 	CompletedAt *time.Time `json:"completed_at,omitempty"`
// }

// type TaskManager struct {
// 	Tasks    []Task `json:"tasks"`
// 	NextID   int    `json:"next_id"`
// 	FilePath string `json:"-"`
// }

// var (
// 	taskFile    string
// 	priority    string
// 	showAll     bool
// 	taskManager *TaskManager
// )

// var rootCmd = &app.Command{
// 	Use:   "taskman",
// 	Short: "A personal task manager",
// 	Long: `taskman is a CLI task manager that helps you organize your work.

// Store tasks locally with priorities, mark them complete, and keep
// track of your productivity over time.`,
// 	Init: func(cmd *cobra.Command) {
// 		cmd.PersistentFlags().StringVar(&taskFile, "file", "", "task file (default is $HOME/.taskman.json)")

// 		config := app.NewConfig("taskman", "yaml", ".")
// 		config.Viper.SetDefault("priority", "medium")
// 		config.Viper.SetDefault("file", filepath.Join(".", ".taskman.json"))

// 		config.Read(func(err error) {
// 			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
// 				fmt.Println("config file not found error")
// 			}
// 		})
// 	},
// 	PreRun:  loadTasks,
// 	PostRun: saveTasks,

// 	SubCommands: []*app.Command{
// 		addCmd,
// 		listCmd,
// 		completeCmd,
// 		deleteCmd,
// 	},
// }

// var addCmd = &app.Command{
// 	Use:   "add [description]",
// 	Short: "Add a new task",
// 	Long:  "Add a new task with description and optional priority (high, medium, low)",
// 	Args:  cobra.MinimumNArgs(1),
// 	Init: func(cmd *cobra.Command) {
// 		cmd.Flags().StringVarP(&priority, "priority", "p", "medium", "task priority (high, medium, low)")
// 	},
// 	Run: addTask,
// }

// var listCmd = &app.Command{
// 	Use:   "list",
// 	Short: "List all tasks",
// 	Long:  "List tasks with optional filtering by completion status",
// 	Init: func(cmd *cobra.Command) {
// 		cmd.Flags().BoolVarP(&showAll, "all", "a", false, "show completed tasks too")
// 	},
// 	Run: listTasks,
// }

// var completeCmd = &app.Command{
// 	Use:   "complete [task-id]",
// 	Short: "Mark a task as completed",
// 	Args:  cobra.ExactArgs(1),
// 	Run:   completeTask,
// }

// var deleteCmd = &app.Command{
// 	Use:   "delete [task-id]",
// 	Short: "Delete a task",
// 	Args:  cobra.ExactArgs(1),
// 	Run:   deleteTask,
// }

// func getTaskFile() string {
// 	if taskFile != "" {
// 		return taskFile
// 	}
// 	return viper.GetString("file")
// }

// func loadTasks(cmd *cobra.Command, args []string) {
// 	file := getTaskFile()
// 	taskManager = &TaskManager{
// 		Tasks:    make([]Task, 0),
// 		NextID:   1,
// 		FilePath: file,
// 	}

// 	if data, err := os.ReadFile(file); err == nil {
// 		json.Unmarshal(data, taskManager)
// 	}
// }

// func saveTasks(cmd *cobra.Command, args []string) {
// 	data, err := json.MarshalIndent(taskManager, "", "  ")
// 	if err != nil {
// 		fmt.Printf("Error saving tasks: %v\n", err)
// 		return
// 	}

// 	os.WriteFile(taskManager.FilePath, data, 0644)
// }

// func addTask(cmd *cobra.Command, args []string) {
// 	description := strings.Join(args, " ")

// 	// Validate priority
// 	validPriorities := map[string]bool{"high": true, "medium": true, "low": true}
// 	if !validPriorities[priority] {
// 		fmt.Printf("Invalid priority '%s'. Use: high, medium, or low\n", priority)
// 		os.Exit(1)
// 	}

// 	task := Task{
// 		ID:          taskManager.NextID,
// 		Description: description,
// 		Priority:    priority,
// 		Completed:   false,
// 		CreatedAt:   time.Now(),
// 	}

// 	taskManager.Tasks = append(taskManager.Tasks, task)
// 	taskManager.NextID++

// 	fmt.Printf("Added task #%d: %s [%s]\n", task.ID, task.Description, task.Priority)
// }

// func listTasks(cmd *cobra.Command, args []string) {
// 	if len(taskManager.Tasks) == 0 {
// 		fmt.Println("No tasks found.")
// 		return
// 	}

// 	fmt.Printf("%-4s %-10s %-50s %-10s %s\n", "ID", "STATUS", "DESCRIPTION", "PRIORITY", "CREATED")
// 	fmt.Println(strings.Repeat("-", 90))

// 	for _, task := range taskManager.Tasks {
// 		if !showAll && task.Completed {
// 			continue
// 		}

// 		status := "PENDING"
// 		if task.Completed {
// 			status = "DONE"
// 		}

// 		fmt.Printf("%-4d %-10s %-50s %-10s %s\n",
// 			task.ID,
// 			status,
// 			truncate(task.Description, 50),
// 			strings.ToUpper(task.Priority),
// 			task.CreatedAt.Format("2006-01-02"))
// 	}
// }

// func completeTask(cmd *cobra.Command, args []string) {
// 	id, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Printf("Invalid task ID: %s\n", args[0])
// 		os.Exit(1)
// 	}

// 	for i := range taskManager.Tasks {
// 		if taskManager.Tasks[i].ID == id {
// 			if taskManager.Tasks[i].Completed {
// 				fmt.Printf("Task #%d is already completed\n", id)
// 				return
// 			}

// 			now := time.Now()
// 			taskManager.Tasks[i].Completed = true
// 			taskManager.Tasks[i].CompletedAt = &now

// 			fmt.Printf("Completed task #%d: %s\n", id, taskManager.Tasks[i].Description)
// 			return
// 		}
// 	}

// 	fmt.Printf("Task #%d not found\n", id)
// 	os.Exit(1)
// }

// func deleteTask(cmd *cobra.Command, args []string) {
// 	id, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Printf("Invalid task ID: %s\n", args[0])
// 		os.Exit(1)
// 	}

// 	for i, task := range taskManager.Tasks {
// 		if task.ID == id {
// 			// Remove task from slice
// 			taskManager.Tasks = append(taskManager.Tasks[:i], taskManager.Tasks[i+1:]...)
// 			fmt.Printf("Deleted task #%d: %s\n", id, task.Description)
// 			return
// 		}
// 	}

// 	fmt.Printf("Task #%d not found\n", id)
// 	os.Exit(1)
// }

// func truncate(s string, max int) string {
// 	if len(s) <= max {
// 		return s
// 	}
// 	return s[:max-3] + "..."
// }

// func main() {
// 	_ = rootCmd.Exec(context.Background())
// }
