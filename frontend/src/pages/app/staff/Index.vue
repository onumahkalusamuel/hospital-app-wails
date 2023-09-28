<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import ActionButton from '@/components/ActionButton.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import { Staff, Roles, Pagination } from '@/interfaces'
import TextField from '@/components/form/TextField.vue';
import { UserPlusIcon, UserIcon, MagnifyingGlassIcon, TrashIcon, PencilIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref, watch } from 'vue';
import { toasts } from '@/stores/toasts';
import Paging from '@/components/Paging.vue';
import { useDebounce } from '@/utils/debounce';
import { popupStore } from '@/stores';
import StaffViewPopup from './popups/StaffViewPopup.vue'
const debounce = useDebounce()

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Staff", current: true },
] as BreadcrumbItem[]);

const pagination = ref({
  limit: 10,
  page: 1,
  rows: [] as Staff[],
  sort_key: '',
  sort_order: 'desc',
  total_pages: 0,
  total_rows: 0,
  query: ''
} as Pagination);

const deleteItem = async (id: string) => {
  const del = await apiRequest.deleteRecord(`staff/${id}`);
  if (del && del.message) toasts.addToast({ message: del.message, type: 'success' });
  fetchStaff();
}

const fetchStaff = async () => {
  debounce(async () => {
    const filter = { ...pagination.value, rows: null };
    pagination.value = await apiRequest.get(`staff?${new URLSearchParams(filter as never as Record<string, string>)}`);
  });
}
const currentRecord = ref({} as Staff)
const staffPopupId = ref('staffview');

const viewRecord = (staff: Staff) => {
  currentRecord.value = staff;
  popupStore.id = staffPopupId.value;
  popupStore.show = true;
}

onMounted(async () => { fetchStaff() })
watch(() => pagination.value.limit, fetchStaff);
watch(() => pagination.value.page, fetchStaff);
watch(() => pagination.value.query, fetchStaff);
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Staff" subtitle="Manage staff" :icon-src="UserIcon">
    </PageHeader>
    <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
      <div class="pr-2">
        <ActionButton dark @click="() => $router.push({ name: 'add-staff' })" :icon-src="UserPlusIcon">Add staff
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
          <div class="cell-data cell-size-1">Name</div>
          <div class="cell-data">Role</div>
          <div class="cell-data">Username</div>
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="st, i in pagination.rows" :key="st.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">
              <router-link class="text-blue-600 hover:underline" to="#" @click="viewRecord(st)">
                {{ `${st.firstname} ${st.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ Roles[st.role as 1 | 2 | 3] }}</div>
            <div class="cell-data">{{ st.username }}</div>
            <div class="cell-data gap-x-2">
              <ActionButton @click="() => $router.push({ name: 'view-staff', params: { id: st.id } })"
                :icon-src="PencilIcon" outline>Edit</ActionButton>
              <ActionButton outline v-if="st.role > 1" @click="deleteItem(st.id)" :icon-src="TrashIcon">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Paging v-model="pagination" />
    <StaffViewPopup :popup-id="staffPopupId" :staff="currentRecord" />
  </div>
</template>