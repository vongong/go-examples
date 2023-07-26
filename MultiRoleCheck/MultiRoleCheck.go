package main

import "fmt"

func main() {
	member1Roles := []int{1, 2, 3, 4, 7, 10}
	member2Roles := []int{1, 10}

	accessRoles := []int{4, 7}

	fmt.Println(checkRoles(member1Roles, accessRoles))
	fmt.Println(checkRoles(member2Roles, accessRoles))
}

func checkRoles(memberRoles, accessRoles []int) bool {
	memberMap := make(map[int]bool)
	for _, v := range memberRoles {
		memberMap[v] = true
	}

	for _, v := range accessRoles {
		if _, found := memberMap[v]; !found {
			return false
		}
	}
	return true
}
