<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {PATH} from "$lib/enum";
  import PartnerJobsView from "$lib/components/partner/PartnerJobsView.svelte";

  const title = 'Jobs Dashboard';
</script>


<PartnerJobsView hideBackBtn {title} baseURL={PATH.ROOFING_PARTNER_JOBS}/>

