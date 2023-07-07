<template>
    <div class="flex flex-col h-screen font-Dela">
        <nav class="flex items-center justify-between bg-black py-4 px-6">
            <router-link to="/">
                <div class="text-white text-lg tracking-wider">FILEDROP</div>
            </router-link>
            <div>
                <router-link to="/upload"><button class="border border-white text-white py-2 px-4 rounded tracking-wide">
                        Upload
                    </button>
                </router-link>
            </div>
        </nav>
        <div class="flex-1 flex">
            <div class="bg-Moderateorange w-[25%] flex justify-center items-center">
                <p class="rotate-90 text-5xl">Download</p>
            </div>
            <div class="flex-1 bg-gray-100 flex justify-center items-center">
                <div v-show="!showFile" class="flex flex-col gap-5 justify-center items-center">
                    <input type="number" class="border border-black rounded py-2 px-4 no-spinner"
                        placeholder="Enter the code" ref="inputField" />
                    <button class="border border-black text-black py-2 px-4 rounded tracking-wide"
                        @click="showDownloadFile">
                        Search
                    </button>
                </div>
                <div v-show="showFile">
                    <div>
                        <p>Sample folder</p>
                        <p>4.0 GB</p>
                    </div>
                    <br />
                    <button class="border border-black text-black py-2 px-4 rounded tracking-wide">
                        Download
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script setup>
import { ref } from 'vue';

const showFile = ref(false);
const inputField = ref('');

const showDownloadFile = () => {
    if (inputField.value === '') {
        console.log('no')
        return;
    }
    showFile.value = true;
};
</script>

<style>
.no-spinner::-webkit-inner-spin-button,
.no-spinner::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
}
</style>