package utils

import "main.go/service"

func HasIsInElement(id string, AllElements []service.Element) bool {
	for _, element := range AllElements {
		if id == element.Id {
			return true
		}
	}
	return false
}
