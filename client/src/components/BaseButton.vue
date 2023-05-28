<template>
    <div>
        <v-btn  :color="color" v-bind="$attrs">
            <font-awesome-icon v-if="icon" :icon="icon" class="mr-2" />
            {{ text }}</v-btn>
    </div>
</template>

<script setup>
import { toRefs } from "vue"

const props = defineProps({
    type: {
        type: String,
        default: "submit",
    },
    color: {
        type: String,
        default: "",
    },
    icon: {
        type: String,
        default: "",
    },
    text: {
        type: [String, Number],
        default: "",
    },
})
const { color, icon, text } = toRefs(props)
</script>

<style lang="scss" scoped></style>