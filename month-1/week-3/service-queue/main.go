package main

import (
	"fmt"
	"slices"
)

func addService(queue []string, name string) []string {
	if name != "" {
		return append(queue, name)
	} else {
		return queue // If nil or anything it will overwrite the old data in queue so just return queue
	}
}

func removeFirst(queue []string) ([]string, string, bool) {
	if len(queue) > 1 {
		return queue[1:], queue[0], true // The true here means was it able to remove the first and send rest
	} else {
		return queue, queue[0], false // False here means there was only one element
	}
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
