<script lang="ts" setup>
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import TextField from '@/components/form/TextField.vue';
import apiRequest from '@/services/http/api-requests';
import { auth } from '@/stores/auth';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import VueQrcode from '@chenfengyuan/vue-qrcode';
import { EyeIcon, EyeSlashIcon } from '@heroicons/vue/24/solid';
import QrCodePopup from '@/components/popups/QrCodePopup.vue';
import { popupStore } from '@/stores';

const hospital_name = ref('');
const hospital_address = ref('');
const hospital_logo = ref('');
const loginForm = ref(null);
const router = useRouter();
const remoteAddress = ref('');

onMounted(async () => {
  auth.setJwt("");
  try {
    const req = await apiRequest.get("hospital-details");
    if (req) {
      hospital_name.value = req.hospital_name;
      hospital_address.value = req.hospital_address;
      hospital_logo.value = `${req.asset_base_url}/files/logo/${req.hospital_logo}`;
    }

    const remote = await apiRequest.get("get-remote-address");
    if (remote.address) remoteAddress.value = remote.address;
  }
  catch (e) {
    console.log(e);
  }
})

const login = async () => {
  const formData = new FormData(loginForm.value as never as HTMLFormElement);
  const login = await apiRequest.post("login", Object.fromEntries(formData.entries()));
  if (login.message) {
    auth.setJwt(login.jwt);
    router.push({ name: "dashboard" });
  }
}

const form = ref({ password: '' });
const passwordVisible = ref(false);
const togglePassword = () => (passwordVisible.value = !passwordVisible.value);

</script>;

<template>
  <div class="h-screen flex flex-row transition-all">
    <div
      class="bg-[url('/hospital-image.png')] bg-cover bg-no-repeat md:flex-1 h-screen md:flex hidden justify-center items-center bg-[#0078d4]">
      <div class="text-xl text-center font-bold text-white bg-[#0078d4] border-[5px] border-white p-[15px]"
        v-if="remoteAddress">
        <vue-qrcode :value="remoteAddress" :options="{ width: 200 }"></vue-qrcode>
        <div class="pt-[10px]">SCAN WITH PHONE</div>
      </div>
    </div>
    <div
      class="md:bg-none md:bg-white bg-[url('/hospital-image.png')] bg-blend-overlay bg-cover bg-no-repeat p-5 flex flex-col justify-between flex-1 border-t-stone-400 border-t-[1px] bg-[#00000088]">
      <div class="m-auto md:w-[24rem] p-5 py-10 flex flex-col justify-center bg-[#ffffffcc] rounded-lg">
        <div class="">
          <img :src="hospital_logo" class="max-w-[60px] mb-[10px] md:mb-[20px]" alt="logo">
          <div class="text-2xl mb-[20px] md:mb-[40px]">Sign in to your account</div>
        </div>
        <div class="form-proper">
          <form v-on:submit.prevent="login" ref="loginForm" autocomplete="off" class="py-md space-y-3">
            <div>
              <TextField name="username" label="Username" required />
            </div>
            <div>
              <TextField name="password" label="Password" :type="passwordVisible ? 'text' : 'password'" required
                v-model="form.password">
                <template #append>
                  <div class="rounded-lg cursor-pointer" @click.prevent="togglePassword" title="Toggle Bible">
                    <!-- <a class="w-6 h-6 text-stone-600 text-sm underline">{{ passwordVisible ? 'hide' : 'show' }}</a> -->
                    <component :is="passwordVisible ? EyeSlashIcon : EyeIcon" class="w-6 h-5 text-stone-600" />
                  </div>
                </template>
              </TextField>
            </div>
            <PrimaryButton type="submit">Sign in</PrimaryButton>
          </form>
          <div class="flex flex-col text-center mt-6 gap-y-5">
            <small @click="() => { popupStore.id = 'qrcode'; popupStore.show = true }"
              class="text-blue-500 underline md:hidden cursor-pointer">Have another device? Scan QRCode</small>
            <small class="">HCMS v1.2.0</small>
          </div>
        </div>
      </div>
    </div>
    <QrCodePopup popup-id="qrcode" :remote-address="remoteAddress" />
  </div>
</template>