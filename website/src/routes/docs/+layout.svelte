<script lang='ts'>
  import Header from '$lib/components/Header.svelte';
  import Drawer from '$lib/components/Drawer.svelte';
  import LeftNav from './LeftNav.svelte';
</script>

<Drawer cls='bg-black'>
  <Header/>
  <div class='sm:px-6 px-4'>
    <slot></slot>
  </div>

  <svelte:fragment slot='left'>
    <LeftNav/>
  </svelte:fragment>
</Drawer>
