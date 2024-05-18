<template>
  <div>
    <h1 class="text-4xl font-bold mb-6">{{ shop.name }}</h1>
    <p class="text-lg mb-6">{{ shop.description }}</p>

    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
      <div
        v-for="product in products"
        :key="product.id"
        class="bg-white rounded-lg shadow-md p-4 hover:shadow-xl transition duration-300"
      >
        <img :src="product.image" :alt="product.name" class="w-full h-48 object-cover rounded-md mb-4" />
        <h2 class="text-xl font-semibold mb-2">{{ product.name }}</h2>
        <p class="text-gray-500">{{ product.description }}</p>
        <p class="text-gray-800 font-semibold mt-2">{{ product.price }}</p>
        <button class="mr-2">3D</button>
        <button>Add to cart</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      shop: {
        name: 'Phina Electronics',
        description: 'Latest and greatest technologies Phina electronics',
      },
      products: [
        {
          id: 1,
          name: 'samsung edge 10',
          description: 'latest samsung edge',
          price: '$400.99',
          image: 'https://cdn.pixabay.com/photo/2017/08/15/10/14/samsung-galaxy-s8-2643381_640.jpg',
        },
        {
          id: 1,
          name: 'samsung edge 10',
          description: 'latest samsung edge',
          price: '$400.99',
          image: 'https://cdn.pixabay.com/photo/2017/08/15/10/14/samsung-galaxy-s8-2643381_640.jpg',
        },
        {
          id: 1,
          name: 'samsung edge 10',
          description: 'latest samsung edge',
          price: '$400.99',
          image: 'https://cdn.pixabay.com/photo/2017/08/15/10/14/samsung-galaxy-s8-2643381_640.jpg',
        },
        {
          id: 1,
          name: 'samsung edge 10',
          description: 'latest samsung edge',
          price: '$400.99',
          image: 'https://cdn.pixabay.com/photo/2017/08/15/10/14/samsung-galaxy-s8-2643381_640.jpg',
        },
        {
          id: 1,
          name: 'samsung edge 10',
          description: 'latest samsung edge',
          price: '$400.99',
          image: 'https://cdn.pixabay.com/photo/2017/08/15/10/14/samsung-galaxy-s8-2643381_640.jpg',
        },
      ],
    };
  },
};
</script>

<style>
</style>
