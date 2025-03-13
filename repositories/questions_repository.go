package repositories

import (
	"database/sql"
	"jogodasperguntas/dto"
	"jogodasperguntas/models"
	"log"
	"time"
)

type QuestionsRepository struct {
	db *sql.DB
}

func NewQuestionsRepository(database *sql.DB) *QuestionsRepository {
	return &QuestionsRepository{db: database}
}

func (r *QuestionsRepository) GetAllQuestions() ([]dto.Question, error) {
	row, err := r.db.Query("select * from perguntas where deleted_at is null")
	if err != nil {
		return nil, err
	}

	var questions []dto.Question
	for row.Next() {
		var question dto.Question
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

func (r *QuestionsRepository) StoreQuestion(question dto.Question) error {
	questionModel := models.Question{
		Question:  question.Question,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := "INSERT INTO perguntas (pergunta, created_at, updated_at) VALUES (?, ?, ?)"
	_, err := r.db.Exec(
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

func (r *QuestionsRepository) FindQuestionById(id int) (*dto.Question, error) {
	question := dto.Question{}

	query := "SELECT * FROM perguntas WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(
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

func (r *QuestionsRepository) UpdateQuestion(id int, question string) error {
	query := "UPDATE perguntas SET pergunta = ? WHERE id = ? AND deleted_at IS NULL"
	_, err := r.db.Exec(query, question, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *QuestionsRepository) DeleteQuestion(id int) error {
	query := "UPDATE perguntas SET deleted_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
