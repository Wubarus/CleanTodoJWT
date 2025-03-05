package domain

type TaskList struct {
	User  User   `json:"user" db:"user_id"`
	Tasks []Task `json:"tasks" db:"task_id"`
}

type Task struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Desc  string `json:"description" db:"description"`
	Done  bool   `json:"status" db:"done"`
}

type User struct {
	Id       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
