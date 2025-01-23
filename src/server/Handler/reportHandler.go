package handler

import (
	"fmt"
	"os"
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

	var tasks []s.Task
	var status string
	pdf := gofpdf.New("P", "mm", "A4", "") // 210mm x 297mm
	pdf.AddPage()
	pdf.SetFont("Times", "B", 28)
	pdf.Cell(150, 10, fmt.Sprintf("Report %v", user.Name))
	t := time.Now()
	pdf.SetFont("Times", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("%v %v", t.Format(time.DateOnly), t.Format(time.TimeOnly)))
	pdf.Ln(15)
	for _, pr := range projects {
		pdf.SetFont("Times", "B", 16)
		pdf.Cell(30, 10, "Project:")
		pdf.SetFont("Times", "", 16)
		pdf.Cell(0, 10, pr.ProjectName)
		pdf.Ln(7)
		pdf.SetFont("Times", "B", 12)
		pdf.Cell(30, 10, "Created time: ")
		pdf.SetFont("Times", "", 12)
		pdf.Cell(0, 10, fmt.Sprintf("%v", pr.CreatedAt))
		pdf.Ln(5)
		pdf.SetFont("Times", "B", 12)
		pdf.Cell(0, 10, fmt.Sprintf("Description:"))
		pdf.Ln(5)
		pdf.SetFont("Times", "", 12)
		pdf.Cell(0, 10, fmt.Sprintf("%v", pr.Description))
		pdf.Ln(7)
		pdf.SetFont("Times", "I", 14)
		pdf.Cell(0, 10, "Tasks:")
		pdf.Ln(7)
		tasks, err = database.GetTasks(payload.UserID, pr.ProjectID)
		if err != nil {
			http.Error(w, "Could not generate report", http.StatusInternalServerError)
			return
		}
		pdf.SetFont("Times", "", 12)
		for _, t := range tasks {
			if t.Status {
				status = "Complite"
			} else {
				status = "Unexecuted"
			}
			pdf.Cell(0, 10, fmt.Sprintf("\t\tName: %v     (%v / %v)", t.Name, t.CreatedAt, t.EndDate))
			pdf.Ln(5)
			pdf.Cell(0, 10, fmt.Sprintf("\t\tCategory: %v \t\t\tStatus: %v", t.Category, status))
			pdf.Ln(5)
			pdf.Cell(0, 10, "\t\tDescription:")
			pdf.Ln(5)
			pdf.Cell(0, 10, fmt.Sprintf("\t\t%v", t.Description))
			pdf.Ln(7)
		}
		if tasks == nil {
			pdf.Cell(0, 10, "\t\tNo tasks")
			pdf.Ln(7)
		}
		pdf.Cell(0, 5, "-------------------------------------------------------------------------------------")
		pdf.Ln(7)

	}
	// Сохранить файл в буфер
	tempFile := "./report.pdf"
	err = pdf.OutputFileAndClose(tempFile)
	if err != nil {
		http.Error(w, "Could not generate report", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile) // Удалить временный файл после отправки

	// Отправить файл клиенту
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	http.ServeFile(w, r, tempFile)
}
