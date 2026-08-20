package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Tasker is a CLI for managing your tasks",
	Long:  `A longer description that spans multiple lines and contains examples.`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadtasks()
		if err != nil {
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Список задач пуст.")
			return
		}

		fmt.Println("Ваш список задач:")
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "X"
			}
			fmt.Printf("[%s] ID: %d - %s\n", status, t.ID, t.Text)
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadtasks()
		if err != nil {
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		newID := 1
		if len(tasks) > 0 {
			newID = tasks[len(tasks)-1].ID + 1
		}

		newTask := task{
			ID:   newID,
			Text: args[0],
			Done: false,
		}

		tasks = append(tasks, newTask)

		err = savetasks(tasks)
		if err != nil {
			fmt.Println("Ошибка保存ения задачи:", err)
			return
		}

		fmt.Printf("Задача успешно добавлена (ID: %d)\n", newID)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Ошибка: ID должен быть числом")
			return
		}

		tasks, err := loadtasks()
		if err != nil {
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		found := false
		var updatedTasks []task
		for _, t := range tasks {
			if t.ID == id {
				found = true
				continue
			}
			updatedTasks = append(updatedTasks, t)
		}

		if !found {
			fmt.Printf("Задача с ID %d не найдена\n", id)
			return
		}

		err = savetasks(updatedTasks)
		if err != nil {
			fmt.Println("Ошибка сохранения файла:", err)
			return
		}

		fmt.Printf("Задача с ID %d успешно удалена\n", id)
	},
}

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Ошибка: ID должен быть числом")
			return
		}

		tasks, err := loadtasks()
		if err != nil {
			fmt.Println("Ошибка загрузки задач:", err)
			return
		}

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Задача с ID %d не найдена\n", id)
			return
		}

		err = savetasks(tasks)
		if err != nil {
			fmt.Println("Ошибка сохранения файла:", err)
			return
		}

		fmt.Printf("Задача с ID %d отмечена как выполненная\n", id)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(doneCmd)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

const dbFile = "tasks.json"

func loadtasks() ([]task, error) {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return []task{}, nil
	}

	data, err := os.ReadFile(dbFile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []task{}, nil
	}

	var tasks []task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func savetasks(tasks []task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(dbFile, data, 0644)
}
