<template>
    <div class="flex items-center flex-col justify-center w-full h-full" :style="{ perspective: '1000px' }">
      <div v-for="(image, index) in images" :key="index" class="rounded-full">
        <motion-img
          :src="image"
          :alt="image"
          class="rounded-[8px] object-fill top-[5%] left-[30%]"
          initial="center"
          :transition="{ duration: 0.7 }"
          :animate="positions[positionindexs[index]]"
          :variants="imageVariants"
          :style="{ width: '40%', position: 'absolute', transform: 'rotate(90deg)' }"
        />
        <div class="flex justify-between items-center z-10 fixed top-[33%] left-[24%] w-[52%]">
          <img src="@/assets/back.png" alt="back" class="z-10 w-26 h-28 cursor-pointer" @click="handlePrev" />
          <img src="@/assets/forward.png" alt="forward" class="z-10 w-28 h-28 cursor-pointer" @click="handleNext" />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { motion } from 'framer-motion';
  import videopic from '@/assets/banner.jpg';
  import image2 from '@/assets/image2.jpg';
  import image3 from '@/assets/image3.jpg';
  import image4 from '@/assets/image4.jpg';
  import image5 from '@/assets/image5.jpg';
  
  export default {
    components: {
      motionImg: motion.img,
    },
    data() {
      return {
        positionindexs: [0, 1, 2, 3, 4],
        images: [videopic, image2, image3, image4, image5],
        positions: ['center', 'left1', 'left', 'right', 'right1'],
        imageVariants: {
          center: { x: '0%', scale: 1.1, zIndex: 5 },
          left1: { x: '-55%', scale: 0.8, zIndex: 2 },
          left: { x: '-95%', scale: 0.6, zIndex: 1 },
          right: { x: '95%', scale: 0.6, zIndex: 1 },
          right1: { x: '55%', scale: 0.8, zIndex: 2 },
        },
      };
    },
    methods: {
      handleNext() {
        this.positionindexs = this.positionindexs.map((index) => (index + 1) % 5);
      },
      handlePrev() {
        this.positionindexs = this.positionindexs.map((index) => (index - 1 + 5) % 5);
      },
    },
  };
  </script>
  
  <style scoped>
  </style>
  