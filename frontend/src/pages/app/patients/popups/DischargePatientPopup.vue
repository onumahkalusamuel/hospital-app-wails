<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import {toasts, popupStore } from '@/stores';
import PopUp from '@/components/PopUp.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import TextArea from '@/components/form/TextArea.vue';


defineProps<{ popupId: string }>()
const dischargePatientRef = ref(null);
const route = useRoute();
const emit = defineEmits(["update:closed"]);

const dischargePatient = async () => {
  const formData = Object.fromEntries(new FormData(dischargePatientRef.value as never as HTMLFormElement).entries())
  const admit = await apiRequest.post(`patients/${route.params.id}/discharge-patient`, formData);
  if(admit.message) {
    toasts.addToast({message: "patient discharged successfully.", type: 'success'});
    emit('update:closed');
    popupStore.show = false;
  }
}

</script>
<template>
  <pop-up title="Discharge Patient" :id="popupId">
    <form method="POST" v-on:submit.prevent="dischargePatient" ref="dischargePatientRef">
      <div class="flex">
        <div class="w-full"><TextArea label="Note" placeholder="Note" name="note"></TextArea></div>
      </div>
      <div class="flex gap-3 mt-5">
        <div><PrimaryButton type="submit">Submit</PrimaryButton></div>
        <div><SecondaryButton type="button" v-on:click.prevent="()=>popupStore.show=false">Cancel</SecondaryButton></div>
      </div>
    </form>
  </pop-up>
</template>