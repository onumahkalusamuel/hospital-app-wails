<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { toasts, popupStore } from '@/stores';
import TextField from '@/components/form/TextField.vue';
import TextArea from '@/components/form/TextArea.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';

defineProps<{ popupId: string }>()
const admitPatientRef = ref(null);
const route = useRoute();
const emit = defineEmits(["update:closed"]);


const admitPatient = async () => {
  const formData = Object.fromEntries(new FormData(admitPatientRef.value as never as HTMLFormElement).entries())
  const admit = await apiRequest.post(`patients/${route.params.id}/admit-patient`, formData);
  if(admit.message) {
    toasts.addToast({message: "patient admitted successfully.", type: 'success'});
    emit('update:closed');
    popupStore.show = false;
  }
}

</script>
<template>
  <pop-up title="Admit Patient" :id="popupId">
    <form method="POST" v-on:submit.prevent="admitPatient" ref="admitPatientRef">
      <div class="flex gap-x-3">
        <div class="w-full"><TextField label="Ward Name" placeholder="Maternity Ward" name="ward_number" required></TextField></div>
        <div class="w-full"><TextField label="Room Number" placeholder="R-024" name="room_number" required></TextField></div>
      </div>
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