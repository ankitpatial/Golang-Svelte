<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import {SavePartnerContacts} from '$lib/gql/partner';
  import {page} from '$app/stores';
  import alerts from '$lib/stores/alerts';
  import ContactCard from '../ContactCard.svelte';
  import {activeStep, contactDesc, contactTitle, isProfilePage, partner, stepName} from './index';
  import type {PartnerContact} from "$lib/models/partner";
  import {PartnerContactType} from "$lib/models/partner";
  import {goto} from "$app/navigation";

  export let isEdit = false;

  const client = getContextClient();
  const stepNo = 2;

  let busy = false,
    contacts: Array<PartnerContact>;

  $:partnerID = $page.params.id;
  $:asProfile = $isProfilePage;
  $: {
    let list = $partner.contacts || [];
    contacts = [];
    [
      PartnerContactType.Primary,
      PartnerContactType.Accounting,
      PartnerContactType.Invoicing,
      PartnerContactType.Operations,
    ].forEach(t => {
      let c = list.find(c => c.type === t) || {
        type: t,
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        otherEmail: ''
      };
      contacts.push(c);
    })
  }

  // save Contact List
  async function submit() {
    if (busy) return;

    busy = true;
    const res = await client.mutation(SavePartnerContacts, {
      partnerID,
      contacts: contacts.map(c => ({
        type: c.type,
        id: c.id,
        userID: c.userID,
        firstName: c.firstName,
        lastName: c.lastName,
        phone: c.phone,
        email: c.email,
        otherEmail: c.otherEmail || "",
        title: c.title || undefined,
        description: c.description || undefined
      }))
    }).toPromise();
    busy = false

    if (res.error) {
      alerts.error(`Error:${$stepName}`, res.error.message);
      return;
    }

    partner.update(p => {
      if (p.setupStepsCompleted < stepNo) {
        p.setupStepsCompleted = stepNo;
      }

      p.contacts = res.data.savePartnerContacts;
      return p;
    });

    if (asProfile) {
      alerts.success($stepName, 'Saved Successfully');
      return;
    }

    $page.url.searchParams.set('step', `${$activeStep + 1}`);
    await goto($page.url.pathname + $page.url.search);

  }
</script>

<H3WithBorder>{$stepName}</H3WithBorder>
<form on:submit|preventDefault={submit}>
  <div class="flex gap-4 flex-col sm:flex-row flex-wrap mt-10" class:blur-sm={busy}>
    {#each (contacts || []) as contact, idx}
      <ContactCard
        {contact}
        {busy}
        title={contactTitle(contact.type)}
        description={contactDesc(contact.type)}
        hideAccInvEmail
      />
    {/each}
  </div>
  <BtnSubmit cls='mt-5' loading={busy}>
    {asProfile ? 'SAVE' : 'SAVE & NEXT'}
  </BtnSubmit>
</form>
