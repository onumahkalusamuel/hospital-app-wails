<script lang="ts" setup>
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import TextField from '@/components/form/TextField.vue';
import apiRequest from '@/services/http/api-requests';
import { auth } from '@/stores/auth';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import VueQrcode from '@chenfengyuan/vue-qrcode';

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
    console.log(remote);
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
</script>;

<template>
  <div class="h-screen flex flex-row">
    <div
      class="bg-[url('/hospital-image.png')] bg-cover bg-no-repeat md:flex-1 h-screen md:flex hidden justify-center items-center bg-[#0078d4]">
      <div class="text-xl text-center font-bold text-white bg-[#0078d4] border-[5px] border-white p-[15px]" v-if="remoteAddress">
        <vue-qrcode :value="remoteAddress" :options="{ width: 200 }"></vue-qrcode>
        <div class="pt-[10px]">SCAN WITH PHONE</div>
      </div>
    </div>
    <div class="p-5 flex flex-col justify-between flex-1 border-t-stone-400 border-t-[1px]">
      <div class="m-auto w-[24rem] flex flex-col h-screen justify-center">
        <div class="">
          <img :src="hospital_logo" class="max-w-[60px] my-[20px]" alt="logo">
          <div class="text-2xl mb-[40px]">Sign in to your account</div>
        </div>
        <div class="form-proper">
          <form v-on:submit.prevent="login" ref="loginForm" autocomplete="off" class="py-md space-y-3">
            <div><TextField name="username" label="Username" required /></div>
            <div><TextField name="password" label="Password" type="password" required /></div>
            <PrimaryButton type="submit">Sign in</PrimaryButton>
          </form>
          <div class="text-center mt-5"><small>HCMS v1.2.0</small></div>
        </div>
      </div>
    </div>
  </div>
</template>