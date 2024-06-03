<script setup>
import axios from 'axios';
import { onMounted, ref } from 'vue'
import { Bar, Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement } from 'chart.js';

import EditProduct from '../Owner/Edit-Product.vue'

onMounted(() => {
  listProducts()
  
}) 


ChartJS.register(Title, Tooltip, Legend, PointElement, LineElement, BarElement, CategoryScale, LinearScale);

const chartData = ref({
  labels: ['January', 'February', 'March'],
  datasets: [
    {
      label: 'Sell',
      data: [40, 20, 12],
      backgroundColor: ['#ff6384', '#36a2eb', '#cc65fe']
    }
  ]
});

const lineData = ref({
  labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
  datasets: [
    {
      label: 'Shoppers',
      backgroundColor: '#f87979',
      data: [40, 39, 10, 40, 39, 80, 40]
    }
  ]
})


const chartOptions = ref({
  responsive: true
});



const products = ref([]);
let totalProducts

const shopID = localStorage.getItem("owner_id")
const error = ref(null);


async function listProducts() {
  try {
    const response = await axios.get(`http://localhost:8000/shops/viewproducts?shop_id=${shopID}`);
    products.value = response.data;
    totalProducts = products.value.shop_products.length;
  } catch (err) {
    error  = err.message || 'An error occurred while fetching products';
  }
  
}



console.log(products)


</script>

<template>
  <div>
    <h3 class="text-3xl font-medium text-gray-700">
      Request
    </h3>

    <div class="mt-4">
      <div class="flex flex-wrap -mx-6">
        <div class="w-full px-6 sm:w-1/2 xl:w-1/3">
          <div
            class="flex items-center px-5 py-6 bg-white rounded-md shadow-sm"
          >
            
        <div class="flex flex-col gap-9">
            <div>
    <Bar :data="chartData" :options="chartOptions" />
  </div>

<<<<<<< HEAD
            <div class="mx-5">
              <h4 class="text-2xl font-semibold text-gray-700">
                8,282
              </h4>
              <div class="text-gray-500">
                New Orders
              </div>
            </div>
=======
  <div >
    <Line :data="lineData" :options="chartOptions" />
  </div>
</div>
>>>>>>> cad65aee50fdeb782aa2c6e5a8e7dd0968f06047
          </div>
        </div>

        <div class="w-full px-6 mt-6 sm:w-1/2 xl:w-1/3 sm:mt-0">
          <div
            class="flex items-center px-5 py-6 bg-white rounded-md shadow-sm"
          >
            <div class="p-3 bg-blue-600 bg-opacity-75 rounded-full">
              <svg
                class="w-8 h-8 text-white"
                viewBox="0 0 28 28"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M4.19999 1.4C3.4268 1.4 2.79999 2.02681 2.79999 2.8C2.79999 3.57319 3.4268 4.2 4.19999 4.2H5.9069L6.33468 5.91114C6.33917 5.93092 6.34409 5.95055 6.34941 5.97001L8.24953 13.5705L6.99992 14.8201C5.23602 16.584 6.48528 19.6 8.97981 19.6H21C21.7731 19.6 22.4 18.9732 22.4 18.2C22.4 17.4268 21.7731 16.8 21 16.8H8.97983L10.3798 15.4H19.6C20.1303 15.4 20.615 15.1004 20.8521 14.6261L25.0521 6.22609C25.2691 5.79212 25.246 5.27673 24.991 4.86398C24.7357 4.45123 24.2852 4.2 23.8 4.2H8.79308L8.35818 2.46044C8.20238 1.83722 7.64241 1.4 6.99999 1.4H4.19999Z"
                  fill="currentColor"
                />
                <path
                  d="M22.4 23.1C22.4 24.2598 21.4598 25.2 20.3 25.2C19.1403 25.2 18.2 24.2598 18.2 23.1C18.2 21.9402 19.1403 21 20.3 21C21.4598 21 22.4 21.9402 22.4 23.1Z"
                  fill="currentColor"
                />
                <path
                  d="M9.1 25.2C10.2598 25.2 11.2 24.2598 11.2 23.1C11.2 21.9402 10.2598 21 9.1 21C7.9402 21 7 21.9402 7 23.1C7 24.2598 7.9402 25.2 9.1 25.2Z"
                  fill="currentColor"
                />
              </svg>
            </div>

            <div class="mx-5">
              <h4 class="text-2xl font-semibold text-gray-700">
                200
              </h4>
              <div class="text-gray-500">
                Total Orders
              </div>
            </div>
          </div>
        </div>

        <div class="w-full px-6 mt-6 sm:w-1/2 xl:w-1/3 xl:mt-0">
          <div
            class="flex items-center px-5 py-6 bg-white rounded-md shadow-sm"
          >
            <div class="p-3 bg-pink-600 bg-opacity-75 rounded-full">
              <svg
                class="w-8 h-8 text-white"
                viewBox="0 0 28 28"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M6.99998 11.2H21L22.4 23.8H5.59998L6.99998 11.2Z"
                  fill="currentColor"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linejoin="round"
                />
                <path
                  d="M9.79999 8.4C9.79999 6.08041 11.6804 4.2 14 4.2C16.3196 4.2 18.2 6.08041 18.2 8.4V12.6C18.2 14.9197 16.3196 16.8 14 16.8C11.6804 16.8 9.79999 14.9197 9.79999 12.6V8.4Z"
                  stroke="currentColor"
                  stroke-width="2"
                />
              </svg>
            </div>

            <div class="mx-5">
              <h4 class="text-2xl font-semibold text-gray-700">
                {{ totalProducts }}
              </h4>
              <div class="text-gray-500">
                Available Products
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-8" />

    <div class="flex flex-col mt-8">
      <div class="py-2 -my-2 overflow-x-auto sm:-mx-6 sm:px-6 lg:-mx-8 lg:px-8">
        <div
          class="inline-block min-w-full overflow-hidden align-middle border-b border-gray-200 shadow sm:rounded-lg"
        >
          <table class="min-w-full">
            <thead>
              <tr>
                <th
                  class="px-6 py-3 text-xs font-medium leading-4 tracking-wider text-left text-gray-500 uppercase border-b border-gray-200 bg-gray-50"
                >
                  Name
                </th>
                <th
                  class="px-6 py-3 text-xs font-medium leading-4 tracking-wider text-left text-gray-500 uppercase border-b border-gray-200 bg-gray-50"
                >
