<script setup lang="ts">
import { ref, watch } from 'vue';
import apiRequest from '@/services/http/api-requests';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import { Patient, Pagination } from '@/interfaces'

defineProps<{ modelValue: Patient }>()
const emit = defineEmits(['update:modelValue', 'update:selected'])

const updateValue = (p: Patient) => {
    selected.value = p;
    emit('update:modelValue', p)
}
const selected = ref({} as Patient);

const pagination = ref({
    limit: 5,
    page: 1,
    rows: [] as Patient[],
    sort_key: '',
    sort_order: 'desc',
    total_pages: 0,
    total_rows: 0,
    query: ''
} as Pagination);

const fetchPatients = async () => {
    const filter = { ...pagination.value, rows: null };
    pagination.value = await apiRequest.get(`patients?${new URLSearchParams(filter as never as Record<string, string>)}`);
}
const patientSelected = () => {
    // pagination.value.rows = [];
    // pagination.value.query = '';
    emit('update:selected', true);
}

watch(() => pagination.value.query, async (y, x) => {
    if (y.length < 2 || y == x) return;
    await fetchPatients();
});

</script>
<template>
    <div class="space-y-5">
        <div class="w-full">
            <input
                class="focus:outline-none focus:bg-gray-100 border-[1px] px-2 py-2 rounded w-full text-center border-blue-600"
                label="Patient Name or Card No" placeholder=" Enter part of name or Card No" v-model="pagination.query">
        </div>
        <div class="w-full space-y-2">
            <div @click="() => updateValue(p)" v-for="p in pagination.rows"
                class="cursor-pointer rounded flex justify-between px-2 py-1 border-[1px] border-blue-600 items-center"
                :key="p.id" :class="selected.id == p.id ? 'bg-blue-600 text-white' : ''">
                <div>{{ p.lastname }} {{ p.firstname }} - {{ p.card_no }}</div>
            </div>
        </div>
    </div>
    <div class="flex gap-3 mt-5" v-if="pagination.rows.length">
        <div>
            <PrimaryButton @click.prevent="patientSelected">Select Patient</PrimaryButton>
        </div>
    </div>
</template>