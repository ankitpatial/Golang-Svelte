<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script context="module" lang="ts">
  import {gql} from "@urql/svelte";

  const qry = gql`query ViewDocURL($docID: String!){
        viewDocURL(docID: $docID)
    }`
</script>

<script lang='ts'>
  import IconDocument from '$lib/assets/svg/IconDocument.svelte';
  import {getContextClient} from "@urql/svelte";
  import alerts from "$lib/stores/alerts";
  import Modal from "$lib/components/Modal.svelte";

  export let data: {
    id: string,
    key: string,
    name: string,
    contentType?: string,
    contentSize: string,
    ready: boolean
  } | undefined;

  const client = getContextClient();
  let busy = false, show = false, url;

  async function open() {
    if (busy) {
      return;
    }

    busy = true;
    url = '';

    // get SignedURL
    const res = await client.query(qry, {docID: data?.id}).toPromise()
    busy = false;
    if (res.error) {
      alerts.error('View Document', res.error.message)
      return
    }

    show = true;
    url = res.data.viewDocURL;
    // open modal
  }

</script>

{#if data}
  <div class='btn btn-ghost' on:click={open} class:loading={busy}>
    <IconDocument/>
    <div>
      {data.name}
    </div>
  </div>
{/if}

{#if data && show}
  <Modal
    open
    size="lg"
    title="File: {data.name}"
    hideOkBtn
    cancelText="CLOSE"
    on:cancel={() => {show = false;}}
  >
    <iframe src={url} class="w-full h-[80vh]">
    </iframe>
  </Modal>
{/if}
