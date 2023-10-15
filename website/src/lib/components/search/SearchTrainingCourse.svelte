<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import SearchBox from '$lib/components/form/SearchBox.svelte';
  import alerts from '$lib/stores/alerts';
  import type Entity from '$lib/models/Entity';
  import {CreateTrainingCourse, QryTrainingCourses} from "$lib/gql/training";
  import BtnSmall from "$lib/components/form/BtnSmall.svelte";
  import Modal from "$lib/components/Modal.svelte";
  import Input from "$lib/components/form/Input.svelte";

  export let value: Entity;

  const title = 'Training Course';
  const client = getContextClient();

  let showModal = false, name = '';

  async function search(q: string) {
    const res = await client.query(QryTrainingCourses, {search: q, page: {first: 10}}).toPromise();
    if (res.error) {
      alerts.push({type: 'error', title: title, body: res.error.message});
      return [];
    }

    return res.data.trainingCourses.edges.map((e) => ({id: e.node.id, name: e.node.name}));
  }

  async function createCourse() {
    if (!name) {
      alerts.warning(title, 'Course name is required');
      return
    }
    const res = await client.mutation(CreateTrainingCourse, {name}).toPromise();
    if (res.error) {
      alerts.error(title, res.error.message);
      return;
    }

    value = res.data.createTrainingCourse
    showModal = false;
  }
</script>

<SearchBox
  label={title}
  placeholder='search...'
  searchHandler={search}
  bind:value
  on:change
>
  <svelte:fragment slot="create">
    <BtnSmall on:click={() => {showModal = true}}>
      Create New
    </BtnSmall>
  </svelte:fragment>
</SearchBox>

<Modal
  open={showModal}
  on:ok={createCourse}
  on:cancel={() => {showModal = false; name= '';}}
>
  <Input label="Course Name" bind:value={name} required/>
</Modal>
