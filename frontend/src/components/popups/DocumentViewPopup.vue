<script setup lang="ts">
import PopUp from '@/components/PopUp.vue';
import { onMounted, ref, watch } from 'vue';
import VuePdfEmbed from 'vue-pdf-embed'

const props = defineProps<{ popupId: string; documentUrl: string }>()
const doctype = ref('');

const getDocType = () => (doctype.value = (props.documentUrl.split('.').pop() as string).toLocaleLowerCase())

onMounted(async () => getDocType())
watch(() => props.documentUrl, () => getDocType())

</script>
<template>
    <pop-up title="Document View" :id="popupId" class="w-screen h-screen">
        <div class="text-center py-2">
            <a :href="documentUrl" download target="_blank"
                class="bg-blue-600 text-sm hover:bg-blue-700 active:bg-blue-800 py-1 px-3 text-white rounded-full">Download document</a>
        </div>
        <div v-if="doctype == 'pdf'">
            <vue-pdf-embed :source="documentUrl" />
        </div>
        <div class="flex">
            <div class="m-auto" v-if="['jpg', 'jpeg', 'png'].includes(doctype)">
                <img :src="documentUrl" alt="image" />
            </div>
        </div>
    </pop-up>
</template>