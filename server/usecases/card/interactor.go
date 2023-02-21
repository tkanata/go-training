package card

import (
    "context"
    "api/server/entities"
)

type CardInteractor struct {
    OutputPort ports.RoleOutputPort
    Repository ports.CardRepository
}

func NewUserInputPort(outputPort ports.RoleOutputPort, repository ports.CardRepository) ports.CardInputPort{
    return &UserInteractor{
        OutputPort: outputPort,
        Repository: repository,
    }
}

func (u *CardInteractor) AddCardName()


