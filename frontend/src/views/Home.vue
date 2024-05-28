<template>
  <Carousels />

  <br />
  <br />

  <hero-section />
  <div class="ml-3 shop-header -mt-20 flex justify-center text-4xl mb-20">
    <p class="text-5xl font-extrabold">Top Shops</p>
  </div>

  <shop />
  <div class="flex justify-center mt-9">
    <router-link to="/shops"
      ><p class="bg-[#DB4444] text-white px-5 py-1 shop-header">
        View All Shops
      </p>
      </router-link>
  </div>

  
  <!-- <div class="flex justify-center mt-9">
    <router-link to="/shops"
      ><p class="bg-[#DB4444] text-white px-5 py-1 shop-header">
        View All Shops
      </p></router-link
    >
  </div> -->

  <div class="ml-3 shop-header mt-12">
    <p class="text-3xl font-extrabold">Explore based on Category</p>
  </div>

  <div class="text-8xl flex justify-center shop-header">
    <div class="relative">
      <i
        class="pi pi-desktop rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] flex p-12"
      ></i>
      <span class="text-xl mt-5 absolute ml-32 flex justify-center"
        >Electronics</span
      >
    </div>

    <div class="relative">
      <i
        class="pi pi-shopping-bag rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] flex p-12"
      ></i>
      <span class="text-xl mt-5 absolute ml-32 flex justify-center"
        >Shoe / Clothes</span
      >
    </div>

    <div class="relative">
      <i
        class="pi pi-palette rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] flex p-12"
      ></i>
      <span class="text-xl mt-5 absolute ml-32 flex justify-center"
        >Cosmetics</span
      >
    </div>
  </div>

  <template>
  <div
    class="main-container flex w-[1170px] flex-col gap-[60px] items-start flex-nowrap relative mx-auto my-0"
  >
    <div
      class="flex w-[216px] flex-col gap-[20px] items-start shrink-0 flex-nowrap relative"
    >
      <div
        class="flex w-[110px] gap-[16px] items-center shrink-0 flex-nowrap relative z-[1]"
      >
        <div class="w-[20px] h-[40px] shrink-0 relative z-[2]">
          <div
            class="w-full h-full bg-[#db4444] rounded-[4px] absolute top-0 left-0 z-[3]"
          ></div>
        </div>
        <span
          class="h-[20px] shrink-0 basis-auto font-['Poppins'] text-[16px] font-semibold leading-[20px] text-[#db4444] relative text-left whitespace-nowrap z-[4]"
          >Featured</span
        >
      </div>
      <span
        class="h-[48px] shrink-0 basis-auto font-['Inter'] text-[36px] font-semibold leading-[48px] text-[#000] tracking-[1.44px] relative text-left whitespace-nowrap z-[5]"
        >New Arrival</span
      >
    </div>
    <div
      class="flex w-[1170px] gap-[30px] items-start shrink-0 flex-nowrap relative z-[6]"
    >
      <div
        class="w-[570px] h-[600px] shrink-0 bg-[#000] rounded-[4px] relative overflow-hidden z-[7]"
      >
        <div
          class="w-[511px] h-[511px] bg-[url(../assets/images/c04b3497-face-44b8-b2ed-ddd7ba97cafb.png)] bg-cover bg-no-repeat absolute top-[89px] left-[29px] z-[8]"
        ></div>
        <div
          class="flex w-[242px] h-[122px] flex-col gap-[16px] items-start flex-nowrap absolute top-[446px] left-[32px] z-[9]"
        >
          <div
            class="flex w-[242px] flex-col gap-[16px] items-start shrink-0 flex-nowrap relative z-10"
          >
            <span
              class="h-[24px] shrink-0 basis-auto font-['Inter'] text-[24px] font-semibold leading-[24px] text-[#f9f9f9] tracking-[0.72px] relative text-left whitespace-nowrap z-[11]"
              >PlayStation 5</span
            ><span
              class="flex w-[242px] h-[42px] justify-start items-start shrink-0 font-['Poppins'] text-[14px] font-normal leading-[21px] text-[#f9f9f9] relative text-left z-[12]"
              >Black and White version of the PS5 coming out on sale.</span
            >
          </div>
          <div
            class="flex w-[81px] flex-col items-start shrink-0 flex-nowrap relative z-[13]"
          >
            <span
              class="h-[24px] shrink-0 basis-auto font-['Poppins'] text-[16px] font-medium leading-[24px] text-[#fff] relative text-left whitespace-nowrap z-[14]"
              >Shop Now</span
            >
            <div class="w-[81px] shrink-0 opacity-50 relative z-[15]">
              <div
                class="w-full h-[Infinity%] bg-[url(../assets/images/d8c55d7c-132c-486b-add0-8f62fe889b11.png)] bg-[length:100%_100%] bg-no-repeat absolute top-[-Infinity%] left-0 z-[16]"
              ></div>
            </div>
          </div>
        </div>
      </div>
      <div
        class="flex w-[570px] flex-col gap-[32px] items-center shrink-0 flex-nowrap relative z-[17]"
      >
        <div
          class="w-[570px] h-[284px] shrink-0 bg-[#0d0d0d] rounded-[4px] relative overflow-hidden z-[18]"
        >
          <div
            class="w-[432px] h-[286px] bg-[url(../assets/images/15396805-453c-4403-9e91-68923ff61e73.png)] bg-cover bg-no-repeat absolute top-0 left-[138px] z-[19]"
          ></div>
          <div
            class="flex w-[255px] h-[122px] flex-col gap-[16px] items-start flex-nowrap absolute top-[138px] left-[24px] z-20"
          >
            <div
              class="flex w-[255px] flex-col gap-[16px] items-start shrink-0 flex-nowrap relative z-[21]"
            >
              <span
                class="h-[24px] shrink-0 basis-auto font-['Inter'] text-[24px] font-semibold leading-[24px] text-[#f9f9f9] tracking-[0.72px] relative text-left whitespace-nowrap z-[22]"
                >Women’s Collections</span
              ><span
                class="flex w-[255px] h-[42px] justify-start items-start shrink-0 font-['Poppins'] text-[14px] font-normal leading-[21px] text-[#f9f9f9] relative text-left z-[23]"
                >Featured woman collections that give you another vibe.</span
              >
            </div>
            <div
              class="flex w-[81px] flex-col items-start shrink-0 flex-nowrap relative z-[24]"
            >
              <span
                class="h-[24px] shrink-0 basis-auto font-['Poppins'] text-[16px] font-medium leading-[24px] text-[#fff] relative text-left whitespace-nowrap z-[25]"
                >Shop Now</span
              >
              <div class="w-[81px] shrink-0 opacity-50 relative z-[26]">
                <div
                  class="w-full h-[Infinity%] bg-[url(../assets/images/e4e45bec-cc20-4245-9faf-3588688b2995.png)] bg-[length:100%_100%] bg-no-repeat absolute top-[-Infinity%] left-0 z-[27]"
                ></div>
              </div>
            </div>
          </div>
        </div>
        <div
          class="flex w-[570px] gap-[30px] justify-center items-center shrink-0 flex-nowrap relative z-[28]"
        >
          <div
            class="w-[270px] h-[284px] shrink-0 bg-[#000] rounded-[4px] relative overflow-hidden z-[29]"
          >
            <div
              class="w-[210px] h-[222px] absolute top-[31px] left-[30px] overflow-hidden z-[31]"
            >
              <div
                class="w-[190px] h-[221px] bg-[url(../assets/images/fa41799a-9988-414d-bb88-6a7b27ee861e.png)] bg-cover bg-no-repeat relative z-[32] mt-0 mr-0 mb-0 ml-[10px]"
              ></div>
            </div>
            <div
              class="w-[196px] h-[196px] bg-[url(../assets/images/7c313ed3-df51-4260-8176-ae3f9306f30b.png)] bg-cover bg-no-repeat absolute top-1/2 left-1/2 translate-x-[-50%] translate-y-[-50%] z-30"
            ></div>
            <div
              class="flex w-[191px] h-[85px] flex-col gap-[8px] items-start flex-nowrap absolute top-[175px] left-[24px] z-[33]"
            >
              <div
                class="flex w-[191px] flex-col gap-[8px] items-start shrink-0 flex-nowrap relative z-[34]"
              >
                <span
                  class="h-[24px] shrink-0 basis-auto font-['Inter'] text-[24px] font-semibold leading-[24px] text-[#f9f9f9] tracking-[0.72px] relative text-left whitespace-nowrap z-[35]"
                  >Speakers</span
                ><span
                  class="flex w-[191px] h-[21px] justify-start items-start shrink-0 basis-auto font-['Poppins'] text-[14px] font-normal leading-[21px] text-[#f9f9f9] relative text-left whitespace-nowrap z-[36]"
                  >Amazon wireless speakers</span
                >
              </div>
              <div
                class="flex w-[81px] flex-col items-start shrink-0 flex-nowrap relative z-[37]"
              >
                <span
                  class="h-[24px] shrink-0 basis-auto font-['Poppins'] text-[16px] font-medium leading-[24px] text-[#fff] relative text-left whitespace-nowrap z-[38]"
                  >Shop Now</span
                >
                <div class="w-[81px] shrink-0 opacity-50 relative z-[39]">
                  <div
                    class="w-full h-[Infinity%] bg-[url(../assets/images/6ee7cf6a-06df-4d5e-90d9-0894fb9e5829.png)] bg-[length:100%_100%] bg-no-repeat absolute top-[-Infinity%] left-0 z-40"
                  ></div>
                </div>
              </div>
            </div>
          </div>
          <div
            class="w-[270px] h-[284px] shrink-0 bg-[#000] rounded-[4px] relative overflow-hidden z-[41]"
          >
            <div
              class="w-[238px] h-[238px] bg-[url(../assets/images/320d1da3-63b3-4ce8-8545-b9937ee04f86.png)] bg-cover bg-no-repeat absolute top-1/2 left-1/2 translate-x-[-50%] translate-y-[-50%] z-[42]"
            ></div>
            <div
              class="w-[210px] h-[222px] absolute top-1/2 left-1/2 translate-x-[-50%] translate-y-[-50.45%] overflow-hidden z-[43]"
            >
              <div
                class="w-[201px] h-[203px] bg-[url(../assets/images/fbe98866-126c-41b3-9c54-272f13e40b46.png)] bg-cover bg-no-repeat relative z-[44] mt-[8px] mr-0 mb-0 ml-[4px]"
              ></div>
            </div>
            <div
              class="flex w-[191px] h-[85px] flex-col gap-[8px] items-start flex-nowrap absolute top-[175px] left-[24px] z-[45]"
            >
              <div
                class="flex w-[191px] flex-col gap-[8px] items-start shrink-0 flex-nowrap relative z-[46]"
              >
                <span
                  class="h-[24px] shrink-0 basis-auto font-['Inter'] text-[24px] font-semibold leading-[24px] text-[#f9f9f9] tracking-[0.72px] relative text-left whitespace-nowrap z-[47]"
                  >Perfume</span
                ><span
                  class="flex w-[191px] h-[21px] justify-start items-start shrink-0 basis-auto font-['Poppins'] text-[14px] font-normal leading-[21px] text-[#f9f9f9] relative text-left whitespace-nowrap z-[48]"
                  >GUCCI INTENSE OUD EDP</span
                >
              </div>
              <div
                class="flex w-[81px] flex-col items-start shrink-0 flex-nowrap relative z-[49]"
              >
                <span
                  class="h-[24px] shrink-0 basis-auto font-['Poppins'] text-[16px] font-medium leading-[24px] text-[#fff] relative text-left whitespace-nowrap z-50"
                  >Shop Now</span
                >
                <div class="w-[81px] shrink-0 opacity-50 relative z-[51]">
                  <div
                    class="w-full h-[Infinity%] bg-[url(../assets/images/69a67a23-f586-47a9-bdcf-e915a30de735.png)] bg-[length:100%_100%] bg-no-repeat absolute top-[-Infinity%] left-0 z-[52]"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>





  <div class="mt-36 bg-slate-900 text-white flex justify-center h-60 rounded">
    <p class="flex mt-28 text-4xl shop-header">
      Enhance Your Shopping Experience
    </p>
  </div>

  <div class="flex justify-center shop-header">
    <div class="flex flex-col">
      <i
        class="pi pi-truck rounded-full text-4xl mt-9 ml-20 w-[200px] flex p-12"
      ></i>
      <span class="text-xl relative ml-24 -mt-5">Fast Delivery</span>
    </div>

    <div class="flex flex-col">
      <i
        class="pi pi-headphones rounded-full text-4xl mt-9 ml-28 flex p-12"
      ></i>
      <span class="text-xl relative ml-24 -mt-5">24/7 Customer Service</span>
    </div>

    <div class="flex flex-col">
      <i
        class="pi pi-money-bill rounded-full text-4xl mt-9 ml-28 w-[200px] flex p-12"
      ></i>
      <span class="text-xl relative ml-24 -mt-5">Money Back Guarantee</span>
    </div>
  </div>

  <div class="bg-slate-900 text-white shop-header mt-20 shop-header">
    <div class="pt-12 pb-9 flex flex-row gap-56 justify-center">
      <div class="">
        <p class="font-bold text-xl">Accounts</p>
        <ul class="flex gap-6 flex-col mt-6 text-sm">
          <li>Login / Register</li>
          <li>Cart</li>
          <li>Shop</li>
        </ul>
      </div>

      <div>
        <p class="font-bold text-xl">Quick Link</p>
        <ul class="flex gap-6 flex-col mt-6 text-sm">
          <li>Privacy Policy</li>
          <li>Terms Of Use</li>
          <li>Contact</li>
        </ul>
      </div>

      <div>
        <p class="font-bold text-xl">Download App</p>
        <img src="../assets/images/qrcode.png" class="bg-white h-36" alt="" />
      </div>
    </div>
  </div>
</template>

<script>
import Carousels from "../components/Carousels.vue";
// import ShopCard from '../components/ShopCard.vue'
import shop from "../components/Shops.vue";
// import Carousels from "../components/Carousels.vue";
// import shop from "../components/Shops.vue";
import HeroSection from "./HeroSection.vue";

export default {
  data() {
    return {
     pageReloaded: false
    }
  },
  name: 'home',
  components: {
      Carousels,
      shop,
      HeroSection
  }

  // data() {
  //   return {
  //     pageReloaded: false,
  //   };
  // },
  // name: "home",
  // components: {
  //   Carousels,
  //   shop,
  //   HeroSection
  // },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");
.shop-header {
  font-family: "Marcellus", sans-serif;
}
</style>