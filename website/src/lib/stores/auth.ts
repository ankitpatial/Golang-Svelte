/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  11/04/22, 1:20 PM
 */

import {derived, writable} from 'svelte/store';
import {PATH} from '$lib/enum';
import {PartnerType, Role, type SessionUser} from "$lib/graph/graphql";

export const authBusy = writable<boolean>(false);

export const authUser = writable<SessionUser | null>(null);

export const authUserDashboard = derived(authUser, ($user) => {
  switch ($user?.role) {
    case Role.Admin:
      return PATH.OVERVIEW;
    case Role.SiteUser:
      switch ($user?.partner?.type) {
        case PartnerType.Roofing:
          return PATH.ROOFING_PARTNER_JOBS;
        case PartnerType.Solar:
          return PATH.SOLAR_PARTNER_EST_COMPLETED;
        case PartnerType.Epc:
          return PATH.EPC_PARTNER;
        case PartnerType.Integration:
          return PATH.INTEGRATION_PARTNER;
        default:
          return PATH.UNKNOWN_PARTNER;
      }

    default:
      return PATH.HOME;
  }
});


export interface MenuCount {
  estimateCount: number;
  unassignedJobCount: number;
  assignedJobCount: number;
  paymentsPending: number;
  paymentsApproved: number;
  paymentsCompleted: number;
}

export const menuCounts = writable<MenuCount | undefined>(undefined)
