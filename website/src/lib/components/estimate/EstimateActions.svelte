<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import {fade} from 'svelte/transition';
  import {createEventDispatcher} from 'svelte';
  import {getContextClient} from '@urql/svelte';
  import {goto} from '$app/navigation';
  import {page} from '$app/stores';
  import {authUser} from '$lib/stores/auth';
  import alerts from '$lib/stores/alerts';
  import BtnView from '$lib/components/form/BtnView.svelte';
  import ConfirmAction from '$lib/components/ConfirmAction.svelte';
  import IconThumbUp from '$lib/assets/svg/IconThumbUp.svelte';
  import IconThumbDown from '$lib/assets/svg/IconThumbDown.svelte';
  import Chip from '$lib/components/Chip.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import InputChk from '$lib/components/form/InputChk.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import {type Estimate, EstimateStatus, JobStatus} from '$lib/graph/graphql';
  import {ApproveEstimate, DenyEstimate, RemoveDenied} from "$lib/gql/estimate";

  export let item: Estimate | undefined = undefined;
  export let basePath = '';
  export let epcs = [];

  const client = getContextClient();
  const dispatch = createEventDispatcher();
  const titleRemoveDenied = 'Remove Denied Status';
  const isAdminOrCoAdmin = $authUser?.isAdmin || $authUser?.isCompanyAdmin || false;

  let showConfirm = false;
  let busy = false;
  let epc = undefined;
  let agree = false;
  let ownerEmail;
  let ownerPhone;
  let set = false;

  $: if (item && !set) {
    ownerEmail = item?.homeOwner?.email;
    ownerPhone = item?.homeOwner?.phone;
    set = true;
  }

  function view() {
    goto(`${basePath}/${item?.id}${$page.url.search}`);
  }

  // update Approve/Disapprove
  async function exec(mut: any, vars, status: string, title: string, msg: string) {
    const res = await client.mutation(mut, vars).toPromise();
    if (res.error) {
      alerts.error(title, res.error.message);
      return;
    }

    alerts.success(title, msg);
    dispatch('changeStatus', {id: item?.id, status: item?.status});
  }

  // approve the estimate
  function approve() {
    const title = 'Approve Estimate';

    if ($authUser?.isAdmin) {
      return exec(ApproveEstimate, {input: {id: item?.id}}, JobStatus.Approved, title, 'Successfully approved"');
    }

    if ($authUser?.partner) {
      if (!ownerEmail) {
        return Promise.reject('Please enter home owner email-address');
      }

      if (!ownerPhone) {
        return Promise.reject('Please enter home owner phone');
      }

      if (epcs.length > 0 && !epc) {
        return Promise.reject('Please ensure that you have selected an EPC');
      }

      if (!agree) {
        return Promise.reject('Please ensure that you agree to the terms and conditions');
      }

      const input = {
        id: item?.id,
        ownerEmail,
        ownerPhone,
        epc,
        agree
      };

      return exec(ApproveEstimate, {id: item?.id, input}, JobStatus.Approved, title, 'Successfully approved"');
    }


    return Promise.reject('not authorize');
  }

  // disapprove the estimate
  function deny(note) {
    return exec(DenyEstimate, {input: {id: item?.id, note}}, JobStatus.Denied, 'Deny Estimate', 'Successfully denied');
  }

  function confirmRemoveDenied() {
    showConfirm = true;
  }

  async function removeDenied() {
    busy = true;
    const res = await client.mutation(RemoveDenied, {id: item?.id}).toPromise();
    busy = false;
    if (res.error) {
      alerts.error(titleRemoveDenied, res.error.message);
      return;
    }

    dispatch('changeStatus', {id: item?.id, status: EstimateStatus.Pending});
    showConfirm = false;
  }

</script>

<div class='flex flex-row  w-36' transition:fade>
  <BtnView on:click={view} tooltip='View Detail' cls='mt-1'/>
  {#if isAdminOrCoAdmin && (EstimateStatus.Pending === item?.status)}
    <ConfirmAction
      title='Approve Estimate'
      confirm={approve}
      tooltip='Approve'
      okText='APPROVE'
    >
      <svelte:fragment slot='message'>
        {#if ($authUser.partner)}
          <div class='flex flex-col space-y-2'>
            <div>
              Are you sure you want to approve this job for production?
            </div>
            <div>
              If so, can you confirm this customer contact information is correct?
            </div>
            <div>
              <div class='flex flex-row flex-wrap'>
                <CardWithBorderTitle title='Owner' cls='sm:w-72 sm:mr-2'>
                  <InputReadonly label='First Name' value={item?.homeOwner?.firstName}/>
                  <InputReadonly label='Last Name' value={item?.homeOwner?.lastName}/>
                  <Input type='email' label='Email' bind:value={ownerEmail}/>
                  <InputPhone label='Phone' bind:value={ownerPhone}/>
                </CardWithBorderTitle>
                <CardWithBorderTitle title='Address' cls='sm:w-72'>
                  <InputReadonly label='Street Number' value={item?.homeOwner?.streetNumber}/>
                  <InputReadonly label='Street Name' value={item?.homeOwner?.streetName}/>
                  <InputReadonly label='City' value={item?.homeOwner?.city}/>
                  <InputReadonly label='State' value={item?.homeOwner?.state}/>
                  <InputReadonly label='Zip' value={item?.homeOwner?.zip}/>
                </CardWithBorderTitle>
              </div>
            </div>
            {#if (epcs.length > 0)}
              <div>
                <DropDown
                  label='Who is the EPC for this job?'
                  options={[{id: undefined, name: ''}].concat(epcs)}
                  bind:value={epc}
                  cls='mt-8'
                />
              </div>
            {/if}
            <div>
              <i>
                By accepting this Proposal, Company agrees to be bound by the terms of this Proposal and of the
                Roofix General
                <a href='/terms-and-conditions' rel='noopener noreferrer' class='link link-primary' target='_blank'>
                  Terms and Conditions
                </a> for Services.
              </i>
            </div>

            <InputChk wrapperCls='w-28' label='I Agree' labelPosition='right' bind:value={agree}/>
          </div>
        {:else }
          Are you sure you want to <strong>approve</strong> this estimate?
        {/if}
      </svelte:fragment>
      <IconThumbUp color='green'/>
    </ConfirmAction>
    <ConfirmAction
      title='Deny Estimate'
      message='Are you sure you want to <strong>deny</strong> this estimate?'
      confirm={deny}
      noteRequired
      noteTitle='Reason'
      tooltip='Deny'
      okText='DENY'
    >
      <IconThumbDown color='red'/>
    </ConfirmAction>
  {:else if EstimateStatus.Approved === item?.status}
    <Chip type='success' content='APPROVED' hideCancel/>
  {:else if EstimateStatus.Denied === item?.status}
    <Chip type='error' content='DENIED' hideCancel={!isAdminOrCoAdmin} on:cancel={confirmRemoveDenied}/>
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
  Are you sure you want to remove DENIED status of this estimate?
</Modal>
