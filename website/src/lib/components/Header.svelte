<script lang='ts'>
  import { slide } from 'svelte/transition';
  import { quintOut } from 'svelte/easing';
  import { authUser, authUserDashboard } from '$lib/stores/auth';
  import LogoRoofix from '$lib/assets/svg/LogoRoofix.svelte';
  import IconBar from '$lib/assets/svg/IconBar.svelte';
  import IconArrowDown from '$lib/assets/svg/IconArrowDown.svelte';
  import Avatar from './Avatar.svelte';
  import { humanize } from '$lib/utils/string';
  import Logout from '$lib/components/Logout.svelte';
  import IconUser from '$lib/assets/svg/IconUser.svelte';
  import { PATH } from '$lib/enum';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconBell from '$lib/assets/svg/IconBell.svelte';
  import { getContextClient } from '@urql/svelte';
  import { QryMyUnreadNotificationsCount } from '$lib/gql/notification.js';
  import type { MyUnreadNotificationsCountQuery } from '$lib/graph/graphql';
  import { onMount } from 'svelte';
  import { wsMessage } from '$lib/stores/socket';
  import { goto } from '$app/navigation';
  import { unreadNotificationsCount } from '$lib/stores/app';

  const client = getContextClient();

  let user;
  let name;
  $: {
    user = $authUser;
    name = `${user?.firstName} ${user?.lastName}`.trim();
  }

  onMount(() => {
    unreadCount();
    return wsMessage.subscribe(msg => {
      if (!msg) return;
      if (msg.data.isNotification) {
        unreadCount();
      }
    });
  });

  function unreadCount() {
    client
      .query <MyUnreadNotificationsCountQuery>(QryMyUnreadNotificationsCount, {})
      .toPromise()
      .then(res => {
        if (res.error) {
          console.error(res.error.message);
          return;
        }
        unreadNotificationsCount.set(res.data?.myUnreadNotificationsCount || 0);
      });
  }

</script>

<div class='flex justify-between p-1 bg-blend-darken sm:border-none border-b mb-6 sticky z-50 top-0 bg-white'>
  <!-- mobile menu -->
  <div>
    <label for='main-drawer' class='drawer-button btn btn-ghost btn-square hover:text-white sm:hidden'>
      <IconBar />
    </label>
  </div>

  <!-- mobile logo -->
  <div class='sm:hidden'>
    <a href={$authUserDashboard} class='btn btn-ghost text-black'>
      <LogoRoofix cls='h-10' />
    </a>
  </div>

  <!-- user section -->
  {#if (user)}
    <div class="flex place-items-center gap-x-3">
      <div class='indicator'>
        {#if ($unreadNotificationsCount > 0)}
        <span
          class='indicator-item indicator-start badge badge-error text-sm'
          transition:slide="{{delay: 250, duration: 300, easing: quintOut, axis: 'x'}}"
        >
          {$unreadNotificationsCount}
        </span>
        {/if}
        <BtnIcon on:click={() => goto(PATH.NOTIFICATIONS)}>
          <IconBell />
        </BtnIcon>
      </div>
      <div class='dropdown dropdown-end'>
        <div tabindex='0'
             class='btn btn-ghost flex group gap-3 pl-0 sm:pr-4 pr-0 rounded-full hover:bg-slate-100'>
          <div class='btn btn-circle btn-ghost'>
            <Avatar name={name} src={user.picture} />
          </div>
          <div class='text-left hidden sm:inline-block'>
            <h4>{name}</h4>
            <small class='text-neutral capitalize'>{humanize(user.role)}</small>
          </div>
          <div class='hidden sm:inline-block'>
				<span>
					<IconArrowDown className='h-3 w-3' />
				</span>
          </div>
        </div>
        <ul tabindex='0' class='dropdown-content menu p-2 shadow-lg bg-base-100 rounded-box w-52'>

          {#if (user?.isAdmin)}
            <li>
              <a href={PATH.PROFILE}>
                <IconUser />
                Profile
              </a>
            </li>
          {/if}
          <li>
            <Logout />
          </li>
        </ul>
      </div>
    </div>
  {/if}
</div>
