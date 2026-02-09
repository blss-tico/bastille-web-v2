package users

import (
	"bastille-web-v2/config"

	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func RegisterUserToFile(user config.UsersModel) error {
	fileBytes, err := os.ReadFile("users.json")
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	var users []config.UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	users = append(users, user)
	updatedBytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}

	if err := os.WriteFile(string(fileBytes), updatedBytes, 0644); err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}

func UpdateUserToFile(user config.UsersModel) error {
	fileBytes, err := os.ReadFile("users.json")
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var users []config.UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for i, u := range users {
		if fmt.Sprintf("%d", u.ID) == strconv.Itoa(user.ID) {
			users[i].Username = user.Username
			if user.Password != "" {
				hashed, _ := config.HashPasswordUtil(user.Password)
				users[i].Password = hashed
			}
		}
	}

	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal updated JSON: %w", err)
	}

	if err := os.WriteFile("users.json", updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
