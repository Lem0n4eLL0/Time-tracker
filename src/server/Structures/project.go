package structures

import "fmt"

type Project struct {
	ProjectID   int    `db:"project_id" json:"id"`
	UserID      int    `db:"user_id" json:"userID"`
	ProjectName string `db:"name" json:"project_name"`
	Description string `db:"description" json:"project_description"`
	CreatedAt   any    `db:"created_at" json:"created_at"`
}

func (p *Project) Display() {
	fmt.Println(
		"ProjectID: ", p.ProjectID,
		"\nUserID: ", p.UserID,
		"\nProjectName: "+p.ProjectName+
			"\nDescription: "+p.Description)
}
