<script lang='ts'>
  // cropper
  import CropperCanvas from '@cropper/element-canvas';
  import CropperImage from '@cropper/element-image';
  import CropperHandle from '@cropper/element-handle';
  import CropperSelection from '@cropper/element-selection';
  import CropperGrid from '@cropper/element-grid';
  import CropperCrosshair from '@cropper/element-crosshair';
  // ~~
  import BtnSmall from '$lib/components/form/BtnSmall.svelte';
  import IconUpload from '$lib/assets/svg/IconUpload.svelte';
  import { getImgDimensions } from '$lib/utils/image';
  import { createEventDispatcher } from 'svelte';

  export let label = 'Browse Image';
  export let height: number = 200;
  export let width: number = 400;
  export let busy = false;
  export let ghost = false;
  export let imageType = 'image/jpg';
  export let aspectRatio: number = NaN;

  const dispatch = createEventDispatcher();

  let inp, imgSrc, cropper, show, imgD;

  // cropper elements
  if (!customElements.get('cropper-canvas')) {
    customElements.define('cropper-canvas', CropperCanvas);
    customElements.define('cropper-image', CropperImage);
    customElements.define('cropper-handle', CropperHandle);
    customElements.define('cropper-selection', CropperSelection);
    customElements.define('cropper-grid', CropperGrid);
    customElements.define('cropper-crosshair', CropperCrosshair);
  }

  function browse() {
    if (busy) return;
    // open file dialog
    inp.click();
  }

  async function handleFile(e) {
    if (!e.target.files || e.target.files.length == 0) return;

    const f = e.target.files[0];
    imgD = await getImgDimensions(f);
    let reader = new FileReader();
    reader.onloadend = function() {
      imgSrc = reader.result;
    };
    reader.readAsDataURL(f);

    show = true;
    try {

      await cropper.$reset();
      await cropper.$change(10, 10, Number(width), Number(height));
      await cropper.$center();
    } catch (ex) {
      dispatch('error', ex.toString());
    }
  }

  async function crop() {
    const c = await cropper.$toCanvas();
    c.toBlob((blob) => {
      dispatch('crop', { url: URL.createObjectURL(blob), blob });
    }, imageType, 1.0);
    show = false;
  }
</script>

<div class='mt-2'>
  <input bind:this={inp} type='file' accept='image/*' on:input={handleFile} class='hidden' />
  <BtnSmall
    ghost={ghost}
    outline={!ghost}
    on:click={browse}
    {busy}
  >
    <IconUpload />
    {label}
  </BtnSmall>
</div>

<div class='modal modal-middle h-full' class:modal-open={show}>
  <div class='modal-box h-[90vh] w-11/12 max-w-5xl overflow-hidden'>
    <cropper-canvas background class='h-[80vh]'>
      <cropper-image src={imgSrc} alt='Picture'></cropper-image>
      <cropper-shade hidden></cropper-shade>
      <cropper-handle action='select' plain></cropper-handle>
      <cropper-selection
        initial-coverage='0.5'
        movable
        resizable
        zoomable
        outlined
        aspect-ratio={aspectRatio}
        bind:this={cropper}
      >
        <cropper-grid role='grid' bordered covered></cropper-grid>
        <cropper-crosshair theme-color='rgba(238, 238, 238, 0.5)' centered></cropper-crosshair>
        <cropper-handle action='move' theme-color='rgba(255, 255, 255, 0.35)'></cropper-handle>
        <cropper-handle action='n-resize'></cropper-handle>
        <cropper-handle action='e-resize'></cropper-handle>
        <cropper-handle action='s-resize'></cropper-handle>
        <cropper-handle action='w-resize'></cropper-handle>
        <cropper-handle action='ne-resize'></cropper-handle>
        <cropper-handle action='nw-resize'></cropper-handle>
        <cropper-handle action='se-resize'></cropper-handle>
        <cropper-handle action='sw-resize'></cropper-handle>
      </cropper-selection>
    </cropper-canvas>
    <div class='modal-action '>
      <BtnSmall primary on:click={crop}>
        CROP
      </BtnSmall>
      <BtnSmall on:click={()=> {show=false}}>
        CANCEL
      </BtnSmall>
    </div>
  </div>
</div>
