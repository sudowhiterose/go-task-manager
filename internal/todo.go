package todo

import (
	"encoding/json"
	"fmt"
	"os"

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

// Изменили Use на "list", так как эта команда не требует аргументов
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Args:  cobra.NoArgs, // Гарантирует, что лишних аргументов нет
	Run: func(cmd *cobra.Command, args []string) {
		// Логика отображения всех задач
		fmt.Println("Listing all tasks")
	},
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1), // Защита: строго 1 аргумент
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Added task:", args[0])
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task by its ID",
	Args:  cobra.ExactArgs(1), // Защита: строго 1 аргумент
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleted task with ID:", args[0])
	},
}

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done by its ID",
	Args:  cobra.ExactArgs(1), // Защита: строго 1 аргумент
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Done task with ID:", args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Объединили всю инициализацию в один init()
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
