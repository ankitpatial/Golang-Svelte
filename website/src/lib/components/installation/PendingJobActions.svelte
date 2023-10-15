<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { createEventDispatcher } from 'svelte';
  import { getContextClient } from '@urql/svelte';
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { authUser } from '$lib/stores/auth';
  import alerts from '$lib/stores/alerts';
  import BtnView from '$lib/components/form/BtnView.svelte';
  import ConfirmAction from '$lib/components/ConfirmAction.svelte';
  import IconThumbUp from '$lib/assets/svg/IconThumbUp.svelte';
  import IconThumbDown from '$lib/assets/svg/IconThumbDown.svelte';
  import Chip from '$lib/components/Chip.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import InputChk from '$lib/components/form/InputChk.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import { Approval, type InstallationJob } from '$lib/graph/graphql';
  import { ApproveInstallation, DenyInstallation, UndoDenyInstallation } from '$lib/gql/installation';

  export let item: InstallationJob | undefined = undefined;
  export let basePath = '';

  const client = getContextClient();
  const dispatch = createEventDispatcher();
  const titleRemoveDenied = 'Remove Denied Status';
  const isAdmin = $authUser?.isAdmin || $authUser?.isCompanyAdmin || false;

  let showConfirm = false;
  let busy = false;
  let epc = undefined;
  let agree = false;
  let ownerEmail;
  let ownerPhone;
  let set = false;

  $: if (item && !set) {
    ownerEmail = item?.ownerEmail;
    ownerPhone = item?.ownerPhone;
    set = true;
  }

  function view() {
    goto(`${basePath}/${item?.id}${$page.url.search}`);
  }


  // approve the estimate
  async function approve() {
    const title = 'Approve Estimate';
    const input = { id: item?.id };
    if ($authUser?.isAdmin) {
      if (await exec(title, 'Successfully approved"', ApproveInstallation, { input })) {
        dispatch('approvalChange', { id: item?.id, approval: Approval.Approved });
        return true;
      }
      return false;
    }

    if ($authUser?.partner) {
      if (!ownerEmail) {
        return Promise.reject('Please enter home owner email-address');
      }

      if (!ownerPhone) {
        return Promise.reject('Please enter home owner phone');
      }

      if (!agree) {
        return Promise.reject('Please ensure that you agree to the terms and conditions');
      }

      input.ownerEmail = ownerEmail;
      input.ownerPhone = ownerPhone;
      input.agree = agree;

      if (await exec(title, 'Successfully approved', ApproveInstallation, { input })) {
        dispatch('approvalChange', { id: item?.id, approval: Approval.Approved });
        return true;
      }
      return false;
    }

    return Promise.reject('not authorize');
  }

  // disapprove the estimate
  async function denied(note) {
    const success = await exec('Deny Estimate', 'Successfully denied', DenyInstallation, {
      id: item?.id,
      reason: note
    });

    if (success) {
      dispatch('approvalChange', { id: item?.id, approval: Approval.Denied });
    }
  }

  function confirmRemoveDenied() {
    showConfirm = true;
  }

  async function removeDenied() {
    busy = true;
    const success = await exec(titleRemoveDenied, 'removed DENIED status', UndoDenyInstallation, { id: item?.id });
    busy = false;
    if (success) {
      dispatch('approvalChange', { id: item?.id, approval: Approval.Pending });
      showConfirm = false;
    }
  }

  // update job approval
  async function exec(title: string, msg: string, mut, input: object) {
    const res = await client.mutation(mut, input).toPromise();
    if (res.error) {
      alerts.error(title, res.error.message);
      return false;
    }

    alerts.success(title, msg);
    return true;
  }
</script>

<div class='flex flex-row  w-36'>
  <BtnView on:click={view} tooltip='View Detail' cls='mt-1' />
  {#if Approval.Pending === item?.approval && isAdmin}
    <ConfirmAction title='Approve Job' confirm={approve} tooltip='Approve' okText='APPROVE'>
      <svelte:fragment slot='message'>
        {#if ($authUser.partner)}
          <div class='flex flex-col space-y-2'>
            <div>
              Are you sure you want to approve this job for installation?
            </div>
            <div>
              If so, can you confirm this customer contact information is correct?
            </div>
            <div>
              <CardWithBorderTitle title='Owner' cls='sm:w-72 sm:mr-2'>
                <InputReadonly label='Customer Full Name' value={item?.ownerName} />
                <InputReadonly label='Customer Address' value={item?.ownerAddress} />
                <Input type='email' label='Email' bind:value={ownerEmail} />
                <InputPhone label='Phone' bind:value={ownerPhone} />
              </CardWithBorderTitle>
            </div>
            <div>
              <i>
                By accepting this Proposal, Company agrees to be bound by the terms of this Proposal and of the
                Roofix General
                <a href='/terms-and-conditions' rel='noopener noreferrer' class='link link-primary' target='_blank'>
                  Terms and Conditions
                </a> for Services.
              </i>
            </div>

            <InputChk wrapperCls='w-28' label='I Agree' labelPosition='right' bind:value={agree} />
          </div>
        {:else }
          Are you sure you want to <strong>approve</strong> this job for installation?
        {/if}
      </svelte:fragment>
      <IconThumbUp color='green' />
    </ConfirmAction>
    <ConfirmAction
      title='Deny Estimate'
      message='Are you sure you want to <strong>deny</strong> this job?'
      confirm={denied}
      noteRequired
      noteTitle='Reason'
      tooltip='Deny'
      okText='DENY'
    >
      <IconThumbDown color='red' />
    </ConfirmAction>
  {:else if Approval.Approved === item?.approval}
    <Chip type='success' content='APPROVED' hideCancel />
  {:else if Approval.Denied === item?.approval}
    <Chip type='error' content='DENIED' hideCancel={!isAdmin} on:cancel={confirmRemoveDenied} />
  {/if}
</div>

<Modal
  open={showConfirm}
  title={titleRemoveDenied}
  okText='Yes'
  size='sm'
  on:ok={removeDenied}
  on:cancel={() => {showConfirm = false}}
  {busy}
>
  Are you sure you want to remove DENIED status of this job?
</Modal>
