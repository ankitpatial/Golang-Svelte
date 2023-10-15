import {derived, readable, writable} from 'svelte/store';
import {page} from '$app/stores';
import {type Partner, PartnerContactType} from '$lib/models/partner';
import Step1 from './Step01CompanyInfo.svelte';
import Step2 from './Step02Contacts.svelte';
import Step3 from './Step03ServiceAreas.svelte';
import Step4 from './Step04Operations.svelte';
import type {ServiceCity} from "$lib/components/onboard/roofing";
import {PartnerType} from "$lib/graph/graphql";

export interface StepItem {
  name: string;
  component: any;
}

export interface ServiceState {
  id?: string;
  name: string;
  licenseNo?: string;
  licenseExpDate?: string;
  cities: Array<ServiceCity>;
  expand: boolean;
}

export const initialValue: Partner = {
  type: PartnerType.Solar,
  name: '',
  contacts: [
    {type: PartnerContactType.Primary, firstName: '', lastName: '', email: '', phone: ''}
  ],
  setupStepsCompleted: 0
};

export const steps = readable<Array<StepItem>>([
  {name: 'Company Details', component: Step1},
  {name: 'Points Of Contact', component: Step2},
  {name: 'Service Areas', component: Step3},
  {name: 'Operations', component: Step4},
]);

export const activeStep = writable<number>(1);

export const stepName = derived([steps, activeStep], ([$s, $a]) => {
  return $s[$a - 1].name;
});

export const partner = writable<Partner>(initialValue);

export const isProfilePage = derived(page, ($page) => {
  return $page.route.id?.endsWith('/(partner)/solar/[id]')
})

export function contactTitle(type: PartnerContactType) {
  switch (type) {
    case PartnerContactType.Primary:
      return 'Master User';
    case PartnerContactType.Accounting:
      return 'Accounting Contact';
    case PartnerContactType.Invoicing:
      return 'Invoicing Contact';
    case PartnerContactType.Operations:
      return 'Operations Contact';
  }
}

export function contactDesc(type: PartnerContactType) {
  switch (type) {
    case PartnerContactType.Primary:
      return 'This individual will be responsible for all of the account’s proceedings. They will ensure all workflow, communication and execution issues are resolved. This is our go-to contact within your organization';
    case PartnerContactType.Accounting:
      return 'This is the individual directly responsible to handle and resolving all accounting interactions';
    case PartnerContactType.Invoicing:
      return 'This individual will be responsible to handle invoicing';
    case PartnerContactType.CustomerService:
      return 'Contact information for the staff in charge of customer service. This is who the customers can call and email to communicate with';
    case PartnerContactType.Operations:
      return 'Head of the installation and operations department. The individual directly responsible over personal who contact, schedule, install, pass inspections and insure the job is successfully completed';
  }
}
