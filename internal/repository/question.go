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
	Store(question domain.Question) (*domain.Question, error)
	FindById(id int) (*domain.Question, error)
	Update(id int, question string) (*domain.Question, error)
	Delete(id int) error
	Sort() (domain.Question, error)
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

func (q *questionRepository) Store(question domain.Question) (*domain.Question, error) {
	query := `
		INSERT INTO perguntas (pergunta, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id, pergunta, created_at, updated_at, deleted_at`

	created := &domain.Question{}
	now := time.Now()
	err := q.db.QueryRow(query, question.Question, now, now).Scan(
		&created.ID,
		&created.Question,
		&created.CreatedAt,
		&created.UpdatedAt,
		&created.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return created, nil
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

func (q *questionRepository) Update(id int, question string) (*domain.Question, error) {
	query := `
		UPDATE perguntas
		SET pergunta = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL
		RETURNING id, pergunta, created_at, updated_at, deleted_at`

	updated := &domain.Question{}
	err := q.db.QueryRow(query, question, time.Now(), id).Scan(
		&updated.ID,
		&updated.Question,
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&updated.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (q *questionRepository) Delete(id int) error {
	query := "UPDATE perguntas SET deleted_at = $1 WHERE id = $2"
	_, err := q.db.Exec(query, time.Now(), id)
	return err
}

func (q *questionRepository) Sort() (domain.Question, error) {
	question := domain.Question{}

	query := "SELECT * FROM perguntas WHERE deleted_at IS NULL ORDER BY RANDOM() LIMIT 1"
	err := q.db.QueryRow(query).Scan(
		&question.ID,
		&question.Question,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.DeletedAt,
	)
	if err != nil {
		return question, err
	}

	return question, nil
}
