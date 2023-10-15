package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"errors"
	"fmt"
	"roofix/config"
	"roofix/ent"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/installation"
	"roofix/pkg/util/str"
	"roofix/server/generated"
	"roofix/server/model"
)

// SalesRep is the resolver for the salesRep field.
func (r *installationJobResolver) SalesRep(ctx context.Context, obj *ent.InstallationJob) (*model.Entity, error) {
	if obj == nil || obj.Edges.SalesRep == nil {
		return nil, nil
	}

	rep := obj.Edges.SalesRep
	return &model.Entity{
		ID:   rep.ID,
		Name: rep.FirstName + " " + rep.LastName,
	}, nil
}

// ImageURL is the resolver for the imageURL field.
func (r *installationJobItemResolver) ImageURL(ctx context.Context, obj *ent.InstallationJobItem) (string, error) {
	if obj == nil || obj.ImgKey == "" {
		return "", nil
	}

	return fmt.Sprintf("%s/%s", config.Read().Website.AssetUrl, obj.ImgKey), nil
}

// BookInstallation is the resolver for the bookInstallation field.
func (r *mutationResolver) BookInstallation(ctx context.Context, typeArg enum.InstallationType, pkgID string, productID *string, owner model.InstallationOwnerInput) (bool, error) {
	var requestingPartnerID *string
	var repUserID, creatorID string

	u := account.CtxUser(ctx)
	repUserID = u.ID
	creatorID = u.ID
	if u.Partner != nil {
		requestingPartnerID = &u.Partner.ID
	}

	id, err := installation.CreateJob(ctx, typeArg, pkgID, productID, &owner, requestingPartnerID, repUserID, creatorID)
	if err != nil {
		return false, err
	}

	// notify
	installation.NotifyCreated(ctx, id)
	return true, nil
}

// ApproveInstallation is the resolver for the approveInstallation field.
func (r *mutationResolver) ApproveInstallation(ctx context.Context, input model.InstallationApproveInput) (bool, error) {
	// only for ADMIN & Co. ADMIN
	if err := installation.AssertAdminOrCompanyAdmin(ctx, input.ID); err != nil {
		return false, err
	}

	var err error
	u := account.CtxUser(ctx)
	if u.IsAdmin {
		err = installation.Approve(ctx, input.ID, "", "")
	} else {
		if input.Agree == nil || !(*input.Agree) {
			return false, errors.New("must agree with terms & conditions")
		}
		err = installation.Approve(ctx, input.ID, str.Val(input.OwnerEmail), str.Val(input.OwnerPhone))
	}

	if err != nil {
		return false, err
	}

	installation.NotifyApproved(ctx, input.ID, u.Name())
	return true, nil
}

// DenyInstallation is the resolver for the denyInstallation field.
func (r *mutationResolver) DenyInstallation(ctx context.Context, id string, reason string) (bool, error) {
	// only for ADMIN & Co. ADMIN
	if err := installation.AssertAdminOrCompanyAdmin(ctx, id); err != nil {
		return false, err
	}

	if err := installation.Deny(ctx, id, reason); err != nil {
		return false, err
	}

	// notify denied
	u := account.CtxUser(ctx)
	installation.NotifyDenied(ctx, id, u.Name())
	return false, nil
}

// UndoDenyInstallation is the resolver for the undoDenyInstallation field.
func (r *mutationResolver) UndoDenyInstallation(ctx context.Context, id string) (bool, error) {
	// only for ADMIN & Co. ADMIN
	if err := installation.AssertAdminOrCompanyAdmin(ctx, id); err != nil {
		return false, err
	}

	if err := installation.RemoveDeny(ctx, id); err != nil {
		return false, err
	}

	// notify denied
	u := account.CtxUser(ctx)
	installation.NotifyRemoveDeny(ctx, id, u.Name())
	return false, nil
}

// PendingInstallations is the resolver for the pendingInstallations field.
func (r *queryResolver) PendingInstallations(ctx context.Context, typeArg enum.InstallationType, approval *enum.Approval, search *string, betweenDates []string, page model.PageInput) (*ent.InstallationJobConnection, error) {
	isAdmin, uID, pID, _ := account.Info(ctx)
	var approvalIn []enum.Approval
	if approval != nil {
		approvalIn = []enum.Approval{
			*approval,
		}
	} else {
		approvalIn = []enum.Approval{
			enum.ApprovalPending,
			enum.ApprovalApproved,
			enum.ApprovalDenied,
		}
	}

	if isAdmin {
		return installation.SearchPending(ctx, typeArg, nil, nil, approvalIn, search, betweenDates, &page)
	} else {
		return installation.SearchPending(ctx, typeArg, pID, uID, approvalIn, search, betweenDates, &page)
	}
}

// ApprovedInstallations is the resolver for the approvedInstallations field.
func (r *queryResolver) ApprovedInstallations(ctx context.Context, typeArg enum.InstallationType, status *enum.InstallationStatus, search *string, betweenDates []string, page model.PageInput) (*ent.InstallationJobConnection, error) {
	isAdmin, uID, pID, _ := account.Info(ctx)
	var statusIn []enum.InstallationStatus
	if status != nil {
		statusIn = []enum.InstallationStatus{
			*status,
		}
	} else {
		statusIn = []enum.InstallationStatus{
			enum.InstallationStatusNew,
			enum.InstallationStatusScheduled,
			enum.InstallationStatusInstalled,
		}
	}

	if isAdmin {
		return installation.SearchApproved(ctx, typeArg, nil, nil, statusIn, search, betweenDates, &page)
	} else {
		return installation.SearchApproved(ctx, typeArg, pID, uID, statusIn, search, betweenDates, &page)
	}
}

// InstallationJob returns generated.InstallationJobResolver implementation.
func (r *Resolver) InstallationJob() generated.InstallationJobResolver {
	return &installationJobResolver{r}
}

// InstallationJobItem returns generated.InstallationJobItemResolver implementation.
func (r *Resolver) InstallationJobItem() generated.InstallationJobItemResolver {
	return &installationJobItemResolver{r}
}

type installationJobResolver struct{ *Resolver }
type installationJobItemResolver struct{ *Resolver }