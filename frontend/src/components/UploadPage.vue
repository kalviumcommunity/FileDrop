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
                <div v-show="!showCode" class="w-full h-full flex flex-col justify-center items-center gap-5">
                    <label
                        class="flex justify-center w-1/2 h-1/2 px-4 transition bg-gray-100 border-2 border-gray-400 border-dashed rounded-md appearance-none cursor-pointer hover:border-gray-600 focus:outline-none">
                        <span class="flex items-center space-x-2">
                            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24"
                                stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round"
                                    d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                            </svg>
                            <span class="font-medium">
                                Drop files to Attach, or
                                <span class="text-blue-600 underline">browse</span>
                            </span>
                        </span>
                        <input type="file" name="file_upload" class="hidden">
                    </label>
                    <button class="border border-black text-black py-2 px-4 rounded" type="submit"
                        @click="showDownloadCode">
                        Upload
                    </button>
                </div>
                <div v-show="showCode" class="text-center">
                    <p>
                        USE THIS CODE TO DOWNLOAD THE FILE
                    </p>
                    <p>
                        25072005
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const showCode = ref(false);

function showDownloadCode() {
    showCode.value = true;
}
</script>