package card

import (
    "context"
    "api/server/entities"
)

type CardInputPort interface {
    AddCardName(ctx context.Context, card *entities.Card) error
    GetCardName(ctx context.Context) error
}

type RoleOutputPort interface {
    OutputRole([]*entities.Card) error
    OutputError(error) error
}

type CardRepository interface {
    AddCardName(ctx context.Context, card *entities.Card) ([]*entities.Card, error)
    GetCardName(ctx context.Context) ([]*entities.Card, error)
}


