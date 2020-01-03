package author

import (
	"time"

	"github.com/adrian83/library/pkg/common"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
)

type Entity struct {
	ID           string    `bson:"_id,omitempty"`
	Name         string    `bson:"name,omitempty"`
	CreationDate time.Time `bson:"creationDate,omitempty"`
}

func NewEntity(athr *Author) *Entity {
	return &Entity{
		ID:           athr.ID,
		Name:         athr.Name,
		CreationDate: athr.CreationDate,
	}
}

func NewEntityFromDoc(doc map[string]interface{}) (*Entity, error) {

	docBytes, err := bson.Marshal(doc)
	if err != nil {
		return nil, err
	}

	var entity Entity
	if err = bson.Unmarshal(docBytes, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

type Author struct {
	ID           string    `json:"id"`
	Name         string    `json:"name,omitempty"`
	CreationDate time.Time `json:"creationDate,omitempty"`
}

func NewAuthor(name string) *Author {
	return NewAuthorWithID(uuid.New().String(), name)
}

func NewAuthorFromEntity(entity *Entity) *Author {
	return &Author{
		ID:           entity.ID,
		Name:         entity.Name,
		CreationDate: entity.CreationDate,
	}
}

func NewAuthorWithID(id, name string) *Author {
	return &Author{
		ID:           id,
		Name:         name,
		CreationDate: time.Now().UTC(),
	}
}

type Authors []*Author

type AuthorsPage struct {
	*common.Page
	Authors Authors `json:"authors"`
}

func NewAuthorsPage(authors Authors, limit, offset int, total int64) *AuthorsPage {
	page := common.NewPage(limit, offset, total)
	return &AuthorsPage{
		Page:    page,
		Authors: authors,
	}
}
