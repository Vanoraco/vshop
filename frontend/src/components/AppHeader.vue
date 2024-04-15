<template>
    <div class="font-serif mt-12 flex justify-center font-medium mx-auto">
       <ul class="flex flex-row gap-12 text-xl mx-auto">
        <li>
          Shops
        </li>
        <li>
          SignUp
        </li>
        <li>
          SignIn
        </li>
       </ul>
       
       <ul>
         <li>
          <IconField iconPosition="left">
            <InputIcon class="pi pi-search"> </InputIcon>
            <InputText placeholder="Search" />
          </IconField>
         </li>
       </ul>
    </div>
  </template>

  <script>
     import IconField from 'primevue/iconfield';
     import InputIcon from 'primevue/inputicon';
     import InputText from 'primevue/inputtext';

     export default {
      name: 'appheader',
      components: {
        IconField,
        InputIcon,
        InputText
      }
     }
   </script>