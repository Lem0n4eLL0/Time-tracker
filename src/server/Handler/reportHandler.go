package handler

import (
	"fmt"
	"time"

	"net/http"

	database "timeTrackerApp/src/server/Database"
	eh "timeTrackerApp/src/server/ErrorHandler"
	s "timeTrackerApp/src/server/Structures"
	token "timeTrackerApp/src/server/Token"

	"github.com/jung-kurt/gofpdf"
)

func GetReportPDF(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	payload, err := token.GetTokenMaker().VerifyToken(c.Value)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := database.GetUserByID(payload.UserID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusNotFound, err.Error())
		return
	}

	projects, err := database.GetProjects(payload.UserID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusNotFound, err.Error())
		return
	}
	fmt.Println("Зашел")
	var tasks []s.Task
	var status string
	pdf := gofpdf.New("P", "mm", "A4", "") // 210mm x 297mm
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(0, 10, fmt.Sprintf("%v user's report for the period: %v", user.Name, time.Now()))
	pdf.Cell(0, 10, "Projects:")
	for _, pr := range projects {
		pdf.Cell(0, 10, fmt.Sprintf("\t%v %v", pr.ProjectName, pr.CreatedAt))
		pdf.Cell(0, 10, fmt.Sprintf("\tDescription%v", pr.Description))
		pdf.Cell(0, 10, "Tasks:")
		tasks, err = database.GetTasks(payload.UserID, pr.ProjectID)
		if err != nil {
			http.Error(w, "Could not generate report", http.StatusInternalServerError)
			return
		}
		for _, t := range tasks {
			if t.Status {
				status = "Complite"
			} else {
				status = "Unexecuted"
			}
			pdf.Cell(0, 10, fmt.Sprintf("\t%v  %v/%v  category: %v status: %v", t.Name, t.CreatedAt, t.EndDate, t.Category, status))
			pdf.Cell(0, 10, fmt.Sprintf("\t%v", t.Description))
		}
	}
	// Сохранить файл в буфер
	tempFile := "report.pdf"
	err = pdf.OutputFileAndClose(tempFile)
	if err != nil {
		http.Error(w, "Could not generate report", http.StatusInternalServerError)
		return
	}
	//defer os.Remove(tempFile) // Удалить временный файл после отправки

	// Отправить файл клиенту
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	http.ServeFile(w, r, tempFile)
}
