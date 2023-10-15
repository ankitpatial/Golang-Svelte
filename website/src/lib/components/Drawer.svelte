<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import LogoRoofix from '$lib/assets/svg/LogoRoofix.svelte';
  import {authUser, authUserDashboard} from '$lib/stores/auth.js';
  import {page} from '$app/stores';

  export let cls = 'bg-secondary';
</script>

<div class='drawer lg:drawer-open' class:blur-sm={!$authUser}>
  <!--  toggle trigger  -->
  {#key ($page.route.id)}
    <input id='main-drawer' type='checkbox' class='drawer-toggle op'/>
  {/key}
  <!-- RightSide: Content -->
  <div class='drawer-content flex flex-col pb-12'>
    <!-- Page content here -->
    <slot></slot>
  </div>
  <!-- LeftSide: Menu/Navigation -->
  <div class='drawer-side'>
    <label for='main-drawer' class='drawer-overlay'></label>
    <div class='sm:rounded-md m-0 sm:m-2 pt-0 pb-4 relative left-panel {cls}'>
      <!-- logo:Desktop Large -->
      <div class='overflow-hidden sticky top-2 pt-4  z-10 rounded-tl-md rounded-tr-md text-center'>
        <a href={$authUserDashboard} class='btn btn-ghost text-white '>
						<span>
							<LogoRoofix cls='h-10'/>
						</span>
        </a>
      </div>
      <!-- Left NavBar -->
      <div class='ml-2 pr-2 overflow-y-auto height-auto'>
        <slot name='left'></slot>
      </div>
    </div>
  </div>
</div>
