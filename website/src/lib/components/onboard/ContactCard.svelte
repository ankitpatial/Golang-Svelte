<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang="ts">
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import Input from "$lib/components/form/Input.svelte";
  import InputAccountEmail from "$lib/components/form/InputAccountEmail.svelte";
  import InputPhone from "$lib/components/form/InputPhone.svelte";
  import TextArea from "$lib/components/form/TextArea.svelte";
  import type {PartnerContact} from "$lib/models/partner";
  import {PartnerContactType} from "$lib/models/partner";

  export let contact: PartnerContact;
  export let title: string
  export let description: string
  export let hideAccInvEmail = false;
  export let busy = false;

</script>

<div class='card w-full sm:w-96' class:blur-sm={busy}>
  <div class='card-body flex-none'>
    {#if contact?.type === PartnerContactType.Custom }
      <H3WithBorder>
        <Input bind:value={contact.title} required maxlength="50"/>
      </H3WithBorder>
    {:else }
      <H3WithBorder>{title}</H3WithBorder>
    {/if}
    {#if contact?.type === PartnerContactType.Custom }
      <TextArea label="Description" bind:value={contact.description}/>
    {:else }
      <p class="text-sm">
        {description}
      </p>
    {/if}

    <div class="mt-4">
      <Input label="First Name" bind:value={contact.firstName} required maxlength="50"/>
      <Input label="Last Name" bind:value={contact.lastName} required maxlength="50"/>
      <InputAccountEmail userID={contact.userID} bind:value={contact.email}/>
      <InputPhone label="Phone" bind:value={contact.phone} required/>
      {#if !hideAccInvEmail && contact.type === PartnerContactType.Accounting}
        <Input label="Invoicing Email" bind:value={contact.otherEmail} required maxlength="100"/>
        <p class="text-sm text-gray-400 italic -mt-2">
          The dedicated email address to send and receive all invoices/payments, etc. This could be the
          same
          email as the “Accounting Contact” email, if so, please fill in the same email again
        </p>
      {/if}
    </div>
  </div>
</div>
