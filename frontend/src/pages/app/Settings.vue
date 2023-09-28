<script lang="ts" setup>
import { hospital, toasts, user } from '@/stores';
import apiRequest from '@/services/http/api-requests';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import { ref } from 'vue';

const hospitalHandle = ref(null);
const hospitalFormKey = ref(50);
const profileHandle = ref(null);
const passwordHandleKey = ref(1);

const updateProfile = async () => {
  const formData = new FormData(profileHandle.value as never as HTMLFormElement)
  const update = await apiRequest.put('profile', Object.fromEntries(formData.entries()));
  toasts.addToast({ message: update.message, type: 'success' });
  
  const req = await apiRequest.get("profile");
  if (req) user.setAll(req);
  passwordHandleKey.value++
}

const updateHospital = async () => {  
  const formData = new FormData(hospitalHandle.value as never as HTMLFormElement)
  const update = await apiRequest.postMulti('hospital-update', formData);
  toasts.addToast({ message: update.message, type: 'success' });

  const req = await apiRequest.get("hospital-details");
  if (req) hospital.setAll(req);
  hospitalFormKey.value++
}

</script>;

<template>
  <div class="p-[15px]">
    <div class="text-2xl pt-2 pb-4">Settings</div>
    <div></div>
    <div class="flex flex-col sm:flex-row gap-4">
      <div class="sm:w-[50%] border-[1px] border-blue-600 px-3 p-2 rounded">
        <h1 class="mt-3 text-xl">Profile Settings</h1>
        <div class="my-4">
          <form @submit.prevent="updateProfile" class="space-y-5" ref="profileHandle">
            <div class="flex space-x-2 mb-2">
              <div class="w-full"><TextField name="lastname" label="Last Name" :value="user.lastname" placeholder="Last Name" required/></div>
              <div class="w-full"><TextField name="firstname" label="First Name" :value="user.firstname" placeholder="First Name" required/></div>
            </div>
            <div class="flex space-x-2 mb-2">
              <div class="w-full"><TextField name="phone" label="Phone" :value="user.phone" placeholder="Phone"/></div>
              <div class="w-full"><TextField name="email" label="Email" :value="user.email" placeholder="Email"/></div>
            </div>
            <div class="w-full"><TextField name="password" label="Password" placeholder="Password (leave empty to skip password update)" :key="passwordHandleKey"/></div>
            <div class="flex">
              <div><PrimaryButton type="submit">Submit</PrimaryButton></div>
            </div>
          </form>
      </div>
      </div>
      <div class="sm:w-[50%] border-[1px] border-blue-600 px-3 p-2 rounded" v-if="user.role == '1'">
        <h1 class="mt-3 text-xl">System Settings</h1>
        <div class="my-4">
          <form v-on:submit.prevent="updateHospital" class="space-y-5" ref="hospitalHandle" enctype="multipart/form-data" :key="hospitalFormKey">
            <div class="flex space-x-2 mb-2">
              <div class="w-full"><TextField name="hospital_name" label="Hospital Name" :value="hospital.hospital_name" required/></div>
              <div class="w-full"><TextField name="hospital_address" label="Physical Address" :value="hospital.hospital_address" required/></div>
            </div>
            <div class="flex space-x-2 mb-2">
              <div class="w-full"><TextField name="hospital_email" label="Email Address" :value="hospital.hospital_email"/></div>
              <div class="w-full"><TextField name="hospital_phone" label="Phone Number" :value="hospital.hospital_phone"/></div>
            </div>
            <div class="w-full"><TextField label="Hospital Logo" name="hospital_logo" type="file" accept=".png, .jpg, .jpeg"/></div>
            <div class="flex">
              <div><PrimaryButton type="submit">Submit</PrimaryButton></div>
            </div>
          </form>
      </div>
      </div>
    </div>
   </div>
</template>