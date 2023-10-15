<script lang='ts'>
  import {page} from '$app/stores';
  import {goto} from '$app/navigation';
  import {PartnerType} from '$lib/graph/graphql';
  import IconVideoCamera from '$lib/assets/svg/IconVideoCamera.svelte';
  import IconDocumentMagnifying from '$lib/assets/svg/IconDocumentMagnifying.svelte';
  import IconUsers from '$lib/assets/svg/IconUsers.svelte';
  import IconEyeOutline from '$lib/assets/svg/IconEyeOutline.svelte';
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import IconEllipsisHorizontal from "../../assets/svg/IconEllipsisHorizontal.svelte";

  export let basePath;
  export let partnerType: PartnerType;
  export let id;
  export let name;

  const q = $page.url.searchParams.get('q');

  function view() {
    goto(`${basePath}/${id}${q ? `?q=${q}&` : '?'}name=${name}`);
  }

  function profile() {
    goto(`${basePath}/${id}/profile${q ? `?q=${q}&` : '?'}name=${name}`);
  }

  function gotoManageUsers() {
    goto(`${basePath}/${id}/users${q ? `?q=${q}&` : '?'}name=${name}`);
  }

  function trainingVideos() {
    goto(`${basePath}/${id}/training${q ? `?q=${q}&` : '?'}name=${name}`);
  }

  function mobileSettings() {
    goto(`${basePath}/${id}/mobile-app-settings${q ? `?q=${q}&` : '?'}name=${name}`);
  }
</script>

<div class='dropdown dropdown-left whitespace-normal z-20'>
  <label tabindex='0' class='btn btn-sm btn-ghost'>
    <IconEllipsisHorizontal/>
  </label>
  <ul tabindex='0' class='dropdown-content menu shadow bg-base-100 rounded-box sm:w-38 w-52 '>
    <li on:click={view}>
      <a>
        <IconEyeOutline/>
        View Jobs
      </a>
    </li>
    {#if (PartnerType.Roofing === partnerType || PartnerType.Solar === partnerType)}
      <li on:click={profile}>
        <a>
          <IconEdit/>
          View Profile
        </a>
      </li>
    {/if}
    <li on:click={trainingVideos}>
      <a>
        <IconVideoCamera/>
        Training Videos
      </a>
    </li>
    <li on:click={gotoManageUsers}>
      <a>
        <IconUsers/>
        Manage Users
      </a>
    </li>
    {#if (PartnerType.Solar === partnerType)}
      <li on:click={mobileSettings}>
        <a>
          <IconDocumentMagnifying/>
          Mobile App Settings
        </a>
      </li>
    {/if}
  </ul>
</div>
