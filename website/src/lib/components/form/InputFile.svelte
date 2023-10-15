<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import { createEventDispatcher } from 'svelte';
  import IconUpload from '$lib/assets/svg/IconUpload.svelte';
  import { DocumentUploadUrl } from '$lib/gql/document';
  import { uploadFile } from '$lib/components/FileUpload/uploader';
  import IconDocument from '$lib/assets/svg/IconDocument.svelte';
  import Label from '$lib/components/form/Label.svelte';
  import type { Document } from '$lib/graph/graphql';

  export let label = '';
  export let name = undefined; // name on storage
  export let fileName = undefined; // display name
  export let uploadBtnLabel = 'Upload';
  export let folder: string;
  export let dir: string;
  export let section: string;
  export let overwrite = false;
  export let accept = '*/*';
  export let ghost = false;

  export let findUploadInfo: (f: File) => Promise<Document> = null;

  const client = getContextClient();
  const dispatch = createEventDispatcher();

  let id = crypto.randomUUID();
  let inp;
  let busy = false;
  let fileUploadStatus, isErr;

  function browse() {
    if (busy) return;
    // open file dialog
    inp.click();
    fileUploadStatus = '';
    isErr = false;
  }

  async function uploadInfo(f: File): Promise<Document> {
    // call passed on function
    if (findUploadInfo && typeof findUploadInfo === 'function') {
      return findUploadInfo(f);
    }

    // call default function
    const res = await client.mutation(DocumentUploadUrl, {
        folder,
        dir,
        section,
        name: name || f.name,
        fileName: f.name,
        fileType: f.type,
        fileSize: f.size,
        overwrite
      }
    ).toPromise();

    if (res.error) {
      busy = false;
      isErr = true;
      fileUploadStatus = res.error.message;
      return;
    }

    return res.data.documentUploadUrl;
  }

  async function handleFile(e) {
    if (!e.target.files || e.target.files.length == 0) return;

    busy = true;
    // file
    const f = e.target.files[0];
    // get upload URL
    const d = await uploadInfo(f);
    // check if empty upload url
    if (!d.uploadUrl) {
      busy = false;
      isErr = true;
      fileUploadStatus = 'Error: Empty upload URL';
      return;
    }

    await uploadFile(1, f, f.name, d.uploadUrl, d.meta, (_, status: string, err: undefined | string | Error) => {
      switch (status) {
        case 'completed':
          fileUploadStatus = '';
          busy = false;
          dispatch('complete', { id: d.id, name: f.name });
          break;
        case 'error':
          busy = false;
          isErr = true;
          fileUploadStatus = 'Error: ' + err;
          break;
        default:
          fileUploadStatus = 'Status: ' + status;
      }
    });
  }
</script>

<div class='form-control '>
  <Label content={label} {id} />

  {#if fileName}
    <div class='flex flex-row mt-2'>
      <div>
        <IconDocument />
      </div>
      <div class='font-bold'>
        {fileName}
      </div>
    </div>
  {/if}
  <div class='mt-2'>
    <input {id} bind:this={inp} type='file' {accept} on:input={handleFile} class='hidden' />
    <div
      class='btn btn-sm'
      class:btn-ghost={ghost}
      class:btn-outline={!ghost}
      class:loading={busy}
      on:click={browse}
    >
      <IconUpload />
      {uploadBtnLabel}
    </div>
  </div>
  {#if fileUploadStatus}
    <label class='badge clear-both ml-5' class:badge-info={!isErr} class:badge-error={isErr}>
      {fileUploadStatus}
    </label>
  {/if}
</div>
