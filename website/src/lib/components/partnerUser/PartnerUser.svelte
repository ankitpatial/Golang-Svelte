<script lang='ts'>
  import { object, string } from 'yup';
  import { createForm } from 'svelte-forms-lib';
  import Input from '$lib/components/form/Input.svelte';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import type { Entity, SavePartnerContactMutation, SavePartnerContactMutationVariables } from '$lib/graph/graphql';
  import { AccountStatus, PartnerContactRole, PartnerContactType } from '$lib/graph/graphql';
  import Btn from '$lib/components/form/Btn.svelte';
  import { getContextClient } from '@urql/svelte';
  import { SavePartnerContact } from '$lib/gql/partner.contact';
  import alerts from '$lib/stores/alerts';
  import { goto } from '$app/navigation';
  import { QryUserEmailAvailable } from '$lib/gql/account';
  import { debounceInstance } from '$lib/utils/debounce';
  import { editUser } from './store';
  import { onMount } from 'svelte';
  import { authUser } from '$lib/stores/auth';

  export let partnerID;
  export let dashboardPath;
  export let isEdit = false;
  export let data = {
    id: (isEdit && $editUser?.id) || '',
    type: PartnerContactType.Custom,
    firstName: (isEdit && $editUser?.firstName) || '',
    lastName: (isEdit && $editUser?.lastName) || '',
    phone: (isEdit && $editUser?.phone) || '',
    email: (isEdit && $editUser?.email) || '',
    role: (isEdit && $editUser?.role) || PartnerContactRole.SalesRep,
    userID: (isEdit && $editUser?.userID) || undefined,
    accountStatus: (isEdit && $editUser?.accountStatus) || undefined
  };

  $:isSelf = $authUser && data && $authUser.id === data.userID;

  const client = getContextClient();
  const allUserRoles: Array<Entity> = [
    { id: PartnerContactRole.Admin, name: 'Admin' },
    { id: PartnerContactRole.SalesRep, name: 'Sales Rep' }
  ];
  const allAccStatuses: Array<Entity> = [
    { id: AccountStatus.Pending, name: 'Pending' },
    { id: AccountStatus.Active, name: 'Active' },
    { id: AccountStatus.Disabled, name: 'Disabled' }
  ];

  const emailAvailable = async (value, resolve) => {
    const res = await client.query(QryUserEmailAvailable, { id: data.userID || '', email: value }).toPromise();
    resolve(res.data?.userEmailAvailable || false);
  };

  onMount(() => {
    if (isEdit && !$editUser && dashboardPath) {
      goto(dashboardPath);
    }
  });

  let busy = false;
  const { form, errors, handleChange, handleSubmit } = createForm({
    initialValues: data,
    validationSchema: object().shape({
      firstName: string().required('First Name is required'),
      lastName: string().required('Last Name is required'),
      phone: string().required('Phone is required'),
      email: string().required('Email is required')
        .email('Invalid email')
        .test('available', 'email ${value} is not available', debounceInstance(emailAvailable)),
      role: string().required('Role is required'),
      userID: string().optional(),
      accountStatus: string().optional()
    }),
    onSubmit: async (values) => {
      if (busy) return;
      busy = true;
      const res = await client.mutation<SavePartnerContactMutation, SavePartnerContactMutationVariables>(
        SavePartnerContact,
        {
          partnerID,
          contact: values
        }).toPromise();
      busy = false;
      if (res.error) {
        alerts.error('Save User', res.error.message);
        return;
      }

      alerts.success('Save User', 'Successfully Saved');
      await back();
    }
  });

  async function back() {
    await goto(dashboardPath);
  }

</script>

{#if (isSelf)}
  <p class='badge badge-info mt-5'>
    You are viewing your own record
  </p>
{/if}
<form on:submit|preventDefault={handleSubmit} class:blur-sm={busy} class='mt-10'>
  <div class='w-full sm:w-10/12 md:w-4/12'>
    <Input
      label='First Name'
      name='firstName'
      bind:value={$form.firstName}
      err={$errors.firstName}
      {handleChange}
    />
    <Input
      label='Last Name'
      name='lastName'
      bind:value={$form.lastName}
      err={$errors.lastName}
      {handleChange}
    />
    <InputPhone
      label='Phone Number'
      name='phone'
      bind:value={$form.phone}
      err={$errors.phone}
      {handleChange}
    />
    <Input
      label='Email'
      name='email'
      bind:value={$form.email}
      err={$errors.email}
      {handleChange}
    />
    {#if isEdit}
      <DropDown
        label='Account Status'
        name='accountStatus'
        options={allAccStatuses}
        bind:value={$form.accountStatus}
        {handleChange}
      />
    {/if}
    <DropDown
      label='Role'
      name='role'
      options={allUserRoles}
      bind:value={$form.role}
      err={$errors.role}
      {handleChange}
      disabled={isEdit}
    />
  </div>
  <Btn type='submit' {busy} primary>
    SAVE
  </Btn>
</form>
