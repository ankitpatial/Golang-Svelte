<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {goto} from '$app/navigation';
  import {getContextClient} from '@urql/svelte';
  import {InvitePartner} from '$lib/gql/partner';
  import alerts from '$lib/stores/alerts';
  import Input from '$lib/components/form/Input.svelte';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import InputAccountEmail from '$lib/components/form/InputAccountEmail.svelte';
  import DDLPartnerType from '$lib/components/form/DDLPartnerType.svelte';
  import InputPartnerName from '$lib/components/form/InputPartnerName.svelte';
  import {PartnerType} from "$lib/graph/graphql";

  export let id: string;
  export let partnerType: PartnerType | undefined = undefined;
  export let returnURL = '';

  const client = getContextClient();

  let busy = false;
  let data = {
    type: partnerType || PartnerType.Roofing,
    companyName: '',
    contactID: undefined,
    userID: '',
    firstName: '',
    lastName: '',
    email: '',
    phone: ''
  };

  async function submit() {
    const input = {id, ...data, createdAt: undefined};
    const res = await client.mutation(InvitePartner, {input}).toPromise();
    if (res.error) {
      alerts.error('Invite Partner', res.error.message);
      return;
    }
    // data = res.data.invitePartner;
    alerts.success('Invite Partner', `Successfully sent an onboarding invite to ${data.companyName}`);
    back();
  }

  function back() {
    if (!returnURL) return;

    goto(returnURL);
  }
</script>

<div class='w-96' class:blur-sm={busy}>
  <form on:submit|preventDefault={submit}>
    <h2>Company</h2>
    {#if !partnerType }
      <DDLPartnerType bind:value={data.type}/>
    {/if}
    <InputPartnerName partnerType={data.type} partnerID={id} bind:value={data.companyName}/>

    <h2>Account Manager</h2>
    <Input label='First Name' required bind:value={data.firstName}/>
    <Input label='Last Name' required bind:value={data.lastName}/>
    <InputAccountEmail userID={data.userID} bind:value={data.email}/>
    <InputPhone label='Phone' bind:value={data.phone}/>

    <button class='btn btn-primary uppercase'>
      Submit
    </button>
  </form>
</div>
