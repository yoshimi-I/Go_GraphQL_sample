// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph_model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

type Node interface {
	IsNode()
	// Primary key
	GetID() uuid.UUID
}

// Character: Marvel ユニバースのキャラクター
type Character struct {
	// id: Primary key
	ID uuid.UUID `json:"id"`
	// name: キャラクターの本名
	Name string `json:"name"`
	// alias: キャラクターのヒーロー名やヴィラン名
	Alias *string `json:"alias,omitempty"`
	// imageURL: キャラクター画像が格納されているオブジェクトストレージの URL
	ImageURL *string `json:"imageURL,omitempty"`
	// characterType: キャラクターの種類（ヒーロー、ヴィラン、アンチヒーロー）
	CharacterType CharacterType `json:"characterType"`
	// powers: キャラクターの能力リスト
	Powers []string `json:"powers"`
	// affiliations: キャラクターの所属グループ
	Affiliations []string `json:"affiliations"`
}

func (Character) IsNode() {}

// Primary key
func (this Character) GetID() uuid.UUID { return this.ID }

// CharacterCreateInput
type CharacterCreateInput struct {
	// name
	Name string `json:"name"`
	// alias
	Alias *string `json:"alias,omitempty"`
	// imageURL
	ImageURL *string `json:"imageURL,omitempty"`
	// characterType
	CharacterType CharacterType `json:"characterType"`
	// powers
	Powers []string `json:"powers"`
	// affiliations
	Affiliations []string `json:"affiliations"`
}

// CharacterCreatePayload: キャラクター新規作成成功時のレスポンス型
type CharacterCreatePayload struct {
	// id: Primary key
	ID uuid.UUID `json:"id"`
	// name: キャラクターの本名
	Name string `json:"name"`
	// alias: キャラクターのヒーロー名やヴィラン名
	Alias *string `json:"alias,omitempty"`
	// imageURL: キャラクター画像が格納されているオブジェクトストレージの URL
	ImageURL *string `json:"imageURL,omitempty"`
	// characterType: キャラクターの種類（ヒーロー、ヴィラン、アンチヒーロー）
	CharacterType CharacterType `json:"characterType"`
	// powers: キャラクターの能力リスト
	Powers []string `json:"powers"`
	// affiliations: キャラクターの所属グループ
	Affiliations []string `json:"affiliations"`
}

// CharacterUpdateInput
type CharacterUpdateInput struct {
	// id
	ID uuid.UUID `json:"id"`
	// name
	Name *string `json:"name,omitempty"`
	// alias
	Alias *string `json:"alias,omitempty"`
	// imageURL
	ImageURL *string `json:"imageURL,omitempty"`
	// characterType
	CharacterType *CharacterType `json:"characterType,omitempty"`
	// powers
	Powers []string `json:"powers,omitempty"`
	// affiliations
	Affiliations []string `json:"affiliations,omitempty"`
}

// CharacterUpdatePayload: キャラクター更新成功時のレスポンス型
type CharacterUpdatePayload struct {
	// id: Primary key
	ID uuid.UUID `json:"id"`
	// name: キャラクターの本名
	Name string `json:"name"`
	// alias: キャラクターのヒーロー名やヴィラン名
	Alias *string `json:"alias,omitempty"`
	// imageURL: キャラクター画像が格納されているオブジェクトストレージの URL
	ImageURL *string `json:"imageURL,omitempty"`
	// characterType: キャラクターの種類（ヒーロー、ヴィラン、アンチヒーロー）
	CharacterType CharacterType `json:"characterType"`
	// powers: キャラクターの能力リスト
	Powers []string `json:"powers"`
	// affiliations: キャラクターの所属グループ
	Affiliations []string `json:"affiliations"`
}

type Mutation struct {
}

type Query struct {
}

// CharacterType: キャラクターの種類を表す列挙型
type CharacterType string

const (
	CharacterTypeHero     CharacterType = "HERO"
	CharacterTypeVillain  CharacterType = "VILLAIN"
	CharacterTypeAntiHero CharacterType = "ANTI_HERO"
)

var AllCharacterType = []CharacterType{
	CharacterTypeHero,
	CharacterTypeVillain,
	CharacterTypeAntiHero,
}

func (e CharacterType) IsValid() bool {
	switch e {
	case CharacterTypeHero, CharacterTypeVillain, CharacterTypeAntiHero:
		return true
	}
	return false
}

func (e CharacterType) String() string {
	return string(e)
}

func (e *CharacterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CharacterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CharacterType", str)
	}
	return nil
}

func (e CharacterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}