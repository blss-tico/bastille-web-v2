package users

import (
	"bastille-web-v2/config"
	"log"

	"encoding/json"
	"fmt"
	"os"
)

var fileName = "./users.json"

func LoadAllUserFromFile() ([]AllUsersModel, error) {
	log.Println("LoadAllUserFromFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return []AllUsersModel{}, fmt.Errorf("Error reading file: %v", err)
	}

	var users []config.UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return []AllUsersModel{}, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	allUsers := []AllUsersModel{}
	for _, u := range users {
		allUsers = append(allUsers, AllUsersModel{Id: u.ID, Username: u.Username})
	}

	return allUsers, nil
}

func LoadUserFromFile(user config.UsersModel) error {
	log.Println("LoadUserFromFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	var users []config.UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	for _, u := range users {
		if fmt.Sprintf("%s", u.Username) == user.Username {
			if !config.CheckPasswordHashUtil(user.Password, u.Password) {
				return fmt.Errorf("Error user %s or paswword incorrect", user.Username)
			}
		}
	}

	return nil
}

func RegisterUserToFile(user config.UsersModel) error {
	log.Println("RegisterUserToFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	var users []config.UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	for _, u := range users {
		if fmt.Sprintf("%s", u.Username) == user.Username {
			return fmt.Errorf("Error user %s exists", user.Username)
		}
	}

	users = append(users, user)
	updatedBytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}

	if err := os.WriteFile(fileName, updatedBytes, 0644); err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}

func UpdateUserToFile(username string, user config.UsersModel) error {
	log.Println("UpdateUserToFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error failed to read file: %w", err)
	}

	var users []config.UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("Error failed to unmarshal JSON: %w", err)
	}

	index := 0
	for i, u := range users {
		if fmt.Sprintf("%s", u.Username) == username {
			users[i].Username = user.Username
			if user.Password != "" {
				hashed, _ := config.HashPasswordUtil(user.Password)
				users[i].Password = hashed
				index = i
			}
		}
	}

	if index == 0 {
		return fmt.Errorf("Error user %s not found", username)
	}

	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("Error failed to marshal updated JSON: %w", err)
	}

	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		return fmt.Errorf("Error failed to write file: %w", err)
	}

	return nil
}

func DeleteUserFromFile(username string) error {
	log.Println("DeleteUserFromFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error failed to read file: %w", err)
	}

	var users []config.UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("Error failed to unmarshal JSON: %w", err)
	}

	index := 0
	for i, u := range users {
		if fmt.Sprintf("%s", u.Username) == username {
			index = i
			break
		}
	}

	if index == 0 {
		return fmt.Errorf("Error user %s not found", username)
	}

	users = append(users[:index], users[index+1:]...)
	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("Error failed to marshal deleted JSON: %w", err)
	}

	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		return fmt.Errorf("Error failed to write file: %w", err)
	}

	return nil
}
