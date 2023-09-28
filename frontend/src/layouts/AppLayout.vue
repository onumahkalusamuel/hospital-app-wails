<script lang="ts" setup>
import { RouterView, RouterLink } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import { user, auth, hospital } from '@/stores';
import { UserIcon, SparklesIcon, Cog8ToothIcon, LockClosedIcon, QueueListIcon, BanknotesIcon, UserGroupIcon, UsersIcon } from '@heroicons/vue/24/solid';
import { onMounted } from 'vue';

onMounted(async () => {
  try {
    // get hospital
    const req = await apiRequest.get("hospital-details");
    if (req)
      hospital.setAll(req);
    // get user
    const profile = await apiRequest.get("profile");
    if (profile)
      user.setAll(profile);
  }
  catch (e) {
    console.log(e);
  }
})

</script>;

<template>
  <div class="flex flex-col h-screen">
    <header class="w-full flex toolbar bg-[#0078d4] h-[50px]">
      <div class="flex items-center text-white">
        <router-link class="header-icon-link px-3 hover:bg-[#1664a7]" :to="{ name: 'dashboard' }">
          <SparklesIcon class="text-white h-6 w-6" />
          <span class="pl-2 hidden lg:block">{{ hospital.get('hospital_name') }}</span>
        </router-link>
      </div>
      <div class="flex m-r-[15px] items-center flex-1 justify-center uppercase">

        <router-link title="Staff" :class="$route.name == 'staff' ? 'bg-[#00000033]' : ''" :to="{ name: 'staff' }"
          class="header-icon-link px-5" v-if="user.role == '1'">
          <UserIcon class="text-white h-5 w-5" />
          <span class="pl-2 hidden lg:block">Staff</span>
        </router-link>

        <router-link title="Patients" :class="$route.name == 'patients' ? 'bg-[#00000033]' : ''" :to="{ name: 'patients' }"
          class="header-icon-link px-5">
          <UsersIcon class="text-white h-5 w-5" />
          <span class="pl-2 hidden lg:block">Patients</span>
        </router-link>

        <router-link title="Deliveries" :class="$route.name == 'deliveries' ? 'bg-[#00000033]' : ''"
          :to="{ name: 'deliveries' }" class="header-icon-link px-5">
          <UserGroupIcon class="text-white h-5 w-5" />
          <span class="pl-2 hidden lg:block">Deliveries</span>
        </router-link>

        <router-link title="Billings" :class="$route.name == 'billings' ? 'bg-[#00000033]' : ''" :to="{ name: 'billings' }"
          class="header-icon-link px-5">
          <BanknotesIcon class="text-white h-5 w-5" />
          <span class="pl-2 hidden lg:block">Billings</span>
        </router-link>

        <router-link title="Reports" :class="$route.name == 'reports' ? 'bg-[#00000033]' : ''" :to="{ name: 'reports' }"
          class="header-icon-link px-5">
          <QueueListIcon class="text-white h-5 w-5" />
          <span class="pl-2 hidden lg:block">Reports</span>
        </router-link>

      </div>
      <div style="display:flex; flex: 0 0 auto">
        <router-link class="header-icon-link w-[48px]" :to="{ name: 'settings' }">
          <cog8-tooth-icon class="text-white h-5 w-5" />
        </router-link>
      </div>
      <a title="Logout" class="header-icon-link px-3"
        @click="() => { auth.setJwt(''); user.reset(); $router.push({ name: 'login' }); }">
        <lock-closed-icon class="text-white h-5 w-5 mr-2" />
        <span class="hidden lg:block">Logout</span>
      </a>
    </header>
    <main class="flex-1 flex-scroll">
      <router-view></router-view>
    </main>
  </div>
</template>

<style scoped>
.toolbar {
  flex: 0 0 auto;
  box-sizing: border-box;
}

.header-icon-link {
  height: 50px;
  display: flex;
  flex: 0 0 auto;
  align-items: center;
  justify-content: center;
  color: white;
  text-decoration: none;
  cursor: pointer;
}

.header-icon-link:hover {
  background-color: #1664a7;
}</style>