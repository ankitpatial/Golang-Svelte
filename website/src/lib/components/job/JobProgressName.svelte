<script lang='ts'>
  import {humanize} from '$lib/utils/string';
  import {localDate} from '$lib/utils/time';
  import {JobProgress} from "$lib/graph/graphql";

  export let progress: JobProgress;
  export let progressAt: Date;

</script>
<div
  class='tooltip tooltip-right badge badge-ghost'
  data-tip='Dated: {localDate(progressAt, false)}'
>
  {#if !progress }
    N/A
  {:else if progress === JobProgress.Invoiced}
    Payment Pending
  {:else if progress === JobProgress.InvoiceApproved}
    Payment Approved
  {:else if progress === JobProgress.InvoicePaid}
    Paid
  {:else}
    {humanize(progress)}
  {/if}
</div>
