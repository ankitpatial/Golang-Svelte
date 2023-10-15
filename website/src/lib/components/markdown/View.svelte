<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script>
  import {marked} from 'marked';
  import './style.css';

  export let content;

</script>

<div class='markdown-body'>
  {@html marked(content)}
</div>

<style>
  @media (max-width: 767px) {
    .markdown-body {
      padding: 0;
    }
  }
</style>

