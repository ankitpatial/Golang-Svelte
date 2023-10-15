package seed

import (
	"context"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
)

func OptionList(ctx context.Context) {
	db := ent.GetClient()
	defer db.CloseClient()

	if c, _ := db.OptionList.Query().Count(ctx); c > 0 { // already seeded
		return
	}

	inp := []*ent.OptionListCreate{
		db.OptionList.Create().SetType(enum.OptionListFinance).SetName("GoodLeap").SetDisplayName("GoodLeap").SetOrder(1),
		db.OptionList.Create().SetType(enum.OptionListFinance).SetName("Mosaic").SetDisplayName("Mosaic").SetOrder(2),
		db.OptionList.Create().SetType(enum.OptionListFinance).SetName("Dividend").SetDisplayName("Dividend").SetOrder(3),
		db.OptionList.Create().SetType(enum.OptionListFinance).SetName("Sunbright").SetDisplayName("Sunbright").SetOrder(3),

		db.OptionList.Create().SetType(enum.OptionListEPC).SetName("LGCY").SetDisplayName("LGCY").SetOrder(1),
		db.OptionList.Create().SetType(enum.OptionListEPC).SetName("V3").SetDisplayName("V3").SetOrder(2),
		db.OptionList.Create().SetType(enum.OptionListEPC).SetName("Titan").SetDisplayName("Titan").SetOrder(3),
		db.OptionList.Create().SetType(enum.OptionListEPC).SetName("FreedomForever").SetDisplayName("Freedom Forever").SetOrder(4),
		db.OptionList.Create().SetType(enum.OptionListEPC).SetName("TriSMART").SetDisplayName("TriSMART").SetOrder(5),
		db.OptionList.Create().SetType(enum.OptionListEPC).SetName("Action").SetDisplayName("Action").SetOrder(6),
		db.OptionList.Create().SetType(enum.OptionListEPC).SetName("Titanium").SetDisplayName("Titanium").SetOrder(7),
	}

	if err := db.OptionList.CreateBulk(inp...).Exec(ctx); err != nil {
		log.PrintError(err)
	} else {
		log.Info("seeded default %q & %q options", "Finance", "EPV")
	}
}
