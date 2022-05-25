<script setup>
import { ref } from 'vue'

defineProps({
  msg: String
})

const count = ref(0)
</script>

<template>
  <h1>{{ msg }}</h1>

  <div id="app">
    <div class="app-wrapper">
      <header class="app-header">
        <h1 class="app-title"> All Task</h1>
        <ul class="todo-list" v-for="todo in todos">
          <li v-on:click="removeTodo($index)" class="todo-item"
              v-bind:class="{'todo-item--removed':todo.removed}">{{todo.txt}}} asd
            <span class="todo-delete"></span></li>
        </ul>
        <div class="app-form">
          <input>
        </div>
      </header>
    </div>
  </div>
</template>
<style scoped>
a {
  color: #42b983;
}
</style>
