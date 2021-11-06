package handler

import "net/http"

func (h *BaseHandler) ListRecipesHandler(w http.ResponseWriter, r *http.Request) {
	n, _ := h.recipeRepository.FindAll()
	w.Write([]byte(n))
}
