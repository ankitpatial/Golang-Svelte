<script>
  import DropDown from '$lib/components/form/DropDown.svelte';
  export let value = 'THIS_WEEK'
</script>

<DropDown
  bind:value
  options={[
    {id: 'TODAY', name: 'Today'},
    {id: 'YESTERDAY', name: 'Yesterday'},
    {id: 'THIS_WEEK', name: 'This Week'},
    {id: 'LAST_WEEK', name: 'Last Week'},
    {id: 'THIS_MONTH', name: 'This Month'},
    {id: 'LAST_MONTH', name: 'Last Month'},
    {id: 'THIS_YEAR', name: 'This Year'},
    {id: 'LAST_YEAR', name: 'Last Year'},
  ]}
  on:change
/>
