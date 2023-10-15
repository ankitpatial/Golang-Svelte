import {derived, readable, writable} from 'svelte/store';
import {page} from '$app/stores';
import {type Partner, PartnerContactType} from '$lib/models/partner';
import BasicDetail from './Frm01BasicDetail.svelte';
import Contacts from './Frm02Contacts.svelte';
import CompanyDocs from './Frm03CompanyDocs.svelte';
import ServiceAreas from './Frm04ServiceAreas.svelte';
import WorkCapabilities from './Frm05WorkCapabilities.svelte';
import {PartnerType} from "$lib/graph/graphql";

export interface ServiceState {
  id?: string;
  name: string;
  licenseNo?: string;
  licenseExpDate?: string;
  cities: Array<ServiceCity>;
  expand: boolean;
}

export interface ServiceCity {
  id?: string; // postalID
  cityName: string;
  cityZip: string;
  active: boolean;
  licenseNo?: string;
  licenseProof?: string;
}


export interface StepItem {
  name: string;
  component: any;
}

export const initialValue: Partner = {
  type: PartnerType.Roofing,
  name: '',
  contacts: [
    {type: PartnerContactType.Primary, firstName: '', lastName: '', email: '', phone: ''}
  ],
  setupStepsCompleted: 0
};

export const steps = readable<Array<StepItem>>([
  {name: 'Company Details', component: BasicDetail},
  {name: 'Points Of Contact', component: Contacts},
  {name: 'Company Documents', component: CompanyDocs},
  {name: 'Service Areas', component: ServiceAreas},
  {name: 'Trades/Work Capabilities', component: WorkCapabilities}
]);


export const activeStep = writable<number>(1);

export const stepName = derived([steps, activeStep], ([$s, $a]) => {
  return $s[$a - 1].name;
});


export const partner = writable<Partner>(initialValue);

export const isProfilePage = derived(page, ($page) => {
  return $page.route.id?.endsWith('/(partner)/roofing/[id]')
})
