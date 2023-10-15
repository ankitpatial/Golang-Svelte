<script>
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';

  export let data;
</script>

<CardWithBorderTitle title='Home Owner'>
  <InputReadonly label='First Name' value={data?.firstName} />
  <InputReadonly label='Last Name' value={data?.lastName} />
  <InputReadonly label='Street Number' value={data?.streetNumber} />
  <InputReadonly label='Street Name' value={data?.streetName} />
  <InputReadonly label='City' value={data?.city} />
  <InputReadonly label='State' value={data?.state} />
  <InputReadonly label='Zip' value={data?.zip} />
  <slot></slot>
</CardWithBorderTitle>
