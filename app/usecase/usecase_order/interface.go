package usecase_order

import "net/http"

type UseCase interface {
	Create(w http.ResponseWriter, r *http.Request)
}
