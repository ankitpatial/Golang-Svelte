<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {PATH} from '$lib/enum';
  import {authUser} from '$lib/stores/auth';
  import {routeMatch} from '$lib/utils/routeMatch.js';
  import {page} from '$app/stores';
  import IconDocumentText from '$lib/assets/svg/IconDocumentText.svelte';
  import {Role} from "$lib/graph/graphql";

  const menuItems = [
    {href: PATH.DOCS, label: 'Api Documentation', icon: IconDocumentText}
  ];

  $:isAdminOrCoAdmin = $authUser?.role === Role.Admin;
  $:routeId = $page.route.id;

  let cls = 'rounded-md border-0 hover:bg-slate-800 focus:text-slate-300 focus:ease-in focus:duration-100 focus:scale-90';

</script>

<ul class='menu text-base-content text-slate-300 mt-4 overflow-hidden'>
  {#each menuItems as item}
    <li>
      <a href={item.href} class={cls} class:active={routeMatch(item.href, routeId )}>
        <svelte:component this={item.icon}/>
        <span> {item.label} </span>
      </a>
    </li>
  {/each}
</ul>


<style>
  .menu li a.active {
    @apply text-black bg-slate-300;
  }

  .menu li a.active:hover {
    @apply bg-slate-300;
  }
</style>
