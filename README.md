# 📋 go-task-manager

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/sudowhiterose/go-task-manager?display_name=tag)](https://github.com/sudowhiterose/go-task-manager/releases)

Simple CLI task manager written in Go.  

---

## 🛠️ Features

- add — add a new task
- list — show all tasks
- done — mark a task as completed
- delete — remove a task by ID
---

## 🚀 Quick Start

```bash
git clone https://github.com/sudowhiterose/go-task-manager.git
./go-task-manager
```
## 📖 Usage
# Add a task
./go-task-manager add "Buy milk"

# List all tasks
./go-task-manager list

# Mark task as done
./go-task-manager done 1

# Delete a task
./go-task-manager delete 1

Example output of list:
[X] ID: 1 - Buy milk

## 📄 License
This project is licensed under the MIT License — see the LICENSE file for details.
