import { ref, watchEffect } from "vue"

export default function useStorage(key, value = []) {
    const data = ref(JSON.parse(localStorage.getItem(key) || value))

    watchEffect(() => localStorage.setItem(key, JSON.stringify(data.value)))

    return data
}