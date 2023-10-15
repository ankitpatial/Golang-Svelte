<script lang='ts'>
  import { createEventDispatcher } from 'svelte';
  import { getContextClient } from '@urql/svelte';
  import InputDate from '$lib/components/form/dtp/InputDate.svelte';
  import FileUpload from '$lib/components/FileUpload';
  import type { Document, FileInfo } from '$lib/components/FileUpload/d';
  import { JobDocUploadUrl } from '$lib/gql/document';
  import type { Answer, Question } from './d';
  import { progress } from './store';

  export let question: Question;

  const client = getContextClient();
  const dispatch = createEventDispatcher();
  let files = [];

  function input(opt, data?: any) {
    const ans: Answer = {
      qID: question.id,
      qType: question.type,
      question: question.text,
      answer: opt,
      data
    };
    dispatch('input', ans);
  }

  async function uploadUrl(name: string, f: FileInfo) {
    const res = await client.mutation(
      JobDocUploadUrl,
      {
        jobID: $progress.job.id,
        name: name || '',
        section: question.upload?.section,
        fileName: f.name,
        fileType: f.type,
        fileSize: f.size
      }
    ).toPromise();

    if (res.error) {
      return Promise.reject(res.error.message);
    }

    // data =>  {id, key, name, contentType, contentSize, ready, uploadUrl}
    return res.data.jobDocUploadUrl;
  }

  function uploadComplete(inpName: string, files: Array<Document>) {
    const names = files.map(d => d.name).join('<br/>');
    input(names, files);
  }
</script>

<div class='mt-5'>
  <h3>
    {question.text}
  </h3>

  {#if (question.type === 'buttons')}
    <div class='flex flex-row gap-x-2  mt-2'>
      {#each question.options as opt}
        <div>
          <button type='button' class='btn btn-outline' on:click={() => input(opt)}>
            {opt}
          </button>
        </div>
      {/each}
    </div>
  {:else if (question.type === 'radio')}
    <div class='flex flex-col gap-y-2 mt-5'>
      {#each question.options as opt, idx}
        <div class='form-control'>
          <label class='label cursor-pointer inline flex justify-start gap-x-2 p-0.5'>
            <input type='radio' name='ans' class='radio checked:bg-red-500' value={opt}
                   on:change={(e) => input(e.target.value)} />
            <span class='label-text'>{opt}</span>
          </label>
        </div>
      {/each}
    </div>

  {:else if (question.type === 'dtp')}
    <div class='flex flex-col gap-y-2  mt-5'>
      <InputDate on:change={({detail})=> input(detail)} />
    </div>
  {:else if (question.type === 'note')}
    <div class='flex flex-col gap-y-2'>
      {#each question.options as opt, idx}
        <p>{@html opt}</p>
      {/each}
    </div>
  {:else if (question.type === 'fileUpload')}
    <div class='flex flex-col gap-y-2'>
      {#if question.description}
        <p>{question.description}</p>
      {/if}
      <FileUpload
        h='h-20'
        name={question.upload?.name}
        accept={question.upload?.accept}
        multiple={question.upload?.multiple}
        {uploadUrl}
        {uploadComplete}
      />
    </div>
  {/if}
</div>
