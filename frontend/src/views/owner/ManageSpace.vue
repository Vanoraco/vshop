<template>
  <div class="manage-bar">
    <ul class="menu-list">
      <li class="menu-item">Add new product</li>
      <li class="menu-item">Pay monthly salary</li>
      <li class="menu-item">Requests</li>
      <li class="menu-item">Analysis</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Navbar'
};
</script>

<style scoped>
.toolbar-title {
  font-size: 24px;
  font-weight: bold;
}

.manage-bar {
  height: 50px;
  display: flex;
  justify-content: flex-end;
  background-color: #f2f2f2;
  border-bottom: 1px solid #ccc;
}

.menu-list {
  min-width: 200px;
  list-style-type: none;
  display: flex;
  align-items: center;
  padding: 0;
  margin: 0;
}

.menu-item {
  font-size: 16px;
  padding: 10px;
  color: #333;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.menu-item:hover {
  background-color: #e6e6e6;
}

.menu-divider {
  margin: 16px 0;
}

.menu-icon {
  margin-right: 8px;
}
</style>