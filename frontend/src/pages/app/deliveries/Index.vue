<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import ActionButton from '@/components/ActionButton.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import { Delivery, Pagination } from '@/interfaces';
import { TrashIcon } from '@heroicons/vue/24/outline';
import { UserPlusIcon, UserGroupIcon, PencilIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref, watch } from 'vue';
import { toasts } from '@/stores/toasts';
import Paging from '@/components/Paging.vue';
import { useDebounce } from '@/utils/debounce';
import DeliveryViewPopup from './popups/DeliveryViewPopup.vue';
import { popupStore } from '@/stores';
import CheckboxField from '@/components/form/CheckboxField.vue';
const debounce = useDebounce()

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Deliveries", current: true },
] as BreadcrumbItem[]);

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

const filters = ref({ delivery_mode: '', baby_sex: '', condition: '' });
const deletItem = async (id: string) => {
  const del = await apiRequest.deleteRecord(`deliveries/${id}`);
  if (del && del.message)
    toasts.addToast({ message: del.message, type: 'success' });
  fetch();
}
const fetch = () => {
  debounce(async () => {
    const filter = { ...pagination.value, rows: null, ...filters.value };
    pagination.value = await apiRequest.get(`deliveries?${new URLSearchParams(filter as never as Record<string, string>)}`);
  })
}

const currentRecord = ref({} as Delivery)
const deliveryPopupId = ref('deliveryview');
const viewRecord = (delivery: Delivery) => {
  currentRecord.value = delivery;
  popupStore.id = deliveryPopupId.value;
  popupStore.show = true;
}

onMounted(async () => { fetch() });
watch(() => pagination.value.limit, fetch);
watch(() => pagination.value.page, fetch);
watch(() => pagination.value.query, fetch);
watch(() => filters.value.baby_sex, fetch);
watch(() => filters.value.condition, fetch);
watch(() => filters.value.delivery_mode, fetch);
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Deliveries" subtitle="Manage deliveries" :icon-src="UserGroupIcon" />
    <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
      <div class="pr-2">
        <ActionButton dark @click="() => $router.push({ name: 'add-delivery' })" :icon-src="UserPlusIcon">Add delivery
        </ActionButton>
      </div>
      <div class="flex gap-x-2 items-center">
        <div class="font-bold">Filters:</div>
        <div class="flex gap-x-5">
          <CheckboxField v-model="filters.delivery_mode" :options="['Vaginal', 'C-section']" />
          <CheckboxField v-model="filters.baby_sex" :options="['Male', 'Female']" />
          <CheckboxField v-model="filters.condition" :options="['Baby cried', 'Baby did not cry']" />
        </div>
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
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="delivery, i in pagination.rows" :key="delivery.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">
              <router-link class="text-blue-600 hover:underline" to="#" @click="viewRecord(delivery)">
                {{ `${delivery.patient?.firstname} ${delivery.patient?.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ delivery.delivery_mode }}</div>
            <div class="cell-data">{{ delivery.baby_sex }}</div>
            <div class="cell-data">{{ delivery.baby_weight }} kg</div>
            <div class="cell-data cell-size-1">{{ delivery.condition }}</div>
            <div class="cell-data gap-x-2">
              <ActionButton outline @click="() => $router.push({ name: 'view-delivery', params: { id: delivery.id } })"
                :icon-src="PencilIcon">Edit</ActionButton>
              <ActionButton outline v-on:click="deletItem(delivery.id)" :icon-src="TrashIcon">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Paging v-model="pagination" />
    <DeliveryViewPopup :popup-id="deliveryPopupId" :delivery="currentRecord" />
  </div>
</template>