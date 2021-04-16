package models

type Team struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Coach string `json:"coach"`
	Email string `json:"email"`
}

type CreateTeamInput struct {
	Name  string `json:"name" binding:"required"`
	Coach string `json:"coach" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateTeamInput struct {
	Name  string `json:"name"`
	Coach string `json:"coach"`
	Email string `json:"email"`
}
