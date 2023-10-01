<script lang="ts" setup>
import { RouterView, RouterLink } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import { user, auth, hospital, popupStore } from '@/stores';
import { UserIcon, SparklesIcon, Cog8ToothIcon, LockClosedIcon, QueueListIcon, BanknotesIcon, UserGroupIcon, UsersIcon, QrCodeIcon, Bars3Icon, XMarkIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';
import QrCodePopup from '@/components/popups/QrCodePopup.vue';

const qrCodePopupId = ref('qrcode');
const showQrCodePopup = () => { popupStore.id = qrCodePopupId.value; popupStore.show = true; }
const remoteAddress = ref('');
const menuActive = ref(false);

const menuItems = ref([
  { name: 'patients', icon: UsersIcon },
  { name: 'deliveries', icon: UserGroupIcon },
  { name: 'billings', icon: BanknotesIcon },
  { name: 'reports', icon: UsersIcon },
  { name: 'settings', icon: Cog8ToothIcon },
]);
const toggleMenu = () => (menuActive.value = !menuActive.value)

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
    // for qrcode
    const remote = await apiRequest.get("get-remote-address");
    if (remote.address) remoteAddress.value = remote.address;
  }
  catch (e) {
    console.log(e);
  }
})

</script>;

<template>
  <div class="flex flex-col h-screen transition-all">
    <header class="w-full flex toolbar justify-between bg-[#0078d4] h-[50px]">
      <div class="flex items-center text-white">
        <router-link class="header-icon-link px-3 hover:bg-[#1664a7]" :to="{ name: 'dashboard' }">
          <SparklesIcon class="text-white h-6 w-6" />
          <span class="pl-2">{{ hospital.get('hospital_name') }}</span>
        </router-link>
      </div>
      <div class="lg:flex hidden m-r-[15px] items-center flex-1 justify-center uppercase">
        <router-link title="Staff" :class="$route.name == 'staff' ? 'bg-[#00000033]' : ''" :to="{ name: 'staff' }"
          class="header-icon-link px-5" v-if="user.role == '1'">
          <UserIcon class="text-white h-5 w-5" />
          <span class="pl-2">Staff</span>
        </router-link>

        <router-link title="Patients" :class="$route.name == 'patients' ? 'bg-[#00000033]' : ''"
          :to="{ name: 'patients' }" class="header-icon-link px-5">
          <UsersIcon class="text-white h-5 w-5" />
          <span class="pl-2">Patients</span>
        </router-link>

        <router-link title="Deliveries" :class="$route.name == 'deliveries' ? 'bg-[#00000033]' : ''"
          :to="{ name: 'deliveries' }" class="header-icon-link px-5">
          <UserGroupIcon class="text-white h-5 w-5" />
          <span class="pl-2">Deliveries</span>
        </router-link>

        <router-link title="Billings" :class="$route.name == 'billings' ? 'bg-[#00000033]' : ''"
          :to="{ name: 'billings' }" class="header-icon-link px-5">
          <BanknotesIcon class="text-white h-5 w-5" />
          <span class="pl-2">Billings</span>
        </router-link>

        <router-link title="Reports" :class="$route.name == 'reports' ? 'bg-[#00000033]' : ''" :to="{ name: 'reports' }"
          class="header-icon-link px-5">
          <QueueListIcon class="text-white h-5 w-5" />
          <span class="pl-2">Reports</span>
        </router-link>

      </div>
      <div class="lg:flex hidden">
        <a v-if="remoteAddress" class="header-icon-link w-[48px]" @click="showQrCodePopup" title="Scan to Login">
          <QrCodeIcon class="text-white h-5 w-5" />
        </a>
        <router-link class="header-icon-link w-[48px]" :to="{ name: 'settings' }" title="Settings">
          <cog8-tooth-icon class="text-white h-5 w-5" />
        </router-link>
        <a title="Logout" class="header-icon-link px-3"
          @click="() => { auth.setJwt(''); user.reset(); $router.push({ name: 'login' }); }">
          <lock-closed-icon class="text-white h-5 w-5 mr-2" />
          <span class="hidden lg:block">Logout</span>
        </a>
      </div>
      <!-- mobile menu -->
      <div class="lg:hidden block">
        <a title="Menu" class="header-icon-link px-3" @click="toggleMenu">
          <Bars3Icon class="text-white h-5 w-5" />
        </a>
      </div>
      <div v-if="menuActive" class="h-screen fixed inset-0 bg-[#00000080] w-full left-0 top-0" @click="toggleMenu">
        <div class="h-screen right-0 fixed flex flex-col w-[320px] top-0 bottom-0 m-0 bg-white">
          <div class="h-[50px] flex items-center justify-between border-t-[1px] border-t-stone-200">
            <a v-if="remoteAddress" class="h-[50px] px-5 flex flex-1 items-center hover:bg-stone-50 cursor-pointer"
              @click="showQrCodePopup" title="Scan to Login">
              <QrCodeIcon class="h-5 w-5 mr-2" />
              Scan to login
            </a>
            <div></div>
            <div>
              <a class="h-[50px] w-[50px] flex border-l-[1px] flex-auto items-center justify-center hover:bg-stone-50 cursor-pointer"
                title="Close menu">
                <XMarkIcon class="h-5 w-5" />
              </a>
            </div>
          </div>
          <div class="border-t-[1px] page-scroll-area border-t-stone-200 text-[#0078d4]">
            <!-- staff -->
            <router-link v-if="user.role == '1'" title="Staff"
              :class="$route.name == 'staff' ? 'bg-[#0078d4] text-white' : ''" :to="{ name: 'staff' }"
              class="h-[50px] flex flex-auto items-center border-b-[1px] border-b-stone-200 cursor-pointer px-5 hover:bg-[#0078d4fe] hover:text-white">
              <component :is="UserIcon" class="h-5 w-5" />
              <span class="pl-2 capitalize">Staff</span>
            </router-link>
            <!-- others -->
            <router-link v-for="menu, i in menuItems" :key="i" :title="`${menu.name.toLocaleUpperCase()}`"
              :class="$route.name == menu.name ? 'bg-[#0078d4] text-white' : ''" :to="{ name: menu.name }"
              class="h-[50px] flex flex-auto items-center border-b-[1px] border-b-stone-200 cursor-pointer px-5 hover:bg-[#0078d4fe] hover:text-white">
              <component :is="menu.icon" class="h-5 w-5" />
              <span class="pl-2 capitalize">{{ menu.name }}</span>
            </router-link>
          </div>
          <div class="h-[50px] border-t-[1px] flex items-center justify-between">
            <a v-if="remoteAddress" class="h-[50px] px-5 flex flex-1 items-center bg-red-600 text-white hover:bg-red-600 cursor-pointer"
            @click="() => { auth.setJwt(''); user.reset(); $router.push({ name: 'login' }); }" title="Logout">
              <lock-closed-icon class="h-5 w-5 mr-2" />
              Logout
            </a>
          </div>
        </div>
      </div>
    </header>
    <main class="flex-1 flex-scroll">
      <router-view></router-view>
      <QrCodePopup :popup-id="qrCodePopupId" :remote-address="remoteAddress" />
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
}
</style>