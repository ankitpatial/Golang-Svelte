<script lang='ts'>
  import type UserInput from '$lib/models/UserInput';
  import Input from '$lib/components/form/Input.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import {ALL_ACCOUNT_STATUS, All_ROLE, PATH} from '$lib/enum.js';
  import IconSave from '$lib/assets/svg/IconSave.svelte';
  import IconArrowLeft from '$lib/assets/svg/IconArrowLeft.svelte';
  import {goto} from '$app/navigation';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import TextArea from '$lib/components/form/TextArea.svelte';
  import {page} from '$app/stores';
  import SetUserPassword from "$lib/components/SetUserPassword.svelte";
  import {Role} from "$lib/graph/graphql";

  export let user: UserInput;
  export let saveHandler: (input: UserInput) => Promise<void>;

  let busy = false;

  $:isEdit = !!user.id;

  async function submit() {
    if (busy) return;

    busy = true;
    await saveHandler(user);
    busy = false;
  }

  function toUsers() {
    goto(PATH.USERS + $page.url.search);
  }
</script>

<form on:submit|preventDefault={submit}>
  <Input type='email' name='email' bind:value={user.email} label='Email' disabled={busy || isEdit} required/>
  <Input name='firstName' bind:value={user.firstName} label='First Name' disabled={busy} maxlength={50}/>
  <Input name='lastName' bind:value={user.lastName} label='Last Name' disabled={busy} maxlength={50}/>
  <InputPhone name='phone' bind:value={user.phone} label='Phone No' disabled={busy}/>
  <DropDown
    bind:value={user.role}
    options={All_ROLE.map(r => ({ id: r, name: r}))}
    label='Role'
    disabled={busy}
  />

  {#if (isEdit)}
    <DropDown
      bind:value={user.status}
      options={ALL_ACCOUNT_STATUS.map(r => ({ id: r, name: r}))}
      label='Account Status'
      placeholder='-- Select Status --'
      disabled={busy}
    />
  {/if}

  <TextArea bind:value={user.note} label='Note' disabled={busy}/>

  {#if isEdit && user.role !== Role.Admin}
    <SetUserPassword userID={user.id}/>
  {/if}
  <div class='mt-6'>
    <button class='btn btn-primary gap-4' class:loading={busy}>
      <IconSave/>
      SAVE
    </button>
    <button type='button' class='btn btn-secondary ml-2 gap-4' on:click={toUsers}>
      <IconArrowLeft/>
      Back
    </button>
  </div>
</form>
