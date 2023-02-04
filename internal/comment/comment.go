package comment

import "github.com/jinzhu/gorm"

// Service - struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - Defines a comment structure
type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created string
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result != nil {
		return comment, result.Error
	}
	return comment, nil
}

// GetCommentBySlug - retrives all comments by slug
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result != nil {
		return comments, result.Error
	}
	return comments, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Create(&comment); result != nil {
		return comment, result.Error
	}
	return comment, nil
}

// UpdateComment - updates comment by ID with new comments
func (s *Service) UpdateComment(ID uint, newcomment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return comment, err
	}

	if result := s.DB.Model(&comment).Update(newcomment); result.Error != nil {
		return comment, result.Error
	}

	return newcomment, nil
}

// DeleteComment - deletes a comment from database by ID
func (s *Service) DeleteComment(ID uint) error {
	comment, err := s.GetComment(ID)
	if err != nil {
		return err
	}

	if result := s.DB.Delete(&comment); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllComments - retrives all comments from the database
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
