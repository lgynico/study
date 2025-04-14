<template>
    <div>
        <h1 @click="add">{{ count }}</h1>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const count = ref(0)
const color = ref("red")

function add() {
    count.value++
    color.value = Math.random() > 0.5 ? "red" : "blue"
}
</script>

<style scoped>
h1 {
    color: v-bind(color);
}
</style>