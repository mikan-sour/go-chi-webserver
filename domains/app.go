package domains

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	DB     *gorm.DB
}
