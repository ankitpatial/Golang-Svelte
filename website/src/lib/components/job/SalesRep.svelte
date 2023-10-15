<script>
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';

  export let company = '';
  export let data;
</script>

<CardWithBorderTitle title='Sales Rep'>
  <InputReadonly label='Company Name' value={company || 'N/A'} />
  <InputReadonly label='First Name' value={data?.firstName} />
  <InputReadonly label='Last Name' value={data?.lastName} />
  <InputReadonly label='Email' value={data?.email} />
  <InputReadonly label='Mobile' value={data?.phone} />
</CardWithBorderTitle>
