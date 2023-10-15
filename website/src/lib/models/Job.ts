/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

import type {JobProgress} from "$lib/graph/graphql";

export default interface Job {
  id?: string;
  ownerFirstName: string;
  ownerLastName: string;
  ownerEmail?: string;
  ownerPhone?: string;
  streetNumber: string;
  streetName: string;
  city: string;
  state: string;
  zip: string;
  latitude?: number;
  longitude?: number;
  repFirstName: string;
  repLastName: string;
  repEmail: string;
  repMobile: string;
  companyName: string;
  currentMaterial: string;
  newRoofingMaterial: string;
  lowSlope: boolean;
  currentMaterialLowSlope?: string;
  newRoofingMaterialLowSlope?: string;
  redeck: boolean;
  layers: 1 | 2 | 3;
  layer2Material?: string;
  layer3Material?: string;
  webhookURL?: string;
  measurementType: 'PrimaryStructureOnly' | 'PrimaryPlusDetachedGarage' | 'AllStructuresOnParcel';
  status?: string;
  statusAt?: Date;
  progress?: JobProgress;
  progressAt?: Date;
  progressFlagged?: boolean;
  price?: number;
  workOrderPrice?: number;
  materialDate?: Date;
  removeDate?: Date;
  estimates?: [object];
  note?: string;
  installDate?: Date;
  inspectionDate?: Date;
  completionDate?: Date;
  permitRequired?: boolean;
  inspectionRequired?: boolean;
  partial?: number;
  partner?: {
    id: string;
    name: string;
  };
  epc?: {
    id: string;
    name: string;
  }
}
