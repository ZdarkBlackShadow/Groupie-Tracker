package utils

import "main.go/service"

func HasIdInDomain(id string, AllDomains []service.Domain) bool {
	for _, domain := range AllDomains {
		if domain.Id == id {
			return true
		}
	}
	return false
}