package repository

import "Web-Chat/internal/domain/model"

type Message interface {
	Save(model.Message) error
	Last(count int) ([]model.Message, error)
}
