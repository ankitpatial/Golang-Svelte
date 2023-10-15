<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import TextArea from '$lib/components/form/TextArea.svelte';
  import SearchTrainingCourse from '$lib/components/search/SearchTrainingCourse.svelte';
  import InputPublicImage from '$lib/components/form/InputPublicImage.svelte';
  import { page } from '$app/stores';
  import InputFile from '$lib/components/form/InputFile.svelte';
  import Btn from '$lib/components/form/Btn.svelte';
  import { SaveTrainingVideo } from '$lib/gql/training';
  import alerts from '$lib/stores/alerts';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();

  let kind, id, name, title, busy = false;
  $: {
    kind = $page.params.kind;
    id = $page.params.id;
    name = $page.url.searchParams.get('kn');
    title = name + ' Video Upload';
  }

  let data = {
    course: undefined,// {id, name}
    title: '',
    description: '',
    poster: undefined,// {id, name, url}
    video: {} // {id, name}
  };

  function back() {
    goto(PATH.TRAINING_VIDEOS.replace(':kind', kind) + $page.url.search);
  }

  async function save() {
    if (busy) return;

    const title = 'Training Video';
    busy = true;
    const res = await client.mutation(SaveTrainingVideo, {
      inp: {
        id,
        kind: kind.toUpperCase(),
        title: data.title,
        description: data.description,
        courseID: data.course?.id,
        posterID: data.poster?.id,
        videoID: data.video?.id
      }
    }).toPromise();
    busy = false;
    if (res.error) {
      alerts.error(title, res.error.message);
      return;
    }

    await goto(PATH.TRAINING_VIDEOS.replace(':kind', kind) + `?kn=${name}`);
    alerts.success(title, 'Saved successfully');
  }
</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall class='btn btn-outline btn-sm uppercase' on:click={back} />
  </svelte:fragment>

  <form on:submit|preventDefault={save}>
    <div class='md:w-6/12 sm:w-full flex flex-col'>
      <SearchTrainingCourse bind:value={data.course} />

      <Input label='Title' bind:value={data.title} required />

      <TextArea label='Description' bind:value={data.description} required />

      <InputPublicImage
        label='Video Poster'
        uploadBtnLabel='Upload Poster'
        docDir={id}
        docName='poster.jpeg'
        docSection='Image'
        aspectRatio={16/9}
        on:complete={({detail}) => data.poster = detail}
      />

      <InputFile
        label='Video'
        uploadBtnLabel='Upload Video'
        folder='TrainingVideos'
        dir={id}
        section='Video'
        name='video'
        accept='video/mp4,video/webm'
        fileName={data.video?.name}
        on:complete={({detail}) => data.video = detail}
        overwrite
      />
    </div>

    <Btn primary type='submit' cls='mt-5'>
      SAVE
    </Btn>
  </form>
</Page>
