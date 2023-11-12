package entity

import (
	"testing"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var fake = faker.New()

func fakeInput() *dto.CreateUserInput {
	return &dto.CreateUserInput{
		Email:       fake.Internet().Email(),
		Username:    fake.Internet().User(),
		DisplayName: fake.Internet().User(),
		Password:    fake.Internet().Password(),
	}
}

func TestNewUser(t *testing.T) {
	assert := assert.New(t)
	input := fakeInput()
	user, err := NewUser(*input)

	assert.Nil(err)
	assert.Equal(user.Email, input.Email)
	assert.Equal(user.Username, input.Username)
	assert.Equal(user.DisplayName, input.DisplayName)
	assert.Equal(user.Password, input.Password)

	input.Email = ""
	user, err = NewUser(*input)
	assert.NotNil(err)
}
