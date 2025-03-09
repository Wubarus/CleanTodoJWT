package domain

// TaskList represent users lists of tasks
type TaskList struct {
	User  User   `json:"user" db:"user_id"`
	Tasks []Task `json:"tasks" db:"task_id"`
}

// Task are contained by TaskList
type Task struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Desc  string `json:"description" db:"description"`
	Done  bool   `json:"status" db:"done"`
}

// User is the subject in the service
type User struct {
	Id       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
