<script lang='ts'>
  import { authUser } from '$lib/stores/auth';
  import Cropper from '$lib/components/Cropper/Cropper.svelte';
  import { UpdateProfile } from '$lib/gql/account';
  import alerts from '$lib/stores/alerts';
  import { getContextClient } from '@urql/svelte';
  import { validatePwd } from '$lib/utils/validate';
  import { PublicDataUploadUrl } from '$lib/gql/document';
  import { uploadFile } from '$lib/components/FileUpload/uploader';
  import Input from '$lib/components/form/Input.svelte';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import Btn from '$lib/components/form/Btn.svelte';
  import InputPassword from '$lib/components/form/InputPassword.svelte';
  import Page from '$lib/components/Page.svelte';

  const msgTitle = 'Save Profile',
    client = getContextClient();

  let busy = false,
    imageType = 'image/jpeg',
    picUploadStatus = undefined,
    avatarURL = $authUser?.picture || '',
    firstName = $authUser?.firstName || '',
    lastName = $authUser?.lastName || '',
    phone = $authUser?.phone || '',
    oldPwd = '',
    newPwd = '',
    confirmNewPwd = '';
  ``;

  // save profile image
  async function imgCrop({ detail }) {
    const name = 'avatar.jpeg';
    avatarURL = detail.url;
    const f = new File([detail.blob], name, { type: imageType });

    // get upload url data
    const res = await client.mutation(PublicDataUploadUrl, {
      dir: $authUser?.id,
      name,
      section: 'Avatar',
      fileName: name,
      fileType: imageType,
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
          save({ picture: d.publicUrl }, 'Profile Image');
          picUploadStatus = '';
          break;
        case 'error':
          picUploadStatus = 'Error: ' + err;
          break;
        default:
          picUploadStatus = 'Status: ' + status;
      }
    });
  }

  // save detail
  async function saveDetail() {
    if (await save({ firstName, lastName, phone }, 'Profile')) {
      $authUser.firstName = firstName;
      $authUser.lastName = lastName;
      $authUser.phone = phone;
    }
  }

  async function updatePwd() {
    const title = 'Manage Password';
    if (!oldPwd || !newPwd || !confirmNewPwd) {
      alerts.warning(title, 'all password fields are required');
      return;
    }

    if (newPwd !== confirmNewPwd) {
      alerts.warning(title, 'confirm password is not same');
      return;
    }

    const err = validatePwd(newPwd);
    if (err) {
      alerts.warning(title, err);
      return;
    }

    await save({ oldPwd, newPwd, confirmNewPwd }, title);
  }

  // save changes to DB
  async function save(inp, msg): Promise<boolean> {
    if (busy) return false;

    busy = true;
    const res = await client.mutation(UpdateProfile, { inp }).toPromise();
    busy = false;
    if (res.error) {
      alerts.error(msgTitle, res.error.message);
      return false;
    }

    if (inp.picture) {
      $authUser.picture = inp.picture;
    }

    alerts.success(msgTitle, `Successfully saved "${msg}"`);
    return true;
  }
</script>

<Page title='Profile'>
  <div class='w-[500px]' class:blur-sm={busy}>
    <div class='avatar placeholder'>
      <div class='bg-neutral-focus text-neutral-content rounded-full w-24'>
        {#if (avatarURL)}
          <img src={avatarURL} alt='profile image' />
        {:else }
        <span class='text-3xl'>
        {$authUser.firstName[0]}{$authUser.lastName[0]}
      </span>
        {/if}
      </div>
    </div>

    <div class='mt-2 mb-10'>
      <Cropper
        ghost
        {imageType}
        label={picUploadStatus ? picUploadStatus : 'Upload New Image'}
        on:crop={imgCrop}
        aspectRatio={1}
      />
    </div>

    <form on:submit={saveDetail}>
      <Input label='Email' value={$authUser.email} disabled />
      <Input label='First Name' bind:value={firstName} required />
      <Input label='Last Name' bind:value={lastName} required />
      <InputPhone label='Phone' bind:value={phone} required />
      <Btn primary type='submit'>
        Save
      </Btn>
    </form>

    <form on:submit={updatePwd}>
      <h3 class='mt-10'>Manage Password</h3>
      <Input label='Old Password' type='password' maxlength='100' bind:value={oldPwd} required />
      <InputPassword label='New Password' bind:value={newPwd} required />
      <Input label='Confirm New Password' type='password' bind:value={confirmNewPwd} required />
      <Btn primary type='submit'>
        Update Password
      </Btn>
    </form>
  </div>
</Page>
