<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {PATH} from "$lib/enum";
  import PartnerJobsView from "$lib/components/partner/PartnerJobsView.svelte";

  const title = 'Approved Jobs';
  let flagged = false, progress;

</script>

<PartnerJobsView isSolarPartner hideBackBtn {title} baseURL={PATH.SOLAR_PARTNER_JOBS}/>
