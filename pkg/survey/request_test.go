package survey

import (
	"context"
	"errors"
	"roofix/ent"
	"testing"
)

func TestReserveSurveySlot(t *testing.T) {
	ctx := context.Background()
	db := ent.GetClient()
	defer db.CloseClient()

	// Create a test user
	creatorID := testUserID(ctx, db)
	type args struct {
		ctx       context.Context
		creatorID string
		date      string
		slotID    string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
	}{
		{
			name: "reserve survey Slot successfully",
			args: args{
				ctx:       ctx,
				creatorID: creatorID,
				date:      "2023-05-01",
				slotID:    allSlots()[0].ID,
			},
			wantErr: false,
		},
		{
			name: "reserve survey Slot with existing request",
			args: args{
				ctx:       ctx,
				creatorID: creatorID,
				date:      "2023-05-01",
				slotID:    allSlots()[0].ID,
			},
			wantErr:     true,
			expectedErr: errors.New("a survey with the same date and time is already in the requesting status"),
		},
	}

	var surveyID string
	var expectedError bool

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			survey, err := Request(tt.args.ctx, tt.args.creatorID, tt.args.date, tt.args.slotID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReserveSurveySlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.expectedErr.Error() {
				t.Errorf("ReserveSurveySlot() expected error = %v, got %v", tt.expectedErr, err)
				return
			}

			if survey != "" {
				surveyID = survey
			}
			expectedError = tt.wantErr
		})
	}

	if surveyID != "" && !expectedError {
		defer func() {
			err := db.Survey.DeleteOneID(surveyID).Exec(ctx)
			if err != nil {
				t.Errorf("Failed To delete created survey after tests: %v", err)
			}
		}()
	}

	dropTestUser(ctx, db, creatorID)
}

func testUserID(ctx context.Context, db *ent.Client) string {
	testUser, _ := db.User.Create().
		SetEmail("test@example.com").
		SetPwd("password").
		Save(ctx)

	return testUser.ID
}

func dropTestUser(ctx context.Context, db *ent.Client, id string) {
	_ = db.User.DeleteOneID(id).Exec(ctx)
}
