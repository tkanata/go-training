package card

import (
    "api/server/entities"
)

type CardInputPort interface {
    AddCardName(name string) (entities.Card, error)
}

type RoleOutputPort interface {
    OutputRole(entities.Role) error
}

type CardRepository interface {
    AddCardName(name string) (entities.Card, error)
}


