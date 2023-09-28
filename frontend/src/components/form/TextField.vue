<script setup lang="ts">

defineProps<{
  ref?: string;
  name?: string;
  placeholder?: string;
  type?: string;
  required?: boolean;
  class?: string;
  label?: string;
  value?: string | number;
  readonly?: boolean;
  disabled?: boolean;
  step?: string;
  max?: string | number;
  min?: string | number;
  modelValue?: string | number;
  accept?: string;
}>()

const emit = defineEmits(['update:modelValue'])

const updateValue = (event: any) => {
  emit('update:modelValue', event.target.value)
}

</script>
<template>
  
    <label v-if="label" class="text-uppercase cursor-pointer py-2 font-bold block" :for="name">
      {{ label }}
      <span v-if="required" class="text-red-600">*</span>
    </label>
    <div class="flex w-full border-[1px] border-[#888] rounded flex items-center justify-center bg-white hover:bg-gray-100"
      :class="class">
      <div v-if="$slots.prepend" class="px-2 flex items-center justify-center hover:bg-gray-100 rounded">
        <slot name="prepend"></slot>
      </div>
      <input :class="class" class="h-[40px] w-full rounded py-1 px-2 hover:bg-gray-100 focus:outline-none" :name="name"
        :placeholder="placeholder" :type="type || 'text'" :required="required" :autocomplete="`new-${name}`" :id="name"
        :value="modelValue || value" :readonly="readonly" :step="step" :ref="ref" @input="updateValue" :disabled="disabled"
        :max="max" :min="min" :accept="accept" />
      <div v-if="$slots.append" class="px-2 flex items-center justify-center hover:bg-gray-100 rounded">
        <slot name="append"></slot>
      </div>
    </div>
</template>