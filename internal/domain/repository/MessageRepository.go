package repository

import "Web-Chat/internal/domain/model"

type Message interface {
	Save(model.Message, int) error
	CheckMessage(int) ([]model.Message, error)
}
