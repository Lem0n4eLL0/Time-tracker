package database

import (
	"context"
	"fmt"
	s "timeTrackerApp/src/server/Structures"
	"timeTrackerApp/src/utils"

	"github.com/jackc/pgx/v5"
)

func connect() (*pgx.Conn, error) {
	return DatabaseConnection()
}

func GetUserByName(username string) (*s.User, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(),
		`SELECT user_id, username, email, password_hash, created_at, updated_at, roles.name 
        FROM users 
        INNER JOIN roles 
        ON roles.role_id=users.role_id 
        WHERE username=$1 ORDER BY user_id ASC`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.User])
	if err != nil {
		return nil, err
	} else if len(users) == 0 {
		return nil, fmt.Errorf("Данный пользователь не найден")
	}
	return &users[0], nil
}

func GetUserByID(userID int) (*s.User, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(),
		`SELECT user_id, username, email, password_hash, created_at, updated_at, roles.name
  FROM users
  INNER JOIN roles ON roles.role_id=users.role_id WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.User])
	if err != nil {
		return nil, err
	} else if len(users) == 0 {
		return nil, fmt.Errorf("Данный пользователь не найден")
	}
	return &users[0], nil
}

func CreateUser(user *s.User) error {
	conn, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "INSERT INTO users (username, password_hash, email, role_id) VALUES ($1, $2, $3, $4)", user.Name, utils.Sha512Hashing(user.Password), user.Email, 2)
	if err != nil {
		return err
	}
	return nil
}

func GetProjects(userID int) ([]s.Project, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT project_id, user_id, name, description, created_at FROM public.projects
  WHERE user_id = $1
  ORDER BY projects.created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.Project])
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func CreateProject(project *s.Project) error {
	conn, err := DatabaseConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `INSERT INTO projects (user_id, name, description) VALUES ($1, $2, $3)`, project.UserID, project.ProjectName, project.Description)
	if err != nil {
		return err
	}
	return nil
}

func GetProject(userID int, projectID int) (*s.Project, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT project_id, user_id, name, description, created_at FROM public.projects
    WHERE project_id = $1 AND user_id=$2
    ORDER BY projects.created_at ASC`, projectID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	project, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.Project])
	if err != nil {
		return nil, err
	} else if len(project) == 0 {
		return nil, nil
	}
	return &project[0], nil
}

func GetLastCreateProject(userID int) (*s.Project, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT project_id, user_id, name, description, created_at FROM public.projects WHERE user_id = $1 ORDER BY projects.created_at DESC LIMIT 1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	project, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.Project])
	if err != nil {
		return nil, err
	} else if len(project) == 0 {
		return nil, fmt.Errorf("Пользователь не имеет проектов")
	}
	return &project[0], nil
}

func DeleteProject(userID, projectID int) error {
	conn, err := DatabaseConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(),
		`DELETE FROM tasks
    WHERE tasks.project_id IN (SELECT projects.project_id FROM projects WHERE user_id=$1 AND project_id=$2)`, userID, projectID)
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(),
		`DELETE FROM projects WHERE user_id=$1 AND project_id=$2`, userID, projectID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProject(userID int, project *s.Project) error {
	conn, err := DatabaseConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(),
		`UPDATE projects SET name=$1, description=$2 WHERE user_id=$3 AND project_id=$4`, project.ProjectName, project.Description, userID, project.ProjectID)
	if err != nil {
		return err
	}
	return nil
}

func GetTasks(userID int, projectID int) ([]s.Task, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT task_id, tasks.project_id, tasks.name, tasks.description, status, end_date, tasks.created_at, category_name
    FROM public.tasks
    INNER JOIN projects ON tasks.project_id = projects.project_id
    INNER JOIN categories ON tasks.category_id = categories.category_id
    WHERE projects.user_id = $1 AND tasks.project_id = $2
    ORDER BY projects.created_at DESC`, userID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.Task])
	if err != nil {
		return nil, err
	} else if len(tasks) == 0 {
		return nil, nil
	}
	return tasks, nil
}

func CreateTask(task *s.Task) error {
	conn, err := DatabaseConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	id, err := GetCategoryId(task.Category)
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), `INSERT INTO tasks (project_id, name, description, status, category_id) VALUES ($1, $2, $3, $4, $5)`,
		task.ProjectID, task.Name, task.Description, task.Status, id)
	if err != nil {
		return err
	}
	return nil
}

func GetCategoryId(categoryName string) (int, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return -1, err
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), `SELECT category_id FROM public.categories
  WHERE category_name = $1`, categoryName)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var category_id int
	for rows.Next() {
		rows.Scan(&category_id)
	}
	return category_id, nil
}

func GetLastCreateTask(projectID int) (*s.Task, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT task_id, tasks.project_id, tasks.name, tasks.description, status, end_date, tasks.created_at, category_name
    FROM public.tasks
    INNER JOIN categories ON tasks.category_id = categories.category_id
    WHERE project_id = $1 ORDER BY tasks.created_at DESC LIMIT 1`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	task, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.Task])
	if err != nil {
		return nil, err
	} else if len(task) == 0 {
		return nil, fmt.Errorf("Пользователь не имеет задач в этом проекте")
	}
	return &task[0], nil
}

func GetTaskCategories() ([]s.Category, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT * FROM public.categories
    ORDER BY category_id ASC `)
	if err != nil {
		return nil, err
	}
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.Category])
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func DeleteTask(taskID int) error {
	conn, err := DatabaseConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(),
		`DELETE FROM tasks WHERE task_id = $1 `, taskID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(task *s.Task) error {
	conn, err := DatabaseConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	CategoryID, err := GetCategoryId(task.Category)
	if err != nil {
		return err
	}
	fmt.Println(task)
	fmt.Println(task.EndDate)
	_, err = conn.Exec(context.Background(),
		`UPDATE tasks SET name=$1, description=$2, status=$3, end_date=$4, category_id=$5 WHERE task_id = $6`,
		task.Name, task.Description, task.Status, task.EndDate, CategoryID, task.TaskID)
	if err != nil {
		return err
	}
	return nil
}
