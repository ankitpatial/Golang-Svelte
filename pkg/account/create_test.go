package account

import (
	"context"
	"github.com/jaswdr/faker"
	"roofix/pkg/enum"
	"testing"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	// nil data will return error
	if _, err := Create(ctx, nil); err == nil {
		t.Error("must have failed for nil input")
	}

	// empty data will return struct validation errors
	if _, err := Create(ctx, &CreateUserInput{}); err == nil {
		t.Error("must have failed for da")
	}

	fake := faker.New()
	inp := &CreateUserInput{
		Email:     fake.Internet().Email(),
		FirstName: fake.Person().FirstName(),
		LastName:  fake.Person().LastName(),
		Role:      enum.RoleSiteUser,
	}

	// must save without err
	if _, err := Create(ctx, inp); err != nil {
		t.Error(err)
	}

}
