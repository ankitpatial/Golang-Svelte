<script lang="ts">

  import {routeMatch} from "$lib/utils/routeMatch.js";
  import {page} from "$app/stores";
  import IconLinkBlank from "$lib/assets/svg/IconLinkBlank.svelte";

  export let item: { href: string, label: string; icon: any, target?: string, startsWith?: boolean }
  export let cls = ''

</script>

{#if (item.target === '_blank')}
  <a
    href={item.href}
    class={cls}
    class:active={routeMatch(item.href, $page.url.pathname)}
    target="_blank"
    rel="noreferrer noopener"
  >
    <svelte:component this={item.icon}/>
    <span> {item.label} </span>
    <IconLinkBlank/>
  </a>
{:else }
  <a href={item.href} class={cls} class:active={routeMatch(item.href, $page.url.pathname, item.startsWith)}>
    <svelte:component this={item.icon}/>
    <span> {item.label} </span>
  </a>
{/if}

<style>
  .link:hover {
    @apply bg-secondary-focus text-white !important ;
  }
</style>
