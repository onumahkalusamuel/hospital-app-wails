<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';

const props = defineProps<{
    name?: string;
    options?: string[];
    required?: boolean;
    label?: string;
    modelValue?: string;
}>()

const selectedValue = ref('')

const emit = defineEmits(['update:modelValue'])

const makeChoice = (choice: string) => {
    selectedValue.value = choice
    emit('update:modelValue', choice)
}

onMounted(() => {
    if (props.modelValue !== undefined) selectedValue.value = props.modelValue
})

watch(() => props.modelValue, (newValue, oldValue) => {
    if (!oldValue && newValue !== undefined) selectedValue.value = newValue
})
</script>
<template>
    <div class="">
        <label v-if="label" class="cursor-pointer py-2 font-bold block" :for="name">
            {{ label }}
            <span v-if="required" class="text-red-600">*</span>
        </label>

        <div class="w-full flex flex-wrap gap-2">
            <div v-for="(choice, i) in options" :key="i"
                class="cursor-pointer flex items-center px-2 border-[1px] rounded h-[40px]"
                :class="selectedValue == choice ? 'bg-blue-600 border-blue-600 text-white' : 'border-gray-400'"
                @click="() => makeChoice(choice)"> {{ choice }} </div>
        </div>
        <input type="hidden" v-model="selectedValue" :name="name" />
    </div>
</template>