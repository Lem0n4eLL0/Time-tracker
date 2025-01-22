package structures

import (
	"fmt"
)

type Task struct {
	TaskID      int    `db:"task_id" json:"id"`
	ProjectID   int    `db:"project_id" json:"projectID"`
	Name        string `db:"name" json:"task_name"`
	Description string `db:"description" json:"task_description"`
	Status      bool   `db:"status" json:"status"`
	EndDate     any    `db:"end_date" json:"end_date"`
	CreatedAt   any    `db:"created_at" json:"created_at"`
	Category    string `db:"category_name" json:"category_name"`
}

func (p *Task) Display() {
	fmt.Println(
		"TaskID: ", p.TaskID,
		"\nProjectID: ", p.ProjectID,
		"\nName: "+p.Name+
			"\nDescription: "+p.Description,
		"\nStatus: ", p.Status,
		"\nEndDate: ", p.EndDate,
		"\nCreatedAt: ", p.CreatedAt,
		"\nCategory: "+p.Category,
	)
}
