<script lang='ts'>
  import {goto} from '$app/navigation';
  import Avatar from '$lib/components/Avatar.svelte';
  import IconEmail from '$lib/assets/svg/IconEmail.svelte';
  import IconPhone from '$lib/assets/svg/IconPhone.svelte';
  import Card from '$lib/components/Card.svelte';
  import {PATH} from '$lib/enum';
  import BtnGhost from '$lib/components/BtnGhost.svelte';
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import {page} from '$app/stores';
  import {AccountStatus, type SessionUser} from "$lib/graph/graphql";

  export let user: SessionUser;
  let color;
  $: {
    switch (user?.status) {
      case AccountStatus.Active:
        color = 'badge-success';
        break;
      case AccountStatus.Pending:
        color = 'badge-warning';
        break;
      case AccountStatus.Disabled:
        color = 'badge-slate-200';
        break;
    }
  }

  function handleEdit() {
    goto(PATH.USERS_EDIT.replace(':id', user.id) + $page.url.search);
  }
</script>

<Card>
  <div slot='title'>
    <div class='place-items-center'>
      <Avatar name='{user?.name}' src='{user?.picture}'
              placeHolderClass='w-28 h-28 bg-slate-200 text-gray-700 rounded-none'
              imgClass='w-28 h-28 rounded-md'
              txtClass='text-xl font-bold'
      />
    </div>
  </div>

  <div slot='header'>
    <p class='card-title text-sm leading-4'>{user?.name}</p>
    <p class='inline-flex text-xs text-gray-400 lowercase w-full'>{user?.role}</p>
  </div>

  <svelte:fragment slot='actions'>
    <BtnGhost on:click={handleEdit} tooltip='Edit User'>
      <IconEdit className='h-5'/>
    </BtnGhost>
  </svelte:fragment>

  <div class='flex mb-px'>
    <p>
      <IconEmail className='h-4 w-4 mr-1 inline'/>
      {user?.email}
    </p>
  </div>

  <div class='flex justify-between'>
    <p>
			<span class='mr-1 align-text-bottom'>
				<IconPhone className='h-4 w-4 inline'/>
			</span>
      {user?.phone}
    </p>
    <span class='badge {color} badge-sm mt-1'>
			{user?.status}
		</span>
  </div>

</Card>
