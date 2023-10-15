package server

import (
	"roofix/ent"
	"roofix/server/model"
)

func toStateList(postalList []*ent.PostalCode) []*model.State {
	var result []*model.State

	for _, c := range postalList {
		var state *model.State
		for _, s := range result {
			if s.ID == c.StateAbr {
				state = s
				break
			}
		}

		if state == nil {
			state = &model.State{
				ID:   c.StateAbr,
				Name: c.State,
			}
			result = append(result, state)
		}

		state.Cities = append(state.Cities, &model.City{
			ID:   c.ID,
			Name: &c.City,
			Zip:  &c.Code,
		})
	}
	return result
}
