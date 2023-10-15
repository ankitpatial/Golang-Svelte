<script lang='ts'>
  import { getContextClient, queryStore } from '@urql/svelte';
  import { JobDocUploadUrl, QryJobDocs } from '$lib/gql/document';
  import PlaceholderBasic from '$lib/components/placeholder/PlaceholderBasic.svelte';
  import Doc from '$lib/components/Doc.svelte';
  import InputFile from '$lib/components/form/InputFile.svelte';
  import { onMount } from 'svelte';
  import { wsMessage } from '$lib/stores/socket';
  import { Channel, type Document, Topic } from '$lib/graph/graphql';

  export let id: string;

  const client = getContextClient();
  let store = load();

  onMount(() => {
    return wsMessage.subscribe((msg) => {
      if (!msg) return;
      if (Channel.Job === msg.channel && Topic.FileUpload === msg.topic && msg.data.id === id) {
        store = load();
      }
    });
  });

  function load() {
    return queryStore<any>({ client, query: QryJobDocs, variables: { jobID: id } });
  }

  async function findUploadInfo(f: File): Promise<Document> {
    // call default function
    const res = await client.mutation(JobDocUploadUrl, {
        jobID: id,
        section: 'Docs',
        name: f.name,
        fileName: f.name,
        fileType: f.type,
        fileSize: f.size
      }
    ).toPromise();

    if (res.error) {
      return Promise.reject(res.error);
    }

    return res.data.jobDocUploadUrl;
  }
</script>

<div class='flex flex-col gap-y-2'>
  {#if $store.fetching}
    <PlaceholderBasic />
  {:else if $store.data?.jobDocs?.length > 0}
    {#each $store.data?.jobDocs as doc}
      <div>
        <Doc data={doc} />
      </div>
    {/each}
  {:else }
    <p>
      no document found
    </p>
  {/if}
  <div class='justify-between'>
    <InputFile folder='JobDocs' section='Docs' dir={id} {findUploadInfo} />
  </div>
</div>

