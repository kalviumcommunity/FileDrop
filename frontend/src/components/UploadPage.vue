<template>
    <div class="flex flex-col h-screen font-Dela">
        <nav class="flex items-center justify-between bg-black py-4 px-6">
            <router-link to="/">
                <div class="text-white text-lg tracking-wider">FILEDROP</div>
            </router-link>
            <div>
                <router-link to="/download"><button class="border border-white text-white py-2 px-4 rounded tracking-wide">
                        Download
                    </button>
                </router-link>
            </div>
        </nav>
        <div class="flex-1 flex">
            <div class="bg-Moderateorange w-[25%] flex justify-center items-center">
                <p class="rotate-90 text-5xl">Upload</p>
            </div>
            <div class="flex-1 bg-gray-100 flex justify-center items-center">
                <div v-if="!showUploader">
                    USE THIS CODE TO DOWNLOAD THE FILE
                    <br><br>
                    {{ uploadedData }}
                </div>
                <lr-file-uploader-regular v-else
                    css-src="https://esm.sh/@uploadcare/blocks@0.22.13/web/file-uploader-regular.min.css"
                    ctx-name="my-uploader" class="my-config">
                </lr-file-uploader-regular>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import * as LR from "@uploadcare/blocks";

const showUploader = ref(true);
const uploadedData = ref('');

onMounted(() => {
    const compilerOptions = {
        isCustomElement: (tag) => tag.startsWith("lr-"),
    };

    LR.registerBlocks(LR, compilerOptions);

    window.addEventListener('LR_DATA_OUTPUT', (e) => {
        if (e.detail.data && e.detail.data.length > 0) {
            showUploader.value = false;
            uploadedData.value = e.detail.data[0].uuid;
        }
    });
});
</script>

<style>
.my-config {
    --cfg-pubkey: "162c19c1d80ea7805643";
    --cfg-img-only: 1;
    --cfg-multiple: 0;
    --cfg-max-local-file-size-bytes: 10000000;
    --cfg-use-cloud-image-editor: 1;
    --cfg-source-list: "local, url, camera, dropbox";
    --darkmode: 0;
    --h-accent: 223;
    --s-accent: 100%;
    --l-accent: 61%;
}
</style>
