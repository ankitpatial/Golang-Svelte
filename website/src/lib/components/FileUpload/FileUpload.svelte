<script lang='ts'>
  import {uploadFile} from './uploader';
  import type {Document, SignedUrlCB} from './d';

  export let id = crypto.randomUUID();
  export let name: string;
  export let multiple = false;
  export let accept = '*/*';
  export let disabled = false;
  export let w = 'w-full';
  export let h = 'h-10';
  export let mt = 'mt-0';
  export let uploadUrl: SignedUrlCB = () => {
    return Promise.reject(new Error('cbSignedUrl is not defined'));
  };
  export let uploadComplete: (inpName: string, docs: Array<Document>) => void = () => {
    return Promise.reject(new Error('cb uploadComplete  is not defined'));
  };

  export {uploadFile} from './uploader';

  export interface FileInfo {
    id: number;
    name: string;
    file: File;
    busy: boolean;
    status?: string;
    doc?: Document;
    err?: string;
  }

  let inp;
  let files: Array<FileInfo>;
  let busy = false;
  let acceptedTypes;
  let disableCss;

  $: allowed = accept.split(',').map(o => o.trim());
  $: {
    if (accept !== '*/*') {
      let a = allowed.map(o => o.endsWith('/*')
        ? o.replace('/*', '')
        : o.substring(o.indexOf('/') + 1));

      acceptedTypes = a.join(', ');
    }
  }
  $: disableCss = disabled ? '' : 'transition cursor-pointer hover:border-gray-600 focus:outline-none';


  function open() {
    if (busy) return;

    inp.click();
  }

  function selected(e) {
    onFiles([...e.target.files]);
    upload(files);
  }

  function drop(e) {
    onFiles([...e.dataTransfer.files]);
    upload(files);
  }

  function onFiles(list) {
    if (!list || list.length == 0) {
      files = [];
      return;
    }

    if (accept !== '*/*') {
      let idx, t, ok;
      list = list.filter((f: File, i) => {
        // console.log(i, f.type)
        idx = allowed.findIndex(o => {
          if (o.endsWith('/*')) {
            t = o.replace('/*', '');
            ok = f.type.startsWith(t);
          } else {
            t = o;
            ok = f.type === t;
          }
          // console.log('  ', t, ok);
          return ok;
        });
        return idx > -1;
      });

    }

    if (list.length === 0) return;

    const now = Date.now();
    files = (multiple ? [...list] : [].concat(list[0])).map((f: File, idx) => ({
      id: (now + idx),
      name: f.name,
      file: f,
      busy: false,
      status: 'pending'
    }));
  }

  async function upload(list: Array<FileInfo>) {
    if (!list) return;

    return Promise.resolve(list.map(f => {
      return uploadUrl(name, f.file)
        .then((res) => {
          f.doc = res;
          return uploadFile(f.id, f.file, name, res.uploadUrl, res.meta, fileStatusChange);
        })
        .catch((ex) => {
          fileStatusChange(f.id, 'error', ex.message || ex);
        });
    }));
  }

  function fileStatusChange(id: number, status: string, err: undefined | string) {
    let idx = files.findIndex(o => o.id === id);
    if (idx > -1) {
      files[idx].status = status;
      files[idx].err = err;
    }

    const notCompleteIdx = files.findIndex(f => f.status != 'completed');
    if (notCompleteIdx == -1) {
      const docs = files.map(f => {
        if (f.doc) f.doc.ready = true;
        return f.doc;
      });
      uploadComplete(name, docs);
      files = [];
    }
  }
</script>

<div
  on:dragover|preventDefault|stopPropagation={(e)=>{e.dataTransfer.dropEffect = 'copy'}}
  on:drop|preventDefault|stopPropagation={drop}
>
  <label
    for={id}
    class='flex justify-center {w} {h} px-4 bg-slate-50 border-2 border-slate-300 border-dashed
    rounded-md appearance-none hover:border-2 hover:border-slate-400 {disableCss} {mt}'
  >
    {#if (files && files.length > 0)}
      <ul class='flex items-center space-x-2'>
        {#each files as file, idx (file.id)}
          <li>
            {file.name}
            {#if file.err}
              <span class='badge badge-error badge-sm'>error: {file.err}</span>
            {:else if file.status}
              <span class='badge badge-outline badge-sm'>{file.status}</span>
            {/if}
          </li>
        {/each}
      </ul>
    {:else }
        <span class='flex items-center space-x-2'>
            <svg xmlns='http://www.w3.org/2000/svg' class='w-6 h-6 text-gray-600' fill='none' viewBox='0 0 24 24'
                 stroke='currentColor' stroke-width='2'>
                <path stroke-linecap='round' stroke-linejoin='round'
                      d='M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12'/>
            </svg>
            <span class='font-medium text-gray-600'>
                Drop {(multiple ? 'files' : 'a file')}{acceptedTypes ? `(${acceptedTypes})` : ''} or
                <span class='text-blue-600 underline'>browse</span>
            </span>
        </span>
    {/if}

    <input id={id} type='file' class='hidden' {accept} on:input={selected} multiple={multiple} disabled={disabled}>
  </label>
</div>

