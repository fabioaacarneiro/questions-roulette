package services

import (
	"errors"
	"jogodasperguntas/config"
	"jogodasperguntas/dto"
	"jogodasperguntas/repositories"
	"log"
	"strconv"
	"strings"
)

type QuestionsService struct {
	repository *repositories.QuestionsRepository
}

func NewQuestionsService() *QuestionsService {
	return &QuestionsService{repository: repositories.NewQuestionsRepository(config.DB)}
}

func (s *QuestionsService) GetAllQuestions() ([]dto.Question, error) {
	questions, err := s.repository.GetAllQuestions()
	if err != nil {
		log.Println("erro ao buscar perguntas no service", err)
		return nil, err
	}

	return questions, nil
}

func (s *QuestionsService) StoreQuestion(question dto.Question) error {
	if strings.TrimSpace(question.Question) == "" {
		return errors.New("pergunta nao pode ser vazia")
	}

	err := s.repository.StoreQuestion(question)
	if err != nil {
		log.Println("erro ao salvar pergunta no service", err)
		return err
	}

	return nil
}

func (s *QuestionsService) FindQuestionById(id string) (*dto.Question, error) {
	questionId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("id precisa ser um inteiro")
	}

	question, err := s.repository.FindQuestionById(questionId)
	if err != nil {
		log.Println("erro ao buscar pergunta pelo id no service", err)
		return nil, err
	}

	return question, nil
}
func (s *QuestionsService) UpdateQuestion(id string, question string) error {
	questionId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id precisa ser um inteiro")
	}

	if strings.TrimSpace(question) == "" || len(question) <= 5 {
		return errors.New("pergunta precisa ter mais de 5 caracteres")
	}

	err = s.repository.UpdateQuestion(questionId, question)
	if err != nil {
		log.Println("erro ao atualizar pergunta no service", err)
		return err
	}

	return nil
}

func (s *QuestionsService) DeleteQuestion(id string) error {
	questionId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id precisa ser um inteiro")
	}

	err = s.repository.DeleteQuestion(questionId)
	if err != nil {
		log.Println("erro ao deletar pergunta no service", err)
		return err
	}

	return nil
}
