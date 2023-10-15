<script lang='ts'>
  import { onMount } from 'svelte';
  import { getContextClient } from '@urql/svelte';
  import alerts from '$lib/stores/alerts';
  import { QryPartnerMobileSettings, SavePartnerMobileSettings } from '$lib/gql/partner.settings';
  import { PublicDataUploadUrl } from '$lib/gql/document';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import InputToggle from '$lib/components/form/InputToggle.svelte';
  import Cropper from '$lib/components/Cropper';
  import InputRadio from '$lib/components/form/InputRadio.svelte';
  import Label from '$lib/components/form/Label.svelte';
  import { uploadFile } from '$lib/components/FileUpload/uploader';

  // exported props
  export let partnerID: string;

  // constants
  const client = getContextClient();
  const logImgType = 'image/png';
  const msgTitle = 'Save Mobile App Settings';
  const tabNames = [
      { id: 'SITE_SURVEYS', name: 'SITE SURVEYS' },
      { id: 'HVAC', name: 'HVAC' },
      { id: 'WINDOWS', name: 'WINDOWS' },
      { id: 'INSULATION', name: 'INSULATION' },
      { id: 'SMART_HOME', name: 'SMART HOME' },
      { id: 'TRAINING_CENTER', name: 'TRAINING CENTER' }
    ],
    colors = [
      { id: '#6E88E3', name: 'Purple' },
      { id: '#18265C', name: 'Navy Blue' },
      { id: '#5AB2B8', name: 'Teal' }
    ];

  // variables
  let busy = false,
    logoURL = '',
    logoUploadStatus = '',
    logoW = 400, logoH = 200,
    tOut,
    primaryColor = '',
    hideTabs = [];

  // load data on mount
  onMount(async () => {
    busy = true;
    const res = await client.query(QryPartnerMobileSettings, { id: partnerID }).toPromise();
    busy = false;
    if (res.error) {
      alerts.error('Get Mobile App Settings', res.error.message);
      return;
    }

    const mas = res.data.partnerMobileSettings;
    logoURL = mas.logoURL;
    primaryColor = mas.primaryColor;
    hideTabs = mas.hideTabs;
  });

  // logo image change
  async function imgCrop({ detail }) {
    const name = 'logo.png';
    const f = new File([detail.blob], name, { type: logImgType });
    logoURL = detail.url;
    // get upload url data
    const res = await client.mutation(PublicDataUploadUrl, {
      dir: partnerID,
      name,
      section: 'Logo',
      fileName: name,
      fileType: logImgType,
      fileSize: f.size
    }).toPromise();
    if (res.error) {
      return Promise.reject(res.error.message);
    }

    // upload file to storage
    const d = res.data.publicDataUploadUrl;
    await uploadFile(1, f, name, d.uploadUrl, d.meta, (_, status: string, err: undefined | string | Error) => {
      switch (status) {
        case 'completed':
          // save changes to DB
          save({ logoURL: d.publicUrl }, 'Logo');
          logoUploadStatus = '';
          break;
        case 'error':
          logoUploadStatus = 'Error: ' + err;
          break;
        default:
          logoUploadStatus = 'Upload Status: ' + status;
      }
    });
  }


  // primary color change
  async function pcChange({ detail }) {
    return save({ primaryColor: detail.value }, 'Primary Color');
  }

  // track hide-tab changes
  function hideTabChange(id, value) {
    const i = hideTabs.indexOf(id);
    if (i === -1) {
      hideTabs = hideTabs.concat(id);
    } else {
      hideTabs = hideTabs.filter((t) => t != id);
    }

    if (tOut) {
      clearTimeout(tOut);
    }

    let tabs = hideTabs.length > 0 ? hideTabs : ['_'];
    tOut = setTimeout(() => {
      save({ hideTabs: tabs }, 'Hide Tabs');
    }, 600);
  }

  // save changes to DB
  async function save(inp, msg) {
    if (busy) return;

    busy = true;
    const res = await client.mutation(SavePartnerMobileSettings, { id: partnerID, inp }).toPromise();
    busy = false;
    if (res.error) {
      alerts.error(msgTitle, res.error.message);
      return;
    }

    alerts.success(msgTitle, `Successfully saved "${msg}"`);
  }
</script>

<div class='flex gap-4 flex-col sm:flex-row flex-wrap' class:blur-sm={busy}>
  <CardWithBorderTitle title='Theme'>
    <div class='form-control'>
      <Label uppercase htmlContent='LOGO <small>{logoW}x{logoH}' />
      {#if logoURL}
        <img src={logoURL} width='{logoW}px' height='{logoH}px' alt='logo' />
      {/if}
      {#if logoUploadStatus}
        <p class='badge badge-info mt-2'>
          {logoUploadStatus}
        </p>
      {/if}
      <Cropper imageType={logImgType} on:crop={imgCrop} />
    </div>
    <InputRadio
      label='Primary Color'
      name='primaryColor'
      options={colors}
      value={primaryColor}
      on:change={pcChange}
    >
      <svelte:fragment slot='option' let:option>
        <div class='label-text flex flex-row space-x-2'>
          <div class='mt-0.5'>
            <svg width='15' height='15'>
              <rect width='15' height='15' style='fill:{option.id};stroke-width:3;stroke:rgb(0,0,0)' />
            </svg>
          </div>
          <div>
            {option.name}
          </div>
        </div>
      </svelte:fragment>
    </InputRadio>
  </CardWithBorderTitle>
  <CardWithBorderTitle title='Hide Tabs'>
    {#each tabNames as tn, idx}
      <InputToggle
        value={hideTabs.includes(tn.id)}
        name='tabs'
        label={tn.name}
        on:change={({detail})=> hideTabChange(tn.id, detail)}
      />
    {/each}
  </CardWithBorderTitle>
</div>
