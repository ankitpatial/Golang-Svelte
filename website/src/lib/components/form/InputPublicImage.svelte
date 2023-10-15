<script lang='ts'>
  import Cropper from '$lib/components/Cropper/Cropper.svelte';
  import { PublicDataUploadUrl } from '$lib/gql/document';
  import { uploadFile } from '$lib/components/FileUpload/uploader';
  import { getContextClient } from '@urql/svelte';
  import { createEventDispatcher } from 'svelte';

  export let label = '';
  export let uploadBtnLabel = 'Upload';
  export let docID: string;
  export let docDir: string;
  export let docName: string;
  export let docType = 'image/jpeg';
  export let docSection: string;
  export let docURL: string;
  export let aspectRatio = NaN;

  const client = getContextClient();
  const dispatch = createEventDispatcher();

  let err, busy = false, fileUploadStatus, isErr;

  $:isTip = !!err;
  $:isTipErr = !!err;
  $:tip = err;

  async function imgCrop({ detail }) {
    busy = true;
    isErr = false;
    docURL = detail.url;
    const f = new File([detail.blob], docName);

    // get upload url data
    const res = await client.mutation(PublicDataUploadUrl, {
      dir: docDir,
      name: docName,
      section: docSection,
      fileName: docName,
      fileType: docType,
      fileSize: f.size
    }).toPromise();
    if (res.error) {
      busy = false;
      return Promise.reject(res.error.message);
    }

    // upload file to storage
    const d = res.data.publicDataUploadUrl;
    await uploadFile(1, f, docName, d.uploadUrl, d.meta, (_, status: string, err: undefined | string | Error) => {
      switch (status) {
        case 'completed':
          docURL = d.publicUrl;
          dispatch('complete', { id: d.id, name: docName, url: docURL });
          // save changes to DB
          fileUploadStatus = '';
          busy = false;
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
  <label class='label p-0 mb-0.5 uppercase'>
    <span>{label}</span>
    {#if $$slots.alt}
      <slot name='alt'></slot>
    {/if}
  </label>

  <div class='tooltip-top' class:tooltip={isTip} class:tooltip-error={isTipErr} data-tip={tip}>
    {#if docURL}
      <img src={docURL} alt={label} />
    {/if}
    <Cropper
      label={uploadBtnLabel}
      {busy}
      {aspectRatio}
      on:crop={imgCrop}
      on:error={({detail}) => console.log('err', detail) }
    />
    {#if fileUploadStatus}
      <label class='badge clear-both ml-5' class:badge-info={!isErr} class:badge-error={isErr}>
        {fileUploadStatus}
      </label>
    {/if}
  </div>
</div>

