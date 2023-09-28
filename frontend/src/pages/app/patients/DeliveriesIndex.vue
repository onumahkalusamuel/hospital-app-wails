<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import ActionButton from '@/components/ActionButton.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import TextField from '@/components/form/TextField.vue';
import { Delivery, Pagination, Patient } from '@/interfaces';
import { MagnifyingGlassIcon, TrashIcon } from '@heroicons/vue/24/outline';
import { UserPlusIcon, UserGroupIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref, watch } from 'vue';
import { toasts } from '@/stores/toasts';
import Paging from '@/components/Paging.vue';
import { useDebounce } from '@/utils/debounce';
import { useRoute } from 'vue-router';
const debounce = useDebounce()

const breadcrumbs = ref([
    { title: "Dashboard", link: { name: "dashboard" } },
] as BreadcrumbItem[]);

const route = useRoute();
const patient = ref({} as Patient)
const pagination = ref({
    limit: 10,
    page: 1,
    rows: [] as Delivery[],
    sort_key: '',
    sort_order: 'desc',
    total_pages: 0,
    total_rows: 0,
    query: ''
} as Pagination);

const deletItem = async (id: string) => {
    const del = await apiRequest.deleteRecord(`deliveries/${id}`);
    if (del && del.message)
        toasts.addToast({ message: del.message, type: 'success' });
    fetch();
}
const fetch = () => {
    debounce(async () => {
        const filter = { ...pagination.value, rows: null, patient_id: route.params.id };
        pagination.value = await apiRequest.get(`deliveries?${new URLSearchParams(filter as never as Record<string, string>)}`);
    })
}

const fetchPatient = async () => {
    patient.value = await apiRequest.get(`patients/${route.params.id}`);
    breadcrumbs.value.push(
        { title: `${patient.value.lastname} ${patient.value.firstname}`, link: { name: "view-patient", params: { id: route.params.id } } },
        { title: "Deliveries", current: true }
    )
}


onMounted(async () => { fetch(), fetchPatient() });
watch(() => pagination.value.limit, fetch);
watch(() => pagination.value.page, fetch);
watch(() => pagination.value.query, fetch);
</script>;

<template>
    <div class="page-wrapper">
        <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
        <PageHeader title="Deliveries" :subtitle="`${patient.lastname} ${patient.firstname} deliveries`" :icon-src="UserGroupIcon" />
        <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
            <div class="pr-2">
                <ActionButton dark
                    @click="() => $router.push({ name: 'add-delivery', params: { patient_id: route.params.id } })"
                    :icon-src="UserPlusIcon">Add delivery
                </ActionButton>
            </div>
            <div class="flex">
                <TextField placeholder="Search" v-model="pagination.query">
                    <template #prepend>
                        <MagnifyingGlassIcon class="h-5 w-5" />
                    </template>
                </TextField>
            </div>
        </div>
        <hr class="ml-4 mr-4" />
        <div class="page-scroll-area">
            <div v-if="!pagination.rows.length">
                <div class="text-center mt-4 mb-4 pt-4">No records found</div>
                <hr />
            </div>
            <div v-else class="table">
                <div class="table-row table-header">
                    <div class="cell-data cell-sn">S/N</div>
                    <div class="cell-data cell-size-1">Patient Name</div>
                    <div class="cell-data">Delivery Mode</div>
                    <div class="cell-data">Sex</div>
                    <div class="cell-data">Baby Weight</div>
                    <div class="cell-data cell-size-1">Condition.</div>
                    <div class="cell-data cell-size-1">Delivery Time</div>
                    <div class="cell-data">Actions</div>
                </div>
                <div class="table-body">
                    <div class="table-row" v-for="delivery, i in pagination.rows" :key="delivery.id">
                        <div class="cell-data cell-sn">{{ i + 1 }}</div>
                        <div class="cell-data cell-size-1">
                            <router-link class="text-blue-600 hover:underline"
                                :to="{ name: 'view-delivery', params: { id: delivery.id } }">
                                {{ `${delivery.patient?.firstname} ${delivery.patient?.lastname}` }}
                            </router-link>
                        </div>
                        <div class="cell-data">{{ delivery.delivery_mode }}</div>
                        <div class="cell-data">{{ delivery.baby_sex }}</div>
                        <div class="cell-data">{{ delivery.baby_weight }} kg</div>
                        <div class="cell-data cell-size-1">{{ delivery.condition }}</div>
                        <div class="cell-data cell-size-1">{{ delivery.delivery_date_time }}</div>
                        <div class="cell-data">
                            <ActionButton outline v-on:click="deletItem(delivery.id)" :icon-src="TrashIcon">Delete</ActionButton>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <Paging v-model="pagination" />
    </div>
</template>