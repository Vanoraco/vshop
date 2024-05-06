<template>
    <div>
      <v-container>
        <v-card class="add-item-card">
          <v-card-title class="add-item-title">
            <h2>Add Item</h2>
          </v-card-title>
          <v-card-text>
            <v-form @submit.prevent="addItem">
              <v-text-field v-model="item.name" label="Name" required></v-text-field>
              <v-text-field v-model="item.price" label="Price" required></v-text-field>
              <v-textarea v-model="item.description" label="Description" required></v-textarea>
              <v-file-input v-model="item.images" label="Images" multiple @change="handleImageChange"></v-file-input>
              <v-row>
                <v-col v-for="(image, index) in item.imagesPreview" :key="index" cols="6">
                  <img :src="image" alt="Preview" class="preview-image" />
                </v-col>
              </v-row>
              <v-text-field v-model="item.phoneNumber" label="Phone Number" required></v-text-field>
              <v-btn type="submit" color="primary">Add</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
  
        <v-card v-if="addedItems.length > 0" class="added-items-card">
          <v-card-title class="added-items-title">
            <h2>Added Items</h2>
          </v-card-title>
          <v-card-text>
            <v-list>
              <v-list-item v-for="(item, index) in addedItems" :key="index" class="added-item">
                <v-list-item-content>
                  <v-list-item-title>{{ item.name }}</v-list-item-title>
                  <v-list-item-subtitle>{{ item.price }}</v-list-item-subtitle>
                  <v-list-item-subtitle>{{ item.description }}</v-list-item-subtitle>
                  <v-list-item-subtitle>{{ item.phoneNumber }}</v-list-item-subtitle>
                </v-list-item-content>
                <v-list-item-action>
                  <v-btn icon @click="removeItem(index)">
                    <v-icon>mdi-delete</v-icon>
                  </v-btn>
                </v-list-item-action>
              </v-list-item>
            </v-list>
          </v-card-text>
        </v-card>
      </v-container>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  
  export default {
    name: 'AddItem',
    setup() {
      const item = ref({
        name: '',
        price: '',
        description: '',
        images: [],
        phoneNumber: ''
      });
  
      const addedItems = ref([]);
  
      const handleImageChange = (event) => {
        const files = event.target.files;
        const images = [];
        const imagesPreview = [];
  
        for (let i = 0; i < files.length; i++) {
          const reader = new FileReader();
          reader.onload = (e) => {
            imagesPreview.push(e.target.result);
          };
          reader.readAsDataURL(files[i]);
          images.push(files[i]);
        }
  
        item.value.images = images;
        item.value.imagesPreview = imagesPreview;
      };
  
      const addItem = () => {
        // Here you can perform any logic to save the item, like sending it to an API
        addedItems.value.push({ ...item.value });
        console.log(item.value);
        // Reset the form
        item.value = {
          name: '',
          price: '',
          description: '',
          images: [],
          phoneNumber: ''
        };
      };
  
      const removeItem = (index) => {
        addedItems.value.splice(index, 1);
      };
  
      return {
        item,
        addedItems,
        handleImageChange,
        addItem,
        removeItem
      };
    }
  };
  </script>
  
  <style scoped>
  .container {
    max-width: 500px;
    margin: 0 auto;
  }
  
  .add-item-card {
    padding: 20px;
    margin-top: 20px;
    background-color: #f5f5f5;
    border-radius: 8px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
  }
  
  .add-item-title {
    margin-bottom: 20px;
    text-align: center;
  }
  
  .preview-image {
    width: 100%;
    height: auto;
    margin-bottom: 10px;
    border-radius: 8px;
  }
  
  .added-items-card {
    margin-top: 40px;
    background-color: #f5f5f5;
    border-radius: 8px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
  }
  
  .added-items-title {
    margin-bottom: 20px;
    text-align: center;
  }
  
  .added-item {
    margin-bottom: 10px;
    padding: 10px;
    background-color: #ffffff;
    border-radius: 8px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
  }
  
  .added-item:hover {
    background-color: #f5f5f5;
  }
  
  .v-btn {
    margin-top: 20px;
  }
  
  /* Custom CSS for responsiveness */
  
  @media (max-width: 600px) {
    .container {
      max-width: 100%;
      padding: 0 20px;
    }
  }
  </style>