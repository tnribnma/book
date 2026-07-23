package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"book-management/repository"
	"book-management/service"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(db *sql.DB) *ReportHandler {
	return &ReportHandler{
		service: service.NewReportService(
			repository.NewReportRepository(db),
			repository.NewBookRepository(db),
			repository.NewBorrowingRepository(db),
		),
	}
}

func (h *ReportHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetDashboard(r.Context())
	if err != nil {
		Error(w, http.StatusInternalServerError, "failed to generate report")
		return
	}
	JSON(w, http.StatusOK, report)
}

func (h *ReportHandler) GetPopularBooks(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	books, err := h.service.GetPopularBooks(r.Context(), limit)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	JSON(w, http.StatusOK, books)
}

func (h *ReportHandler) GetSystemSummary(w http.ResponseWriter, r *http.Request) {
	summary, err := h.service.GetSystemSummary(r.Context())
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	JSON(w, http.StatusOK, summary)
}