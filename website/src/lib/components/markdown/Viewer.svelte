<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script>
  import View from "./View.svelte";
  import BackToTop from './BackToTop.svelte';

  export let url;
  let fetchImage = Promise.resolve();

  $: if (url) {
    fetchImage = (async () => {
      const response = await fetch(url);
      return await response.text();
    })();
  }

</script>

{#await fetchImage }
  <p>Loading...</p>
{:then data}
  <View content={data}/>
  <BackToTop/>
{:catch error}
  <p>An error occurred!</p>
{/await}

