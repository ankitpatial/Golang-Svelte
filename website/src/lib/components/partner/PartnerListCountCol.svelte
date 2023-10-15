<script lang='ts'>
  import {goto} from "$app/navigation";
  import {page} from "$app/stores";
  import {JobProgress} from "$lib/graph/graphql";

  export let id: string;
  export let name: string;
  export let count: number;
  export let countFlagged: number;
  export let status: JobProgress | undefined = undefined;

  let hasFlagged = (countFlagged || 0) > 0

  function go() {
    if (!hasFlagged) return;

    $page.url.searchParams.set("name", name);
    $page.url.searchParams.set("flagged", "1");
    $page.url.searchParams.set("progress", status || '');
    goto(`${$page.url.pathname}/${id}${$page.url.search}`);
  }
</script>
<div class='flex flex-row space-x-2'>
  <div>
    {count || 0}
  </div>
  <div class='tooltip-top' data-tip='Click to view flagged jobs' class:tooltip={hasFlagged} on:click={go}>
    <div
      class='badge badge-error'
      class:link={hasFlagged}
      class:link-primary={hasFlagged}
    >
      {countFlagged || 0}
    </div>
  </div>
</div>
