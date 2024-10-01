package projects

import (
	"tomato/internal/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

type ProjectRepositoryInterface interface {
	GetProjects() []models.Project
	GetProjectByID(id string) models.Project
	CreateProject(project ProjectRequest) uint
	UpdateProject(id string, project ProjectRequest)
	DeleteProject(id string)
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

func (r *ProjectRepository) CreateProject(project ProjectRequest) uint {
	tx := r.DB.Create(&models.Project{
		Name:        project.Name,
		UserID:      project.UserID,
		Description: project.Description,
	})

	return uint(tx.RowsAffected)
}

func (r *ProjectRepository) UpdateProject(id string, project ProjectRequest) {
	var projectModel models.Project
	r.DB.First(&projectModel, id)

	projectModel.Name = project.Name
	projectModel.UserID = project.UserID
	projectModel.Description = project.Description

	r.DB.Save(&projectModel)
}

func (r *ProjectRepository) DeleteProject(id string) {
	var project models.Project
	r.DB.First(&project, id)

	r.DB.Delete(&project)
}
