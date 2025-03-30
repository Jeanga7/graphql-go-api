package db

import (
	"fmt"
	"github.com/Jeanga7/graphql-go-api.git/graph/model"
)

// Exemple de fonction pour récupérer tous les utilisateurs
func GetAllUsers() ([]*model.User, error) {
	var users []*model.User

	// Exemple de requête SQL pour récupérer les utilisateurs
	rows, err := DB.Raw("SELECT id, name, email FROM users").Rows()
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return users, nil
}
