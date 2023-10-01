<script setup lang="ts">
import dayjs from 'dayjs';
import PopUp from '@/components/PopUp.vue';
import { Patient, PatientHistory, PatientHistoryTypes } from '@/interfaces';
import { hospital, popupStore } from '@/stores';
import { ref } from 'vue';
import CheckboxField from '@/components/form/CheckboxField.vue';
import DocumentViewPopup from '@/components/popups/DocumentViewPopup.vue';
defineProps<{ popupId: string, history: PatientHistory, patient: Patient }>()
const historyTypes = ref(['General', 'Appointment', 'Admission', 'Discharge', 'Diagnosis', 'Examination', 'TestResult', 'Treatment'] as PatientHistoryTypes[]);

const activeTab = ref('General' as PatientHistoryTypes);

const currentDocumentUrl = ref('');
const viewDocument = (documentUrl: string) => {
    currentDocumentUrl.value = documentUrl;
    popupStore.child_id = 'document';
    popupStore.child_show = true;
}

</script>
<template>
    <pop-up :width="920" :title="`${patient.lastname} ${patient.firstname} - History Details`" :id="popupId"
        pop-class="min-h-[95vh]">
        <div class="mb-5">
            <div>
                <CheckboxField :options="historyTypes" label="History Type" v-model="activeTab" />
            </div>
            <div v-for="his, index in history.details" :key="index"
                :class="activeTab.toLowerCase() == index ? '' : 'hidden'">
                <div class="grid grid-cols-2 gap-x-2 mt-4">
                    <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                        <div class="title">Date:</div>
                        <div class="">{{ dayjs(history.date_time).format('DD-MM-YYYY hh:mm A') }}</div>
                    </div>
                    <div></div>
                    <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded" v-if="index == 'appointment'">
                        <div class="title">Appointment Type:</div>
                        <div class="">{{ his.appointment_type || "N/A" }}</div>
                    </div>
                    <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded" v-if="index == 'admission'">
                        <div class="title">Ward:</div>
                        <div class="">{{ his.ward_number || "N/A" }}</div>
                    </div>
                    <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded" v-if="index == 'admission'">
                        <div class="title">Room:</div>
                        <div class="">{{ his.room_number || "N/A" }}</div>
                    </div>
                </div>

                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Note:</div>
                    <div class="">{{ his.note || "N/A" }}</div>
                </div>
                <div class="mb-2 border-[1px] bg-blue-100 p-2 rounded pointer-cusrsor" v-if="his.document.length">
                    <div class="title">Document:</div>
                    <div @click="viewDocument(`${hospital.asset_base_url}/files/images_${patient.id}/${his.document}`)" class="cursor-pointer">
                        <div>
                            <img v-if="(['jpg', 'jpeg', 'png'].includes(his.document.substring(his.document.length - 3)))"
                                class="max-h-[250px] border-[1px] rounded-lg inline-block p-1 border-stone-500"
                                :src="`${hospital.asset_base_url}/files/images_${patient.id}/${his.document}`"
                                alt="Click to view" />
                            <div v-else
                                class="max-h-[250px] border-[1px] rounded-lg inline-block px-4 py-1 border-stone-500">Click
                                to view</div>
                        </div>
                    </div>
                </div>
            </div>

        </div>
        <DocumentViewPopup :document-url="currentDocumentUrl" popup-id="document" />
    </pop-up>
</template>

<style scoped>
.title {
    color: #0078d4;
    border-radius: 5px;
    font-weight: bold;
    margin-right: 5px;
}
</style>