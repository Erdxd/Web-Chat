package repository

import "Web-Chat/internal/domain/model"

type Message interface {
	Save(model.Message, int) error
	CheckMessages(int) ([]model.Message, error)
}
