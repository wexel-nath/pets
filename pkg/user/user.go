package user

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Status    int64  `json:"userStatus"`
}

func NewUser(row map[string]interface{}) User {
	user := User{
		ID:        row["user_id"].(int64),
		Username:  row["username"].(string),
		FirstName: row["first_name"].(string),
		LastName:  row["last_name"].(string),
		Email:     row["email"].(string),
		Password:  row["password"].(string),
		Phone:     row["phone"].(string),
		Status:    row["status"].(int64),
	}
	return user
}

func Insert(user User) (User, error) {
	row, err := insert(user)
	if err != nil {
		return User{}, err
	}

	return NewUser(row), nil
}
