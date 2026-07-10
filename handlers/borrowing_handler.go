package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	//"strconv"

	"book-management/middleware"
	"book-management/models"
	"book-management/service"
)

type BorrowingHandler struct {
	service *service.BorrowingService
}

func NewBorrowingHandler(db *sql.DB) *BorrowingHandler {
	return &BorrowingHandler{service: service.NewBorrowingService(db)}
}

func (h *BorrowingHandler) IssueBook(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	var req models.BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	borrowing, err := h.service.IssueBook(r.Context(), req, userID)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	JSON(w, http.StatusCreated, borrowing)
}

func (h *BorrowingHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var req models.ReturnRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	fine, err := h.service.ReturnBook(r.Context(), req.BorrowingID)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	JSON(w, http.StatusOK, map[string]float64{"fine": fine})
}

func (h *BorrowingHandler) GetMyBorrowings(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	borrowings, err := h.service.GetUserBorrowings(r.Context(), userID)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	JSON(w, http.StatusOK, borrowings)
}