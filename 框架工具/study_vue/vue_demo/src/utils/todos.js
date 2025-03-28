import { ref, computed } from "vue"
import useStorage from "@/utils/storage"

function useTodos() {

    let title = ref("")
    // let todos = ref([{ title: "学习", done: false }])
    let todos = useStorage("todos", [])

    function addTodo() {
        todos.value.push({
            title: title.value,
            done: false,
        })
        title.value = ""
    }

    function clear() {
        todos.value = todos.value.filter((v) => !v.done)
    }

    let active = computed(() => {
        return todos.value.filter((v) => !v.done).length
    })

    let all = computed(() => todos.value.length)

    let allDone = computed({
        get() {
            return active.value === 0
        },
        set(val) {
            todos.value.forEach((v) => (v.done = val))
        },
    })

    return {
        title,
        todos,
        addTodo,
        clear,
        active,
        all,
        allDone,
    }
}

export default useTodos