<<<<<<< HEAD
                  product
=======
                  Description
>>>>>>> cad65aee50fdeb782aa2c6e5a8e7dd0968f06047
                </th>
                <th
                  class="px-6 py-3 text-xs font-medium leading-4 tracking-wider text-left text-gray-500 uppercase border-b border-gray-200 bg-gray-50"
                >
<<<<<<< HEAD
                  quantity
                </th>
              
=======
                  Price
                </th>
                <th
                  class="px-6 py-3 text-xs font-medium leading-4 tracking-wider text-left text-gray-500 uppercase border-b border-gray-200 bg-gray-50"
                >
                Quantity
                </th>
                <th class="px-6 py-3 border-b border-gray-200 bg-gray-50" />
>>>>>>> cad65aee50fdeb782aa2c6e5a8e7dd0968f06047
              </tr>
            </thead>

            <tbody class="bg-white">
              <tr v-for="(p, index) in products.shop_products" :key="index">
                <td
                  class="px-6 py-4 border-b border-gray-200 whitespace-nowrap"
                >
                  <div class="flex items-center">
                    <div class="flex-shrink-0 w-10 h-10">
                      <img
                        class="w-10 h-10 rounded-full"
                        :src="p.image"
                        alt=""
                      >
                    </div>

                    <div class="ml-4">
                      <div class="text-sm font-medium leading-5 text-gray-900">
                        {{ p.product_name }}
                      </div>
                      <div class="text-sm leading-5 text-gray-500">
                        {{  }}
                      </div>
                    </div>
                  </div>
                </td>

                <td
                  class="px-6 py-4 border-b border-gray-200 whitespace-nowrap"
                >
                  <div class="text-sm leading-5 text-gray-900">
                    {{ p.product_description }}
                  </div>
                  <div class="text-sm leading-5 text-gray-500">
                    {{  }}
                  </div>
                </td>

                <td
                  class="px-6 py-4 border-b border-gray-200 whitespace-nowrap"
                >
                  <span
                    class="inline-flex px-2 text-xs font-semibold leading-5 text-green-800 bg-green-100 rounded-full"
                  >{{ p.price }}</span>
                </td>

<<<<<<< HEAD
              
=======
                <td
                  class="px-6 py-4 text-sm leading-5 text-gray-500 border-b border-gray-200 whitespace-nowrap"
                >
                  {{ p.quantity }}
                </td>
>>>>>>> cad65aee50fdeb782aa2c6e5a8e7dd0968f06047

                <td
                  class="px-6 py-4 text-sm font-medium leading-5 text-right border-b border-gray-200 whitespace-nowrap"
                >
<<<<<<< HEAD
                  <a href="#" class="text-indigo-600 hover:text-indigo-900">Accept</a>
                </td>
                <td
                  class="px-6 py-4 text-sm font-medium leading-5 text-right border-b border-gray-200 whitespace-nowrap"
                >
                  <a href="#" class="text-indigo-600 hover:text-indigo-900">Decline</a>
=======
                <router-link  :to="{ name: 'editprod', params: { id: p._id } }"
           class="h-[350px] w-full">
       Edit 
      </router-link>
>>>>>>> cad65aee50fdeb782aa2c6e5a8e7dd0968f06047
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
