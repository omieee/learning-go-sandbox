package main

import (
	"fmt"
	"slices"
)

func addService(queue []string, name string) []string {
	if name != "" {
		return append(queue, name)
	} else {
		return nil
	}
}

func removeFirst(queue []string) ([]string, string, bool) {
	return queue[1:], queue[0], true
}

func contains(queue []string, name string) bool {
	return slices.Contains(queue, name)
}

func removeService(queue []string, name string) []string {
	newqueue := []string{}
	for _, v := range queue {
		if v != name {
			newqueue = append(newqueue, v)
		}
	}
	return newqueue
}

func main() {
	queue := []string{
		"payments-api",
		"users-api",
		"inventory-api",
	}
	queue = addService(queue, "notifications-api")
	fmt.Println("Added Notification:", queue)
	queue, _, _ = removeFirst(queue)
	fmt.Println("Removed first", queue)
	fmt.Println("Does `inventory-api` exist: ", contains(queue, "inventory-api"))
	fmt.Println("Removing `users-api`", removeService(queue, "users-api"))
}
