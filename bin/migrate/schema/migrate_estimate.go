package schema

import (
	"context"
)

func MigrateEstimate(ctx context.Context) {
	//db := ent.GetClient()
	//defer db.CloseClient()
	//
	//count, err := db.Estimate.Query().Count(ctx)
	//if err != nil {
	//	log.PrintError(err)
	//	return
	//}
	//
	//// pages
	//size := 1000
	//pages := count / size
	//log.InfoBullet("%d batches to migrate estimates", pages)
	//// run update in batch of 100
	//for i := 0; i <= pages; i++ {
	//	log.InfoBullet("migrating estimates batch: %d of %d", i, pages)
	//	update(ctx, db, i*size, size)
	//}
}

//func update(ctx context.Context, db *ent.Client, skip, take int) {
//	all, err := db.Estimate.Query().
//
//		WithUpdates(func(up *ent.EstimateResponseQuery) {
//			up.Order(ent.Desc(estimateresponse.FieldCreatedAt))
//		}).
//		Where(entEstimate.HasJob()).
//		Offset(skip).
//		Limit(take).
//		All(ctx)
//	if err != nil {
//		log.PrintError(err)
//		return
//	}
//	for _, est := range all {
//		estQry := est.Update()
//		j := est.Edges.Job
//		jobQry := j.Update()
//		if j != nil {
//			refCoID := str.Val(j.CompanyRefID)
//			if refCoID == "" {
//				_, refCoID = rfx.CompanyIDByName(j.CompanyName)
//			}
//
//			// find partner by external CompanyID
//			var partnerID *string
//			if refCoID != "" {
//				co, er1 := db.Partner.Query().Where(partner.ExternalID(refCoID)).Select(partner.FieldID).First(ctx)
//				if er1 != nil && ent.IsNotFound(er1) { // partner not found
//					// this must not happen, in normal proces partner:company relation is created before estimate request is made
//					//
//					// create a new solar partner
//					if id, er2 := pkgPartner.QuickCreateSolar(ctx, db, refCoID, j.CompanyName); er2 != nil {
//						log.Error(err)
//					} else {
//						partnerID = &id
//					}
//				} else if er1 != nil {
//					log.Error(er1)
//				} else {
//					partnerID = &co.ID
//				}
//			}
//
//			if refCoID != "" && partnerID == nil {
//				// ref ID is there but company is not in our system
//				// create solar partner
//				log.Info("➢ Create Company, RefID: %s Name: %s", str.Val(j.CompanyRefID), j.CompanyName)
//				if id, er1 := pkgPartner.QuickCreateSolar(ctx, db, refCoID, j.CompanyName); er1 != nil {
//					log.PrintError(er1)
//				} else {
//					partnerID = &id
//				}
//			}
//
//			salesRepID := estimate.SalesRepID(ctx, j.RepEmail)
//			if j.RepEmail != "" && salesRepID == nil && refCoID != "" {
//				log.Info("➢ Creating user %s under company %s", j.RepEmail, j.CompanyName)
//				if uID := pkgPartner.QuickAddSalesRep(
//					ctx, db, refCoID, j.RepEmail, j.RepFirstName, j.RepLastName, j.RepMobile,
//				); uID != "" {
//					salesRepID = &uID
//				}
//			}
//
//			// save HomeOwner
//			hoInp := &homeOwner.Input{
//				FirstName:    j.OwnerFirstName,
//				LastName:     j.OwnerLastName,
//				Email:        &j.OwnerEmail,
//				Phone:        &j.OwnerPhone,
//				StreetNumber: j.StreetNumber,
//				StreetName:   j.StreetName,
//				City:         j.City,
//				State:        j.State,
//				Zip:          j.Zip,
//				Latitude:     j.Latitude,
//				Longitude:    j.Longitude,
//			}
//			inpHash := homeOwner.Hash(hoInp)
//			if first, e := db.HomeOwner.Query().Where(homeowner.Hash(inpHash)).Select(homeowner.FieldID).First(ctx); e != nil && !ent.IsNotFound(e) {
//				log.Warn(fmt.Sprintf("failed to find home owner info for hash: %s with error: %v", inpHash, e))
//				continue
//			} else if first != nil {
//				// exists
//				estQry.SetHomeOwnerID(first.ID)
//				jobQry.SetHomeOwnerID(first.ID)
//			} else {
//				// need to be created
//				hoQry := homeOwner.CreateQry(db.HomeOwner.Create(), hoInp, &inpHash, partnerID)
//				if ho, er := hoQry.Save(ctx); err != nil {
//					log.Warn(fmt.Sprintf("failed to home owner info for jobID: %s with error: %v", est.ID, er))
//					continue
//				} else {
//					estQry.SetHomeOwnerID(ho.ID)
//					jobQry.SetHomeOwnerID(ho.ID)
//				}
//			}
//
//			estQry.
//				SetStatus(toStatus(j.Status)).
//				SetRegionID(postal.GetStateRegion(j.State)).
//				SetCurrentMaterial(j.CurrentMaterial).
//				SetNewRoofingMaterial(j.NewRoofingMaterial).
//				SetLowSlope(j.LowSlope).
//				SetCurrentMaterialLowSlope(j.CurrentMaterialLowSlope).
//				SetNewRoofingMaterialLowSlope(j.NewRoofingMaterialLowSlope).
//				SetRedeck(j.Redeck).
//				SetLayers(j.Layers).
//				SetLayer2Material(j.Layer2Material).
//				SetLayer3Material(j.Layer3Material).
//				SetMeasureType(toMeasureType(j.MeasurementType)).
//				SetPartialPercentage(j.PartialPercentage).
//				SetExtraChargeType(j.ExtraChargeType).
//				SetExtraCharges(j.ExtraCharges).
//				SetNillableExtraChargeNote(j.ExtraChargeNote).
//				SetNillableMaterialMappingNote(j.OverrideSummary).
//				SetNillableSalesRepID(salesRepID).
//				SetNillablePartnerID(partnerID).
//				SetNillableCompanyRefID(j.CompanyRefID).
//				SetCompanyRefName(j.CompanyName)
//
//			jobQry.
//				SetEstimateID(est.ID).
//				SetNillableSalesRepID(salesRepID).
//				SetNillableRequesterID(partnerID)
//
//			updates := est.Edges.Updates
//			if len(updates) > 0 {
//				estQry.SetEstimatorRawResponse(updates[0].Raw)
//			}
//
//			if err = estQry.Exec(ctx); err != nil {
//				log.Warn(fmt.Sprintf("failed to save %s with error: %v", est.ID, err))
//			} else {
//				if err = jobQry.Exec(ctx); err != nil {
//					log.Warn(fmt.Sprintf("failed to save %s with error: %v", est.ID, err))
//				}
//			}
//		}
//	}
//}
//
//func toStatus(status enum.JobStatus) enum.EstimateStatus {
//	switch status {
//	case enum.JobStatusEstimated:
//		return enum.EstimateStatusPending
//	case enum.JobStatusDenied:
//		return enum.EstimateStatusDenied
//	case enum.JobStatusApproved, enum.JobStatusAssigned, enum.JobStatusAccepted:
//		return enum.EstimateStatusApproved
//
//	default:
//		return enum.EstimateStatusNew
//	}
//}
//
//// toMeasureType converts the measurement type from the job to the estimate
//func toMeasureType(v uint8) enum.Measure {
//	var mt enum.JobMeasurement
//	mt.FromUInt8(v)
//	switch mt {
//	case enum.PrimaryAndDetachedGarage:
//		return enum.MeasurePrimaryPlusDetachedGarage
//	case enum.AllStructures:
//		return enum.MeasureAllStructuresOnParcel
//	default:
//		return enum.MeasurePrimaryStructureOnly
//	}
//}
