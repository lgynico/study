<template>
    <h1>这是首页 {{ x }},{{ y }}</h1>
    <button @click="loading">use favicon</button>
    <button @click="toggle">切换全屏</button>
    <Counter />
    <Todolist />
    <Rate :value="score" @update-rate="updateRate" theme="orange"></Rate>
    <Rate :value="3">
        <img width="14" src="/favicon.ico">
    </Rate>
</template>

<script setup>
import { ref } from 'vue';
import Todolist from '@/components/Todolist.vue';
import Counter from '@/components/Counter.vue';
import { useMouse } from '@/utils/mouse';
import useFavicon from '@/utils/favicon';
import { useFullscreen } from '@vueuse/core';
import Rate from '@/components/Rate.vue';

let { x, y } = useMouse();

let { favicon } = useFavicon();

function loading() {
    favicon.value = "./geek.webp"
}

const { toggle } = useFullscreen();

let score = ref(4)

function updateRate(val) {
    score.value = val
}
</script>