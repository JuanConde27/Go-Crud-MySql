package models

import "gorm.io/gorm"

type Persona struct {
    gorm.Model

    Email    string `json:"email" gorm:"uniqueIndex;not null"`
    Password string `json:"password" gorm:"not null"`
    Nombre   string `json:"nombre" gorm:"not null"`
    Apellido string `json:"apellido" gorm:"not null"`
    Edad     string `json:"edad" gorm:"not null"`
}
