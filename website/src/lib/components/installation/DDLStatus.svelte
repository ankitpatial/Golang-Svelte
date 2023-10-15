<script lang='ts'>
  import DropDown from '$lib/components/form/DropDown.svelte';
  import { page } from '$app/stores';
  import { InstallationStatus } from '$lib/graph/graphql';

  export let value: string | undefined;

  const options = [
    { id: '', name: 'All Statuses' },
    { id: InstallationStatus.New, name: 'New' },
    { id: InstallationStatus.Scheduled, name: 'Scheduled' },
    { id: InstallationStatus.Installed, name: 'Installed' }
  ];

  function progressChange({ detail }) {
    if (detail) {
      value = detail;
      $page.url.searchParams.set('progress', detail);
    } else {
      value = undefined;
      $page.url.searchParams.delete('progress');
    }
  }
</script>

<DropDown
  value={value || ''}
  cls='sm:mt-0 mt-2'
  {options}
  on:change={progressChange}
/>
