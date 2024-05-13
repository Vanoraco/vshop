<!-- <template>
  <div class="rounded-md max-w-96 bg-white shadow-sm">
    <div class="w-full px-6 pt-4 flex">
      <p class="text-lg font-semibold">{{ items.title }}</p>
    </div>
    <div class="grid gap-4 p-6 grid-cols-2">
      <figure
        v-for="(item, index) in items"
        :key="index"
        class="ring-1 ring-blue-100 overflow-hidden rounded-sm col-span-1 h-32 flex items-center justify-center"
      >
        <img
          :src="item.src"
          :alt="item.alt"
          class="rounded-sm h-full w-full hover:scale-110 object-cover transition-transform duration-500 ease-in-out"
        />
      </figure>
    </div>
    <div class="p-4">
      <router-link
        class="px-4 font-semibold hover:cursor-pointer text-sm text-indigo-700"
        :to="'/ShowMoreCategory/' + category"
        >See more</router-link
      >
    </div>
  </div>
</template>
    
    <script>
  export default {

    data(){
  items: [
    { 
      category: "summer",
      title: "Summer Collection",  
      src: "https://via.placeholder.com/150",
      alt: "Item 1"
    },
    { 
      title: "Summer Collection",  
      src: "https://via.placeholder.com/150",
      alt: "Item 2"
    },
    { 
      title: "Summer Collection",
      src: "https://via.placeholder.com/150",
      alt: "Item 3"
    },
    {
      title: "Summer Collection",  
      src: "https://via.placeholder.com/150",
      alt: "Item 4"
    }
  ],


    }
}
    
   
        
    
  </script>
    
<style scoped>
/* Add your component-specific styles here */
</style>
     -->