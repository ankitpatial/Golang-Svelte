<script lang='ts'>
  import DropDown from '$lib/components/form/DropDown.svelte';
  import { page } from '$app/stores';
  import { Approval } from '$lib/graph/graphql';

  export let value: string | undefined;

  const options = [
    { id: '', name: 'All Statuses' },
    { id: Approval.Pending, name: 'Pending' },
    { id: Approval.Approved, name: 'Approved' },
    { id: Approval.Denied, name: 'Denied' }
  ];

  function handleChange({ detail }) {
    if (detail) {
      value = detail;
      $page.url.searchParams.set('status', detail);
    } else {
      value = undefined;
      $page.url.searchParams.delete('status');
    }
  }
</script>

<DropDown
  value={value || ''}
  cls='sm:mt-0 mt-2'
  {options}
  on:change={handleChange}
/>
