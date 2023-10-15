/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

import type {PartnerType} from "$lib/graph/graphql";

export enum PartnerContactType {
  Primary = 'PRIMARY',
  Invoicing = 'INVOICING',
  Operations = 'OPERATIONS',
  Accounting = 'ACCOUNTING',
  CustomerService = 'CUSTOMER_SERVICE',
  Custom = 'CUSTOM'
}

export interface Partner {
  id?: string;
  createdAt?: Date;
  type: PartnerType;
  name: string;
  address?: string;
  latitude?: number;
  longitude?: number;
  website?: string;
  isNationWide?: boolean;
  contacts: Array<PartnerContact>;
  crewCount?: number;
  jobCapacity?: number;
  setupStepsCompleted: number;
  salesVolume?: number,
  financeOption?: number,
  downPayment?: number,
  pifDate?: number,
  installInHouse?: boolean,
  epcOption?: string,
}


export interface PartnerContact {
  type: PartnerContactType,
  id?: string;
  userID?: string;
  firstName: string;
  lastName: string;
  phone: string;
  email: string;
  otherEmail?: string;
  title?: string;
  description?: string;
}

export interface UploadedDoc {
  name: string;
  bucket: string;
  key: string;
}


export function partnerContactTitle(type: PartnerContactType) {
  switch (type) {
    case PartnerContactType.Primary:
      return 'Account Manager(Primary)';
    case PartnerContactType.Operations:
      return 'Operations/Construction';
    case PartnerContactType.Accounting:
      return 'Accounting Contact';
    case PartnerContactType.CustomerService:
      return 'Customer Service';
  }
}

export function partnerContactDesc(type: PartnerContactType) {
  switch (type) {
    case PartnerContactType.Primary:
      return 'This individual will be responsible for all of the account’s proceedings. They will ensure all workflow, communication and execution issues are resolved. This is our go-to contact within your organization';
    case PartnerContactType.Operations:
      return 'Head of the installation and operations department. The individual directly responsible over personal who contact, schedule, install, pass inspections and insure the job is successfully completed';
    case PartnerContactType.Accounting:
      return 'This is the individual directly responsible to handle and resolving all accounting interactions';
    case PartnerContactType.CustomerService:
      return 'Contact information for the staff in charge of customer service. This is who the customers can call and email to communicate with';
  }
}
