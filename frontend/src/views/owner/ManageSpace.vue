<script setup lang="ts">
defineProps({
  title: {
    type: String,
    default: "",
  },
});
</script>

<template>
  <div>
    <div
      class="flex w-full flex-wrap items-center justify-between rounded-2xl px-12 py-6 md:flex-nowrap"
      :style="{
        backgroundColor: `url('https://www.transparenttextures.com/patterns/cubes.png')`,
        backgroundImage: `linear-gradient(to right top, #cf4af3,#e73bd7, #f631bc,#fd31a2,#ff3a8b, #ff4b78,#ff5e68,#ff705c,#ff8c51,#ffaa49,#ffc848,#ffe652)`,
      }"
    >
      <div>
        <h3 class="m-0 flex items-center text-lg font-medium"></h3>
        <div class="my-4 overflow-hidden text-sm font-normal text-white">
          Try our Premium plan to get amaizing experience
        </div>
        <button class="rounded-2xl bg-blue-700 px-4 py-2">
          Premium Plan
        </button>
      </div>
      <img
        class="w-24 object-cover object-center md:w-44"
        src="../../assets/images/hand-made-virtual-shopping-made-tangible-3d-render-of-hand-holding-bag-emerging-from-smartphone_9627836.jpg"
        alt=""
      />
    </div>
    <div class="mx-auto flex justify-center font-extrabold">
      <h1>BOKACHA PHONE SALES</h1>
    </div>

    <div class="mb-3 mt-6 text-gray-400">Items uploaded</div>
    <div
      class="rounded-2xl p-4 font-normal"
      style="background-color: rgb(146 151 179 / 13%)"
    >
      <div class="mb-6 flex justify-end border-b border-gray-600 pb-2">
        <div class="mr-auto flex">
          <img
            viewBox="0 0 52 52"
            class="mr-4 h-7 w-7 rounded"
            style="border: 1px solid #3291b8"
            src=""
         />
           
          
          Iphone 14 pro max
        </div>
        <div class="mx-auto flex">
          <div>
            <button
              class="rounded-full border border-gray-500 px-5 py-1 text-gray-500"
            >
              detail
            </button>
          </div>
        </div>
        <div class="ml-auto flex">
          <div>
            <button
              class="rounded-full border border-gray-500 px-5 py-1 text-gray-500"
            >
              delete
            </button>
          </div>
        </div>
      </div>
      <div class="mb-6 flex justify-end border-b border-gray-600 pb-2">
        <div class="mr-auto flex">
          <img
            viewBox="0 0 52 52"
            class="mr-4 h-7 w-7 rounded"
            style="border: 1px solid #3291b8"
            src=""
         />
          Illustrator
        </div>
        <div class="mx-auto flex">
          <div>
            <button
              class="rounded-full border border-gray-500 px-5 py-1 text-gray-500"
            >
              detail
            </button>
          </div>
        </div>
        <div class="ml-auto flex">
          <div>
            <button
              class="rounded-full border border-gray-500 px-5 py-1 text-gray-500"
            >
              delete
            </button>
          </div>
        </div>
      </div>
      <div class="flex justify-end">
        <div class="mr-auto flex">
          <img
            viewBox="0 0 52 52"
            class="mr-4 h-7 w-7 rounded"
            style="border: 1px solid #3291b8"
            src=""
         />
          After Effects
        </div>
        <div class="mx-auto flex">
          <div>
            <button
              class="rounded-full border border-gray-500 px-5 py-1 text-gray-500"
            >
              detail
            </button>
          </div>
        </div>
        <div class="ml-auto flex">
          <div>
            <button
              class="rounded-full border border-gray-500 px-5 py-1 text-gray-500"
            >
              delete
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <div class="mb-3 mt-6 text-gray-400">Request</div>


    <div class="mt-8 rounded-2xl" style="background: rgb(146 151 179 / 13%)">
      <div class="container mx-auto">
        <div class="py-8">
          <div class="py-4">
            <div class="max-w-full overflow-x-auto rounded-lg">
              <table class="w-full leading-normal text-black">
                <thead>
                  <tr>
                    <th
                      scope="col"
                      class="border-b border-gray-200 px-5 py-3 text-left text-sm font-normal uppercase"
                    >
                      User
                    </th>
                    <th
                      scope="col"
                      class="border-b border-gray-200 px-5 py-3 text-left text-sm font-normal uppercase"
                    >
                      Product
                    </th>
                 
                    <th
                      scope="col"
                      class="border-b border-gray-200 px-5 py-3 text-left text-sm font-normal uppercase"
                    >
                      ORDER AT
                    </th>
                    <th
                      scope="col"
                      class="border-b border-gray-200 px-5 py-3 text-left text-sm font-normal uppercase"
                    >
                      STATUS
                    </th>
                    <th
                      scope="col"
                      class="border-b border-gray-200 px-5 py-3 text-left text-sm font-normal uppercase"
                    >
                      detail
                    </th>
                    <th
                      scope="col"
                      class="border-b border-gray-200 px-5 py-3 text-left text-sm font-normal uppercase"
                    >
                   quantity
                    </th>
                    <th
                      scope="col"
                      class="border-b border-gray-200 px-5 py-3 text-left text-sm font-normal uppercase"
                    />
                  </tr>
                </thead>
                <tbody >
                  <tr>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <div class="flex items-center">
                        <div class="shrink-0">
                          <a href="#" class="relative block">
                            <img
                              alt="profil"
                              src="https://res.cloudinary.com/beloved/image/upload/v1623196231/Assets/person/8_yfarwk.jpg"
                              class="mx-auto h-10 w-10 rounded-full object-cover"
                            />
                          </a>
                        </div>
                        <div class="ml-3">
                          <p class="whitespace-nowrap">seare hagos</p>
                        </div>
                      </div>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">Admin</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">12/09/2020</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <span
                        class="relative inline-block px-3 py-1 font-semibold leading-tight"
                      >
                        <span
                          aria-hidden="true"
                          class="absolute inset-0 rounded-full bg-green-200 opacity-50"
                        />
                        <span class="relative">active</span>
                      </span>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <a href="#" class="text-indigo-600 hover:text-indigo-900">
                        detail
                      </a>
                    </td>
                  </tr>
                  <tr>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <div class="flex items-center">
                        <div class="shrink-0">
                          <a href="#" class="relative block">
                            <img
                              alt="profil"
                              src="https://res.cloudinary.com/beloved/image/upload/v1623196231/Assets/person/4_iauuag.jpg"
                              class="mx-auto h-10 w-10 rounded-full object-cover"
                            />
                          </a>
                        </div>
                        <div class="ml-3">
                          <p class="whitespace-nowrap">Marcus coco</p>
                        </div>
                      </div>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">Designer</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">01/10/2012</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <span
                        class="relative inline-block px-3 py-1 font-semibold leading-tight text-green-900"
                      >
                        <span
                          aria-hidden="true"
                          class="absolute inset-0 rounded-full bg-green-200 opacity-50"
                        />
                        <span class="relative">active</span>
                      </span>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <a href="#" class="text-indigo-600 hover:text-indigo-900">
                        detail
                      </a>
                    </td>
                  </tr>
                  <tr>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <div class="flex items-center">
                        <div class="shrink-0">
                          <a href="#" class="relative block">
                            <img
                              alt="profil"
                              src="https://res.cloudinary.com/beloved/image/upload/v1623196231/Assets/person/3_iscinv.jpg"
                              class="mx-auto h-10 w-10 rounded-full object-cover"
                            />
                          </a>
                        </div>
                        <div class="ml-3">
                          <p class="whitespace-nowrap">Eric marc</p>
                        </div>
                      </div>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">Developer</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">02/10/2018</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <span
                        class="relative inline-block px-3 py-1 font-semibold leading-tight"
                      >
                        <span
                          aria-hidden="true"
                          class="absolute inset-0 rounded-full bg-green-200 opacity-50"
                        />
                        <span class="relative">active</span>
                      </span>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <a href="#" class="text-indigo-600 hover:text-indigo-900">
                        detail
                      </a>
                    </td>
                  </tr>
                  <tr>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <div class="flex items-center">
                        <div class="shrink-0">
                          <a href="#" class="relative block">
                            <img
                              alt="profil"
                              src="https://res.cloudinary.com/beloved/image/upload/v1623196231/Assets/person/6_v4ytve.jpg"
                              class="mx-auto h-10 w-10 rounded-full object-cover"
                            />
                          </a>
                        </div>
                        <div class="ml-3">
                          <p class="whitespace-nowrap">Julien Huger</p>
                        </div>
                      </div>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">User</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <p class="whitespace-nowrap">23/09/2010</p>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <span
                        class="relative inline-block px-3 py-1 font-semibold leading-tight text-green-900"
                      >
                        <span
                          aria-hidden="true"
                          class="absolute inset-0 rounded-full bg-green-200 opacity-50"
                        />
                        <span class="relative">active</span>
                      </span>
                    </td>
                    <td class="border-b border-gray-200 p-5 text-sm">
                      <a href="#" class="text-indigo-600 hover:text-indigo-900">
                        detail
                      </a>
                    </td>
                  </tr>
                </tbody>
              </table>
              <div
                class="xs:flex-row xs:justify-between flex flex-col items-center p-5"
              >
                <div class="flex items-center">
                  <button
                    type="button"
                    class="w-full rounded-l-xl border bg-white p-4 text-base text-gray-600 hover:bg-gray-100"
                  >
                    <svg
                      width="9"
                      fill="currentColor"
                      height="8"
                      class=""
                      viewBox="0 0 1792 1792"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M1427 301l-531 531 531 531q19 19 19 45t-19 45l-166 166q-19 19-45 19t-45-19l-742-742q-19-19-19-45t19-45l742-742q19-19 45-19t45 19l166 166q19 19 19 45t-19 45z"
                      />
                    </svg>
                  </button>
                  <button
                    type="button"
                    class="w-full border-y bg-white px-4 py-2 text-base text-indigo-500 hover:bg-gray-100"
                  >
                    1
                  </button>
                  <button
                    type="button"
                    class="w-full border bg-white px-4 py-2 text-base text-gray-600 hover:bg-gray-100"
                  >
                    2
                  </button>
                  <button
                    type="button"
                    class="w-full border-y bg-white px-4 py-2 text-base text-gray-600 hover:bg-gray-100"
                  >
                    3
                  </button>
                  <button
                    type="button"
                    class="w-full border bg-white px-4 py-2 text-base text-gray-600 hover:bg-gray-100"
                  >
                    4
                  </button>
                  <button
                    type="button"
                    class="w-full rounded-r-xl border-y border-r bg-white p-4 text-base text-gray-600 hover:bg-gray-100"
                  >
                    <svg
                      width="9"
                      fill="currentColor"
                      height="8"
                      class=""
                      viewBox="0 0 1792 1792"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M1363 877l-742 742q-19 19-45 19t-45-19l-166-166q-19-19-19-45t19-45l531-531-531-531q-19-19-19-45t19-45l166-166q19-19 45-19t45 19l742 742q19 19 19 45t-19 45z"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
