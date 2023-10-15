<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import {type PartnerContact, partnerContactDesc, partnerContactTitle, PartnerContactType} from '$lib/models/partner';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import {SavePartnerContacts} from '$lib/gql/partner';
  import {page} from '$app/stores';
  import alerts from '$lib/stores/alerts';
  import {goto} from '$app/navigation';
  import ContactCard from '../ContactCard.svelte';
  import {activeStep, isProfilePage, partner, stepName} from './index';

  export let isEdit = false;

  const client = getContextClient();

  let busy = false,
    contacts: Array<PartnerContact>,
    contactTypes = Object.values(PartnerContactType);

  $:partnerID = $page.params.id;
  $:asProfile = $isProfilePage;
  $: {
    contacts = $partner.contacts;
    ([
      PartnerContactType.Primary,
      PartnerContactType.Operations,
      PartnerContactType.Accounting,
      PartnerContactType.CustomerService
    ]).forEach(t => {
      if (contacts.findIndex(c => c.type === t) === -1) {
        contacts.push({type: t, firstName: '', lastName: '', email: '', phone: ''})
      }
    })
    contacts.sort((a, b) => {
      if (a.type === PartnerContactType.Custom) {
        return 0;
      }

      const idxA = contactTypes.indexOf(a.type);
      const idxB = contactTypes.indexOf(b.type);
      if (idxA < idxB) {
        return -1;
      }
      if (idxA > idxB) {
        return 1;
      }
      return 0;
    });
  }
  $: otherContactTypes = contactTypes.filter(k => (contacts || []).findIndex(c => c.type === k) === -1);

  // add new contact type to list
  function addCustomContact() {
    partner.update((p) => {
      p.contacts = (p.contacts || []).concat({
        type: PartnerContactType.Custom,
        title: 'Contact Title',
        description: 'Contact description',
        firstName: "",
        lastName: "",
        email: ""
      });
      return p;
    });
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
    busy = false;

    if (res.error) {
      alerts.error(`Error:${$stepName}`, res.error.message);
      return;
    }

    if (asProfile) {
      alerts.success($stepName, 'Saved Successfully');
      return;
    }

    partner.update(p => {
      if (p.setupStepsCompleted < 2) {
        p.setupStepsCompleted = 2;
      }

      p.contacts = res.data.savePartnerContacts;
      return p;
    });

    $page.url.searchParams.set("step", `${$activeStep + 1}`);
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
        title={partnerContactTitle(contact.type)}
        description={partnerContactDesc(contact.type)}
      />
    {/each}

    <div class="card w-96 bg-slate-50">
      <div class="card-body place-content-center text-center">
        <div class="btn btn-ghost" on:click={addCustomContact}>
          <h1 class="text-2xl font-bold">Add New Contact</h1>
          <p>click to add new contact to list</p>
        </div>
      </div>
    </div>
  </div>
  <BtnSubmit cls="mt-5" loading={busy}>
    {asProfile ? 'SAVE' : 'SAVE & NEXT'}
  </BtnSubmit>
</form>
