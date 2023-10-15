<script lang="ts">
  import DropDown from "$lib/components/form/DropDown.svelte";

  export let value;
  export let ignore = [];

  const list = [
    {id: 'Request', name: 'New Request'},
    {id: 'Estimating', name: 'Estimating'},
    {id: 'Estimated', name: 'Completed'},
    {id: 'Approved', name: 'Approved'},
    {id: 'Denied', name: 'Denied'},
  ]

  const options = list.filter(o => !ignore.includes(o.id))

</script>

<DropDown
  placeholder='All Statuses'
  {value}
  cls='sm:mt-0 mt-2'
  {options}
  on:change
/>
