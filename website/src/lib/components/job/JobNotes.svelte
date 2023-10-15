<script lang="ts">
  import Modal from "$lib/components/Modal.svelte";
  import PlaceholderBasic from "$lib/components/placeholder/PlaceholderBasic.svelte";
  import {getContextClient, queryStore} from "@urql/svelte";
  import {onMount} from "svelte";
  import {wsMessage} from "$lib/stores/socket";
  import type {
    CreatJobNoteMutation,
    CreatJobNoteMutationVariables,
    EditJobNoteMutation,
    EditJobNoteMutationVariables,
    JobNotesQueryStore
  } from "$lib/graph/graphql";
  import {Channel} from "$lib/graph/graphql";
  import {authUser} from "$lib/stores/auth";
  import Btn from "$lib/components/form/Btn.svelte";
  import IconPlusSmall from "$lib/assets/svg/IconPlusSmall.svelte";
  import TextArea from "$lib/components/form/TextArea.svelte";
  import alerts, {modalAlert} from "$lib/stores/alerts";
  import {CreatJobNote, EditJobNote, QryJobNotes} from "$lib/gql/job.notes";
  import Avatar from "$lib/components/Avatar.svelte";
  import {localDate} from "$lib/utils/time";

  export let id;
  export let assignedPartnerID;

  const client = getContextClient();
  let open = false
  let busy = false
  let store = load();
  let noteID;
  let note;
  let allowed: boolean
  $:allowed = ($authUser?.partner && $authUser.partner.id === assignedPartnerID) || false

  onMount(() => {
    return wsMessage.subscribe((msg) => {
      if (!msg) return;
      if (Channel.JobNote === msg.channel && msg.data.id === id) {
        store = load();
      }
    })
  })

  function load() {
    return queryStore<JobNotesQueryStore>({client, query: QryJobNotes, variables: {jobID: id}})
  }

  function toggleModal() {
    open = !open;
    noteID = '';
    note = '';
  }

  async function save() {
    if (busy || !note) return;
    if (note.length > 500) {
      modalAlert.warning('Save Note', 'Message limit exceeded');
      return;
    }

    let res;
    if (noteID) {
      res = await client.mutation<EditJobNoteMutation, EditJobNoteMutationVariables>(EditJobNote, {
        jobID: id,
        noteID,
        note,
      }).toPromise();
    } else {
      res = await client.mutation<CreatJobNoteMutation, CreatJobNoteMutationVariables>(CreatJobNote, {
        jobID: id,
        note,
      }).toPromise();
    }

    if (res.error) {
      modalAlert.error('Save Note', res.error.message);
      return;
    }

    toggleModal();
    alerts.success('Save Note', 'Successfully saved');
  }

  async function edit(id, n) {
    if (!allowed) return;

    toggleModal();
    noteID = id;
    note = n;
  }
</script>

<div>
  {#if $store.fetching}
    <PlaceholderBasic/>
  {:else if $store.data?.jobNotes.length > 0}
    <div class="flex flex-col space-y-2 justify-items-start">
      {#each $store.data?.jobNotes as n}
        <div class="flex flex-row flex-wrap justify-start space-x-2 border-b p-1">
          <div>
            <Avatar name="{n.creator.firstName} {n.creator.lastName}" src={n.creator.picture}/>
          </div>
          <div class="-mt-1" class:cursor-pointer={allowed} on:click={() => edit(n.id, n.note)}>
            <span class="text-xs">{localDate(n.createdAt)}</span>
            <p class="whitespace-pre text-lg">
              {n.note}
            </p>
          </div>
        </div>
      {/each}
    </div>
  {:else }
    <p>
      no note found
    </p>
  {/if}
  {#if (allowed)}
    <Btn small on:click={toggleModal} cls="mt-5">
      <IconPlusSmall/>
      Add New Note
    </Btn>
  {/if}
</div>

<Modal title="Add New Note" okText="SAVE" {open} {busy} on:cancel={toggleModal} on:ok={save}>
  <form on:submit|preventDefault={save}>
    <TextArea bind:value={note}/>
  </form>
</Modal>
