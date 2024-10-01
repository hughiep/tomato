package projects

import (
	"tomato/internal/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

func (r *ProjectRepository) GetProjects() []models.Project {
	var projects []models.Project
	r.DB.Find(&projects)

	return projects
}

func (r *ProjectRepository) GetProjectByID(id string) models.Project {
	var project models.Project
	r.DB.First(&project, id)

	return project
}

func (r *ProjectRepository) CreateProject(project models.Project) uint {
	r.DB.Create(&project)

	return project.ID
}

func (r *ProjectRepository) UpdateProject(id string, project models.Project) {
	r.DB.Save(&project)
}

func (r *ProjectRepository) DeleteProject(id string) {
	var project models.Project
	r.DB.Delete(&project, id)
}
