<template>
    <div>
      <div class="event-card">
        <h1>Setting and privacy</h1>
      </div>
      <div class="event-card">
        <h1>Setting and privacy</h1>
        <p>change your shipping address</p>
      </div>
      <div class="event-card">
        <h1>Setting and privacy</h1>
        <p>change your shipping address</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
   data(){

   }
  }
  </script>
  
  <style scoped>
  .event-card {
    padding: 20px;
    width: 250px;
    cursor: pointer;
    border: 1px solid #39495c;
    margin-bottom: 18px;
  }
  .event-card:hover {
    transform: scale(1.01);
    box-shadow: 0 3px 12px 0 rgba(0, 0, 0, 0.2);
  }
  h1 {
    font-weight: 500;
    font-size: 2.6rem;
    position: relative;
    top: -10px;
  }
  h3 {
    font-size: 1.2rem;
  }
  .greetings h1,
  .greetings h3 {
    text-align: center;
  }
  @media (min-width: 1024px) {
    .greetings h1,
    .greetings h3 {
      text-align: left;
    }
  }
  </style>