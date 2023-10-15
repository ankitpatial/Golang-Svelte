<script lang='ts'>
  import IconGear from '$lib/assets/svg/IconGear.svelte';
  import IconCheckCircle from '$lib/assets/svg/IconCheckCircle.svelte';
  import IconExclamationTriangle from '$lib/assets/svg/IconExclamationTriangle.svelte';
  import IconXCircle from '$lib/assets/svg/IconXCircle.svelte';
  import Page from '$lib/components/Page.svelte';
  import IconFolder from "../../../lib/assets/svg/IconFolder.svelte";
</script>

<Page title='Theme' subTitle='(Testing Components)'>

  <div>
    <!-- category:Table -->
    <div class='overflow-x-auto max-h-[500px]'>
      <table class='table table-compact w-full '>
        <thead>
        <tr>
          <th></th>
          <th>Name</th>
          <th>Job</th>
          <th>company</th>
          <th>location</th>
          <th>Last Login</th>
          <th>Favorite Color</th>
        </tr>
        </thead>
        <tbody>
        <tr>
          <th>1</th>
          <td><a class="link text-red-400 hover:text-green-400" href="#">Cy Ganderton</a></td>
          <td>Quality Control Specialist</td>
          <td>Littel, Schaden and Vandervort</td>
          <td>Canada</td>
          <td>12/16/2020</td>
          <td>Blue</td>
        </tr>
        <tr>
          <th>2</th>
          <td><a class="link" href="https://www.google.com" target="_blank">Hart Hagerty</a></td>
          <td>Desktop Support Technician</td>
          <td>Zemlak, Daniel and Leannon</td>
          <td>United States</td>
          <td>12/5/2020</td>
          <td>Purple</td>
        </tr>
        <tr>
          <th>3</th>
          <td>Brice Swyre</td>
          <td>Tax Accountant</td>
          <td>Carroll Group</td>
          <td>China</td>
          <td>8/15/2020</td>
          <td>Red</td>
        </tr>
        <tr>
          <th>4</th>
          <td>Marjy Ferencz</td>
          <td>Office Assistant I</td>
          <td>Rowe-Schoen</td>
          <td>Russia</td>
          <td>3/25/2021</td>
          <td>Crimson</td>
        </tr>
        <tr>
          <th>5</th>
          <td>Yancy Tear</td>
          <td>Community Outreach Specialist</td>
          <td>Wyman-Ledner</td>
          <td>Brazil</td>
          <td>5/22/2020</td>
          <td>Indigo</td>
        </tr>
        <tr>
          <th>6</th>
          <td>Irma Vasilik</td>
          <td>Editor</td>
          <td>Wiza, Bins and Emard</td>
          <td>Venezuela</td>
          <td>12/8/2020</td>
          <td>Purple</td>
        </tr>
        <tr>
          <th>7</th>
          <td>Meghann Durtnal</td>
          <td>Staff Accountant IV</td>
          <td>Schuster-Schimmel</td>
          <td>Philippines</td>
          <td>2/17/2021</td>
          <td>Yellow</td>
        </tr>
        <tr>
          <th>8</th>
          <td>Sammy Seston</td>
          <td>Accountant I</td>
          <td>O'Hara, Welch and Keebler</td>
          <td>Indonesia</td>
          <td>5/23/2020</td>
          <td>Crimson</td>
        </tr>
        <tr>
          <th>9</th>
          <td>Lesya Tinham</td>
          <td>Safety Technician IV</td>
          <td>Turner-Kuhlman</td>
          <td>Philippines</td>
          <td>2/21/2021</td>
          <td>Maroon</td>
        </tr>
        <tr>
          <th>10</th>
          <td>Zaneta Tewkesbury</td>
          <td>VP Marketing</td>
          <td>Sauer LLC</td>
          <td>Chad</td>
          <td>6/23/2020</td>
          <td>Green</td>
        </tr>
        <tr>
          <th>11</th>
          <td>Andy Tipple</td>
          <td>Librarian</td>
          <td>Hilpert Group</td>
          <td>Poland</td>
          <td>7/9/2020</td>
          <td>Indigo</td>
        </tr>
        <tr>
          <th>12</th>
          <td>Sophi Biles</td>
          <td>Recruiting Manager</td>
          <td>Gutmann Inc</td>
          <td>Indonesia</td>
          <td>2/12/2021</td>
          <td>Maroon</td>
        </tr>
        <tr>
          <th>13</th>
          <td>Florida Garces</td>
          <td>Web Developer IV</td>
          <td>Gaylord, Pacocha and Baumbach</td>
          <td>Poland</td>
          <td>5/31/2020</td>
          <td>Purple</td>
        </tr>
        <tr>
          <th>14</th>
          <td>Maribeth Popping</td>
          <td>Analyst Programmer</td>
          <td>Deckow-Pouros</td>
          <td>Portugal</td>
          <td>4/27/2021</td>
          <td>Aquamarine</td>
        </tr>
        <tr>
          <th>15</th>
          <td>Moritz Dryburgh</td>
          <td>Dental Hygienist</td>
          <td>Schiller, Cole and Hackett</td>
          <td>Sri Lanka</td>
          <td>8/8/2020</td>
          <td>Crimson</td>
        </tr>
        <tr>
          <th>16</th>
          <td>Reid Semiras</td>
          <td>Teacher</td>
          <td>Sporer, Sipes and Rogahn</td>
          <td>Poland</td>
          <td>7/30/2020</td>
          <td>Green</td>
        </tr>
        <tr>
          <th>17</th>
          <td>Alec Lethby</td>
          <td>Teacher</td>
          <td>Reichel, Glover and Hamill</td>
          <td>China</td>
          <td>2/28/2021</td>
          <td>Khaki</td>
        </tr>
        <tr>
          <th>18</th>
          <td>Aland Wilber</td>
          <td>Quality Control Specialist</td>
          <td>Kshlerin, Rogahn and Swaniawski</td>
          <td>Czech Republic</td>
          <td>9/29/2020</td>
          <td>Purple</td>
        </tr>
        <tr>
          <th>19</th>
          <td>Teddie Duerden</td>
          <td>Staff Accountant III</td>
          <td>Pouros, Ullrich and Windler</td>
          <td>France</td>
          <td>10/27/2020</td>
          <td>Aquamarine</td>
        </tr>
        <tr>
          <th>20</th>
          <td>Lorelei Blackstone</td>
          <td>Data Coordiator</td>
          <td>Witting, Kutch and Greenfelder</td>
          <td>Kazakhstan</td>
          <td>6/3/2020</td>
          <td>Red</td>
        </tr>
        </tbody>
      </table>
    </div>
    <div class='clear-both'>hello</div>

    <div class="mt-20">
      <a class="text-yellow-300 hover:bg-amber-400"><IconFolder/> Request an Estimate </a>

      <ul class="menu menu-xs rounded-lg max-w-xs w-full bg-secondary text-white dark:bg-yellow-200">
        <li><a><IconFolder/> COMPANY OVERVIEW </a></li>
        <li>
          <details>
            <summary>
              <IconFolder/> SERVICES</summary>
            <ul>
              <li>
                <details>
                  <summary> <IconFolder/>ROOFING </summary>
                  <ul>
                    <li><a class="link text-yellow-300 hover:bg-amber-400"><IconFolder/> Request an Estimate </a></li>
                    <li><a><IconFolder/> View Estimate </a></li>
                    <li><a><IconFolder/> Unassigned Jobs </a></li>
                    <li><a><IconFolder/> Assigned Jobs </a></li>
                    <li><a><IconFolder/> Pending Approval </a></li>
                    <li><a><IconFolder/> Approved </a></li>
                    <li><a><IconFolder/> Completed </a></li>
                  </ul>
                </details>
              </li>
              <li>
                <details>
                  <summary> <IconFolder/>SURVEYS </summary>
                  <ul>
                    <li><a><IconFolder/> View Dashboard</a></li>
                  </ul>
                </details>
              </li>
              <li><a><IconFolder/> Roofing Partner </a></li>
              <li><a><IconFolder/> Solar Partner </a></li>
              <li><a><IconFolder/> EPC </a></li>
              <li><a><IconFolder/> Integration Partner </a></li>
            </ul>
          </details>
        </li>
        <li>
          <details>
            <summary>
              <IconFolder/> PARTNERS</summary>
            <ul>
              <li><a><IconFolder/> Request an Estimate </a></li>
              <li><a><IconFolder/> View Estimate </a></li>
              <li><a><IconFolder/> Unassigned Jobs </a></li>
              <li><a><IconFolder/> Assigned Jobs </a></li>
              <li><a><IconFolder/> Pending Approval </a></li>
              <li><a><IconFolder/> Approved </a></li>
              <li><a><IconFolder/> Completed </a></li>
            </ul>
          </details>
        </li>
         </ul>
    </div>

    <!-- category:Theme -->
    <div class='mt-16'>
      <h4 class='uppercase font-bold bg-slate-50 py-1 mb-2'>Theme palate</h4>
      <div class='ml-6'>
        <div class='mb-6 pb-3 border-b'>
          <div>
            <h6 class='mb-1'>Color Palate</h6>
            <ul class='gap-2 space-y-3 ulList'>
              <li class='flex gap-3'>
                <div>
                  <div class='bg-primary'></div>
                  <small>primary </small>
                </div>
                <div>
                  <div class=' bg-primary-focus'></div>
                  <small>primary-focus </small>
                </div>
                <div>
                  <div class='  bg-primary-content'></div>
                  <small>primary-content </small>
                </div>
              </li>

              <li class='flex gap-3'>
                <div>
                  <div class=' bg-secondary'></div>
                  <small>secondary</small>
                </div>
                <div>
                  <div class=' bg-secondary-focus'></div>
                  <small>secondary-focus </small>
                </div>
                <div>
                  <div class=' bg-secondary-content'></div>
                  <small>secondary-content </small>
                </div>
              </li>

              <li class='flex gap-3'>
                <div>
                  <div class=' bg-accent'></div>
                  <small>accent</small>
                </div>
                <div>
                  <div class=' bg-accent-focus'></div>
                  <small>accent-focus</small>
                </div>
                <div>
                  <div class=' bg-accent-content'></div>
                  <small>accent-content</small>
                </div>
              </li>

              <li class='flex gap-3'>
                <div>
                  <div class='bg-neutral'></div>
                  <small>neutral</small>
                </div>
                <div>
                  <div class=' bg-neutral-focus'></div>
                  <small>neutral-focus</small>
                </div>
                <div>
                  <div class=' bg-neutral-content'></div>
                  <small>neutral-content</small>
                </div>
              </li>

              <li class='flex gap-3'>
                <div>
                  <div class='bg-info'></div>
                  <small>info</small>
                </div>
                <div>
                  <div class='bg-info-content'></div>
                  <small>info-content</small>
                </div>
              </li>

              <li class='flex gap-3'>
                <div>
                  <div class='bg-success'></div>
                  <small>success</small>
                </div>
                <div>
                  <div class='bg-success-content'></div>
                  <small>success-content</small>
                </div>
              </li>

              <li class='flex gap-3'>
                <div>
                  <div class='bg-warning'></div>
                  <small>warning</small>
                </div>
                <div>
                  <div class='bg-warning-content'></div>
                  <small>warning-content</small>
                </div>
              </li>

              <li class='flex gap-3'>
                <div>
                  <div class='bg-error'></div>
                  <small>error</small>
                </div>
                <div>
                  <div class='bg-error-content'></div>
                  <small>error-content</small>
                </div>
              </li>

            </ul>
          </div>
        </div>
      </div>
    </div>


    <!-- category:FONT -->
    <div>
      <h4 class='uppercase font-bold bg-slate-50 py-1 mb-2'>Font Size</h4>
      <div class='ml-6'>
        <div class='mb-6 pb-3 border-b'>
          <div>
            <h6 class='mb-1'>Heading Fonts</h6>
            <h1>h1 - heading 1</h1>
            <h2>h2 - heading 2</h2>
            <h3>h3 - heading 3</h3>
            <h4>h4 - heading 4</h4>
            <h5>h5 - heading 5</h5>
            <h6>h6 - heading 6</h6>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div>
            <h6 class='mb-1'>Typography</h6>
            <p class='text-xs'>The quick brown fox jumps over the lazy dog. - text-xs</p>
            <p class='text-sm'>The quick brown fox jumps over the lazy dog. - text-sm</p>
            <p class='text-base'>The quick brown fox jumps over the lazy dog. - text-base</p>
            <p class='text-lg'>The quick brown fox jumps over the lazy dog. - text-lg</p>
            <p class='text-xl'>The quick brown fox jumps over the lazy dog. - text-xl</p>
          </div>
        </div>
      </div>
    </div>


    <!-- category:BUTTON -->
    <div>
      <!-- category title -->
      <h4 class='uppercase font-bold bg-slate-50 py-1 mb-2'>Buttons</h4>
      <div class='ml-6'>
        <div class='mb-6 pb-3 border-b'>
          <div>
            <!-- section title -->
            <h6 class='mb-1'>Button default</h6>
            <button class='btn'>btn</button>
            <button class='btn btn-primary'>btn-primary</button>
            <button class='btn btn-secondary'>btn-secondary</button>
            <button class='btn btn-accent'>btn-accent</button>
            <button class='btn btn-ghost'>btn-ghost</button>
            <button class='btn btn-link'>Link</button>
          </div>
          <div class='mt-4'>
            <h6 class='mb-1'>Button with active state</h6>
            <button class='btn btn-active'>btn</button>
            <button class='btn btn-active btn-primary'>btn-primary</button>
            <button class='btn btn-active btn-secondary'>btn-secondary</button>
            <button class='btn btn-active btn-accent'>btn-accent</button>
            <button class='btn btn-active btn-ghost'>btn-ghost</button>
            <button class='btn btn-active btn-link'>Link</button>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <h6 class='mb-1'>Button with state colors</h6>
          <button class='btn btn-info'>btn-info</button>
          <button class='btn btn-success'>btn-success</button>
          <button class='btn btn-warning'>btn-warning</button>
          <button class='btn btn-error'>btn-wrror</button>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <h6 class='mb-1'>Outline Button</h6>
          <button class='btn btn-outline'>btn-outline</button>
          <button class='btn btn-outline btn-primary'>btn-primary</button>
          <button class='btn btn-outline btn-secondary'>btn-secondary</button>
          <button class='btn btn-outline btn-accent'>btn-accent</button>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <h6 class='mb-1'>Button sizes</h6>
          <button class='btn btn-lg'>btn-lg</button>
          <button class='btn'>btn (normal)</button>
          <button class='btn btn-sm'>btn-sm</button>
          <button class='btn btn-xs'>btn-xs</button>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <h6 class='mb-1'>Disabled Button</h6>
          <button class='btn' disabled='disabled'>Disabled using attribute</button>
          <button class='btn btn-disabled' tabindex='-1' role='button' aria-disabled='true'>Disabled using class name
          </button>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <h6 class='mb-1'>Square Button</h6>
          <button class='btn btn-square'>
            <IconGear />
          </button>
          <button class='btn btn-square btn-outline'>
            <IconGear />
          </button>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <h6 class='mb-1'>Circle Button</h6>

          <button class='btn btn-circle btn-lg'>
            <IconGear />
          </button>

          <button class='btn btn-circle'>
            <IconGear />
          </button>

          <button class='btn btn-circle btn-sm'>
            <IconGear />
          </button>

          <button class='btn btn-circle btn-xs'>
            <IconGear />
          </button>

          <button class='btn btn-circle btn-outline'>
            <IconGear />
          </button>

        </div>

      </div>
    </div>


    <!-- category:FORM/ DATA INPUT -->
    <div>
      <!-- category title -->
      <h4 class='uppercase font-bold bg-slate-50 py-1 mb-2'>FORM</h4>
      <div class='ml-6'>
        <div class='mb-6 pb-3 border-b'>
          <div>
            <h6 class='mb-1'>Text input</h6>
            <input type='text' placeholder='Type here' class='input w-full max-w-xs' />
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div>
            <h6 class='mb-1'>Text input with border</h6>
            <input type='text' placeholder='Type here' class='input input-bordered w-full max-w-xs' />
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div>
            <h6 class='mb-1'>With form-control and labels</h6>
            <div class='form-control w-full max-w-xs'>
              <label class='label'>
                <span class='label-text'>What is your name?</span>
                <span class='label-text-alt'>Alt label</span>
              </label>
              <input type='text' placeholder='Type here' class='input input-bordered w-full max-w-xs' />
              <label class='label'>
                <span class='label-text-alt'>Alt label</span>
                <span class='label-text-alt'>Alt label</span>
              </label>
            </div>
          </div>
        </div>


        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Sizes</h6>
            <div class='flex flex-col gap-4'>
              <!-- xs -->
              <input type='text' placeholder='input-xs' class='input input-bordered input-xs w-full max-w-xs' />
              <!-- sm -->
              <input type='text' placeholder='input-sm' class='input input-bordered input-sm w-full max-w-xs' />
              <!-- md -->
              <input type='text' placeholder='input-md (default)'
                     class='input input-bordered input-md w-full max-w-xs' />
              <!-- lg -->
              <input type='text' placeholder='input-lg' class='input input-bordered input-lg w-full max-w-xs' />
            </div>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Disabled</h6>
            <input type='text' placeholder="You can't touch this" class='input w-full max-w-xs' disabled />
          </div>
        </div>


        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Checkbox</h6>
            <input type='checkbox' checked='checked' class='checkbox' />
            <input type='checkbox' checked='checked' class='checkbox checkbox-primary' />
            <input type='checkbox' class='checkbox' disabled />
            <input type='checkbox' class='checkbox' disabled checked />
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Radio</h6>
            <input type='radio' name='radio-1' class='radio' checked />
            <input type='radio' name='radio-1' class='radio radio-primary' checked />
            <input type='radio' name='radio-1' class='radio' />
            <input type='radio' name='radio-5' class='radio' disabled checked />
            <input type='radio' name='radio-5' class='radio' disabled />
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>With label and form-control and custom colors!</h6>
            <div class='form-control w-56'>
              <label class='label cursor-pointer'>
                <span class='label-text'>Red pill</span>
                <input type='radio' name='radio-6' class='radio checked:bg-red-500' checked />
              </label>
            </div>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='form-control w-56'>
            <label class='label cursor-pointer'>
              <span class='label-text'>blue pill</span>
              <input type='radio' name='radio-6' class='radio checked:bg-blue-500' checked />
            </label>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Select</h6>
            <div class='form-control'>
              <select class='select w-full max-w-xs'>
                <option disabled selected>Pick your favorite Simpson</option>
                <option>Homer</option>
                <option>Marge</option>
                <option>Bart</option>
                <option>Lisa</option>
                <option>Maggie</option>
              </select>
            </div>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Select with border</h6>
            <div class='form-control'>
              <select class='select select-bordered w-full max-w-xs'>
                <option disabled selected>Who shot first?</option>
                <option>Han Solo</option>
                <option>Greedo</option>
              </select>
            </div>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Ghost (no background)</h6>
            <div class='form-control'>
              <select class='select select-ghost w-full max-w-xs'>
                <option disabled selected>Pick the best JS framework</option>
                <option>Svelte</option>
                <option>Vue</option>
                <option>React</option>
              </select>
            </div>
          </div>
        </div>


        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Sizes</h6>
            <div class='form-control gap-4 '>
              <!-- xs -->
              <select class='select select-bordered select-xs w-full max-w-xs'>
                <option disabled selected>select-xs</option>
                <option>Tiny Apple</option>
                <option>Tiny Orange</option>
                <option>Tiny Tomato</option>
              </select>

              <!-- sm -->
              <select class='select select-bordered select-sm w-full max-w-xs'>
                <option disabled selected>select-sm</option>
                <option>Small Apple</option>
                <option>Small Orange</option>
                <option>Small Tomato</option>
              </select>

              <!-- md -->
              <select class='select select-bordered w-full max-w-xs'>
                <option disabled selected>Normal (default)</option>
                <option>Normal Apple</option>
                <option>Normal Orange</option>
                <option>Normal Tomato</option>
              </select>

              <!-- lg -->
              <select class='select select-bordered select-lg w-full max-w-xs'>
                <option disabled selected>select-lg</option>
                <option>Large Apple</option>
                <option>Large Orange</option>
                <option>Large Tomato</option>
              </select>
            </div>
          </div>
        </div>

        <div class='mb-6 pb-3 border-b'>
          <div class='gap-2'>
            <h6 class='mb-1'>Disabled</h6>
            <div class='form-control w-56'>
              <select class='select w-full max-w-xs' disabled>
                <option>You can't touch this</option>
              </select>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>


  <!-- category:alert -->
  <div>
    <h4 class='uppercase font-bold bg-slate-50 py-1 mb-2'>Alert</h4>
    <div class='ml-6'>
      <div class='mb-6 pb-3 border-b'>
        <div>
          <h6 class='mb-1'>Alert</h6>
          <div class='alert shadow-lg'>
            <div>
              <svg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24'
                   class='stroke-info flex-shrink-0 w-6 h-6'>
                <path stroke-linecap='round' stroke-linejoin='round' stroke-width='2'
                      d='M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z'></path>
              </svg>
              <span>12 unread messages. Tap to see.</span>
            </div>
          </div>
        </div>

        <div>
          <h6 class='mb-1'>Info color</h6>
          <div class='alert alert-info shadow-lg'>
            <div>
              <svg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24'
                   class='stroke-current flex-shrink-0 w-6 h-6'>
                <path stroke-linecap='round' stroke-linejoin='round' stroke-width='2'
                      d='M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z'></path>
              </svg>
              <span>New software update available.</span>
            </div>
          </div>
        </div>

        <div>
          <h6 class='mb-1'>Success color</h6>
          <div class='alert alert-success shadow-lg'>
            <div>
              <IconCheckCircle />
              <span>Your purchase has been confirmed!</span>
            </div>
          </div>
        </div>

        <div>
          <h6 class='mb-1'>Warning color</h6>
          <div class='alert alert-warning shadow-lg'>
            <div>
              <IconExclamationTriangle />
              <span>Warning: Invalid email address!</span>
            </div>
          </div>
        </div>

        <div>
          <h6 class='mb-1'>Error</h6>
          <div class='alert alert-error shadow-lg'>
            <div>
              <IconXCircle />
              <span>Error! Task failed successfully.</span>
            </div>
          </div>
        </div>


      </div>

    </div>
  </div>

</Page>

<style>
  /*a {*/
  /*  @apply text-amber-300 hover:text-green-400;*/
  /*}*/
  .ulList li div:first-child {
    width: 130px;
    height: 24px;
  }
</style>
