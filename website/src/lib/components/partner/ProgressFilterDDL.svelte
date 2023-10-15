<script lang="ts">
  import DropDown from "$lib/components/form/DropDown.svelte";
  import {page} from "$app/stores";

  export let value: string | undefined;
  export let skipPaymentOptions: boolean = false;

  const options = [
    {id: '', name: 'All Statuses'},
    {id: 'New', name: 'New'},
    {id: 'Delayed', name: 'Delayed'},
    {id: 'CustomerContacted', name: 'Customer Contacted'},
    {id: 'JobDetailsConfirmed', name: 'Job Details Confirmed'},
    {id: 'Permitting', name: 'Permitting'},
    {id: 'Scheduled', name: 'Scheduled'},
    {id: 'InProgress', name: 'In Progress'},
    {id: 'Installed', name: 'Installed'},
  ]

  if (skipPaymentOptions) {
    options.push({id: 'Invoiced', name: 'Invoiced'})
  } else {
    options.push({id: 'Invoiced', name: 'Payment Pending'})
    options.push({id: 'InvoiceApproved', name: 'Payment Approved'})
    options.push({id: 'InvoicePaid', name: 'Paid'})
  }

  function progressChange({detail}) {
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
