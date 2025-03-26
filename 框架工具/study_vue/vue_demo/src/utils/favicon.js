import { ref, watch } from "vue"

export default function useFavicon(newIcon) {
    const favicon = ref(newIcon)

    const updateIcon = (icon) => {
        document.head.querySelectorAll("[rel*='icon']").forEach((el) => (el.href = `${icon}`))
    }

    const reset = () => favicon.value = "/favicon.ico"

    watch(favicon, updateIcon)

    return { favicon, reset }
}