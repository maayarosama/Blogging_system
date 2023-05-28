<template>
    <div>

        <v-text-field :rules="rules" :type="type" :label="label" :placeholder="placeholder" v-model="value" class="mb-2" clearable bg-color="accent" variant="outlined"
            density="compact" @update:model-value="UpdateInput()"></v-text-field>

    </div>
</template>

<script setup>
import {ref, defineEmits,toRefs} from "vue"
const value = ref("")
const emit = defineEmits(["emit-update-input"])


const UpdateInput=()=> {
    emit("emit-update-input",value.value ,label.value )

}

const props = defineProps({
  label: String,
  type: String,
  placeholder: String,
  rules:Object,
})
const { label,type,placeholder ,rules} = toRefs(props)
// const emit = defineEmits(['eventA', 'eventB'])
// function btnClick(params) {
//     emit('eventA')
//     emit('eventB', params)
// }

</script>

<style lang="scss" scoped></style>