package repository

import (
	"database/sql"
	"jogodasperguntas/internal/domain"
	"log"
	"time"
)

type questionRepository struct {
	db *sql.DB
}

type QuestionRepository interface {
	FindAll() ([]domain.Question, error)
	Store(question domain.Question) error
	FindById(id int) (*domain.Question, error)
	Update(id int, question string) error
	Delete(id int) error
}

func NewQuestionRepository(db *sql.DB) QuestionRepository {
	return &questionRepository{
		db: db,
	}
}

func (q *questionRepository) FindAll() ([]domain.Question, error) {
	row, err := q.db.Query("select * from perguntas where deleted_at is null")
	if err != nil {
		return nil, err
	}

	var questions []domain.Question
	for row.Next() {
		var question domain.Question
		if err := row.Scan(
			&question.ID,
			&question.Question,
			&question.CreatedAt,
			&question.UpdatedAt,
			&question.DeletedAt,
		); err != nil {
			log.Println("erro ao escanear a pergunta", err)
			continue
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (q *questionRepository) Store(question domain.Question) error {
	questionModel := domain.Question{
		Question:  question.Question,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := "INSERT INTO perguntas (pergunta, created_at, updated_at) VALUES ($1, $2, $3)"
	_, err := q.db.Exec(
		query,
		questionModel.Question,
		questionModel.CreatedAt,
		questionModel.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q *questionRepository) FindById(id int) (*domain.Question, error) {
	question := domain.Question{}

	query := "SELECT * FROM perguntas WHERE id = $1"
	err := q.db.QueryRow(query, id).Scan(
		&question.ID,
		&question.Question,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (q *questionRepository) Update(id int, question string) error {
	query := "UPDATE perguntas SET pergunta = $1 WHERE id = $2 AND deleted_at IS NULL"
	_, err := q.db.Exec(query, question, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *questionRepository) Delete(id int) error {
	query := "UPDATE perguntas SET deleted_at = $1 WHERE id = $2"
	_, err := q.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
