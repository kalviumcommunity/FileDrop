<template>
    <div class="flex flex-col h-screen font-Dela">
        <nav class="flex items-center justify-between bg-black py-4 px-6">
            <router-link to="/">
                <div class="text-white text-lg tracking-wider">FILEDROP</div>
            </router-link>
            <div>
                <router-link to="/upload">
                    <button class="border border-white text-white py-2 px-4 rounded tracking-wide">
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
                    <input type="text" class="border border-black rounded py-2 px-4 no-spinner" placeholder="Enter the code"
                        v-model="inputField" />
                    <button class="border border-black text-black py-2 px-4 rounded tracking-wide" @click="searchFile">
                        Search
                    </button>
                </div>
                <div v-show="showFile" class="flex flex-col items-center">
                    <div>
                        <img :src="getImageUrl()" alt="Sample folder" style="width: 400px; height: auto;" />
                    </div>
                    <br />
                    <button class="border border-black text-black py-2 px-4 rounded tracking-wide" @click="forceDownload()">
                        Download
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script setup>
import { ref } from 'vue';
import axios from 'axios';

const showFile = ref(false);
const inputField = ref('');

const searchFile = () => {
    if (inputField.value === '') {
        console.log('no')
        return;
    }
    console.log(inputField.value);
    showFile.value = true;
};

const getImageUrl = () => {
    if (showFile.value) {
        const code = inputField.value;
        return `https://ucarecdn.com/${code}/`;
    } else {
        return '';
    }
};


const forceDownload = async () => {
    const imageUrl = getImageUrl();
    try {
        const response = await axios.get(imageUrl, { responseType: 'blob' });
        const blob = new Blob([response.data]);
        const url = window.URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'file.png');
        link.style.display = 'none';
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
        window.URL.revokeObjectURL(url);
    } catch (error) {
        console.error('Error downloading the file:', error);
    }
};
</script>
  
<style>
.no-spinner::-webkit-inner-spin-button,
.no-spinner::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
}
</style>