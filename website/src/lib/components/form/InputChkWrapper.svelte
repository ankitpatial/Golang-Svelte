<script lang='ts'>
  export let label = '';
</script>

<div class='form-control'>
  <label class='label cursor-pointer'>
    <slot></slot>
    <span class='label-text ml-2'>{label}</span>
  </label>
</div>

