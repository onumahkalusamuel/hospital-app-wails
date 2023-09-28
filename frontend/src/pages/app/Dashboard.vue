<script lang="ts" setup>
import { user } from '@/stores/user';
import { UsersIcon, UserGroupIcon, BanknotesIcon, QueueListIcon, UserIcon } from '@heroicons/vue/24/solid';
import { ref } from 'vue';

const links = ref([
  { label: 'Patients', icon: UsersIcon, url: { name: 'patients' } },
  { label: 'Deliveries', icon: UserGroupIcon, url: { name: 'deliveries' } },
  { label: 'Billings', icon: BanknotesIcon, url: { name: 'billings' } },
  { label: 'Reports', icon: QueueListIcon, url: { name: 'reports' } },
]);

</script>;

<template>
  <div class="p-[15px]">
    <div class="pt-[50px] pb-[20px] text-xl">Welcome, {{ user.role > '1' ? (user.role == '2' ? 'Dr.' : 'Nrs.') : '' }} {{
      user.lastname }} {{ user.firstname }}</div>
    <hr />
    <div class="flex flex-wrap justify-around gap-2 m-auto my-5">
      <router-link v-if="user.role == '1'" class="flex flex-col justify-center items-center text-white rounded bg-[#0078d4] hover:bg-blue-700 dashboard-body-item" :to="{ name: 'staff' }">
        <div class="h-[156px] w-[156px]">
          <component :is="UserIcon" />
        </div>
        <div class="mt-5 text-2xl font-bold">Staff</div>
      </router-link>
      <router-link v-for="link, i in links" :key="i" class="flex flex-col justify-center items-center text-white rounded bg-[#0078d4] hover:bg-blue-700 dashboard-body-item" :to="link.url">
        <div class="h-[156px] w-[156px]">
          <component :is="link.icon" />
        </div>
        <div class="mt-5 text-2xl font-bold">{{ link.label }}</div>
      </router-link>
    </div>
  </div>
</template>
<style scoped>
.dashboard-body-item {
  width: 212px;
  height: 259px;
  margin-top: 47px;
  cursor: pointer;
  text-decoration: none;
}
</style>