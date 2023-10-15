<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script>
  import Logout from '$lib/components/Logout.svelte';
  import Page from '$lib/components/Page.svelte';
</script>

<Page title='Thank You'>
  <p class='mb-5'>
    We are reviewing your application and will soon contact you.
  </p>
  <Logout label='You May Logout Now' cls='btn btn-outline' />
</Page>
