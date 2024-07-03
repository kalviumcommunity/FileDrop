<template>
    <div class="h-screen">
        <div class="w-full h-5/6 bg-black flex justify-center items-center text-white">
            <div>
                <div>
                    <img src="../assets/img1.png" alt="Image 1"
                        class="absolute top-0 left-0 rounded-xl h-2/3 filter grayscale-100 mix-blend-luminosity">
                    <img src="../assets/img2.png" alt="Image 2"
                        class="absolute bottom-0 right-0 rounded-xl h-2/3 filter grayscale-100 mix-blend-luminosity">
                </div>
                <p class="text-lightPink font-Dela font-bold text-xxl  relative">FILEDROP</p>
                <div class="flex justify-center items-center gap-36">
                    <router-link to="/upload"><button type="button" class="font-Lime text-3xl rounded  relative"
                            :class="{ 'blink': blinkAnimation }">Upload</button></router-link>
                    <div class="w-2 h-2 bg-white rounded-full"></div>
                    <router-link to="/download"><button type="button" class="font-Lime text-3xl rounded  relative"
                            :class="{ 'blink': blinkAnimation }">Download</button></router-link>
                </div>
            </div>
        </div>
        <div class="flex">
            <div class="w-5/12 flex flex-col justify-center py-24 ml-40">
                <p class="font-Dela text-3xl">
                    Welcome to Easy File Sharing!
                </p>
                <br>
                <p class="font-Libre text-2xl w-10/12">
                    Are you tired of complicated and messy file sharing solutions? Look no further! Our incredible platform
                    allows users to effortlessly upload and download files in a snap.
                </p>
            </div>
            <div
                class="box-border w-[600px] h-[300px] overflow-x-hidden opacity-30 rounded-100 border border-black absolute rotate-[134deg] right-[10%]">
            </div>
        </div>
        <div
            class="box-border w-[480px] h-[250px] overflow-hidden opacity-30 rounded-100 border border-black absolute rotate-[134deg] left-[10%]">
        </div>
        <div class="bg-Moderateorange py-14 px-20 flex flex-col justify-center items-center gap-40">
            <div class="flex items-center justify-evenly">
                <img src="../assets/img3.jpg" class="w-80 h-80 rounded-2xl">
                <div class="w-4/12 flex flex-col justify-center py-20">
                    <p class="font-Dela text-3xl">
                        One-click uploads for lightning-fast file sharing
                    </p>
                    <br>
                    <p class="font-Libre text-base w-10/12">
                        Just click, drag, and go! Our file sharing platform makes it easier than ever to upload those
                        presentations, images, and videos that your team needs.
                    </p>
                </div>
            </div>
            <div class="flex flex-row-reverse items-center justify-evenly">
                <div class="w-80 h-80 flex flex-col">
                    <div class="absolute flex items-center justify-center h-1.5 w-56 self-end">
                        <div class="absolute w-px h-52 bg-black opacity-30"></div>
                        <div class="absolute w-px h-52 bg-black opacity-30 rotate-[30deg]"></div>
                        <div class="absolute w-px h-52 bg-black opacity-30 rotate-[60deg]"></div>
                        <div class="absolute w-px h-52 bg-black opacity-30 rotate-[90deg]"></div>
                        <div class="absolute w-px h-52 bg-black opacity-30 rotate-[120deg]"></div>
                        <div class="absolute w-px h-52 bg-black opacity-30 rotate-[150deg]"></div>
                    </div>
                    <img src="../assets/img3.jpg" class="w-full h-full rounded-xxl rounded-tl-xl">
                </div>
                <div class="w-4/12 flex flex-col justify-center py-20">
                    <p class="font-Dela text-3xl">
                        Secure downloads with a no-hassle experience
                    </p>
                    <br>
                    <p class="font-Libre text-base w-10/12">
                        Say goodbye to sketchy downloads. Easy File Sharing provides a secure and user-friendly environment
                        for downloading files rapidly.
                    </p>
                </div>
            </div>
            <div class="w-4/12 text-center">
                <div class="crossOval absolute left-[10%]">
                    <div
                        class="box-border w-[480px] h-[250px] overflow-hidden opacity-30 rounded-100 border border-black absolute">
                    </div>
                    <div
                        class="box-border w-[480px] h-[250px] overflow-hidden opacity-30 rounded-100 border border-black absolute rotate-[70deg]">
                    </div>
                </div>
                <p class="font-Dela text-3xl">
                    Get Started
                </p>
                <br>
                <p class="font-Libre text-base">
                    Are you ready to revolutionize your file sharing experience? Join the ranks of our satisfied users and
                    never look back. Click below to embark on the journey towards file sharing nirvana.
                </p>
                <br>
                <button class="font-Libre px-4 py-2 bg-lightPink rounded" @click=scrollToTop()>Try Now</button>
            </div>
        </div>
        <footer class="h-48 bg-Moderateorange text-center">
            <div class="h-full flex flex-col justify-center items-center gap-6">
                <div class="flex space-x-4">
                    <a href="https://www.instagram.com/" target="_blank">
                        <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="30" height="30" viewBox="0 0 30 30">
                            <path
                                d="M 9.9980469 3 C 6.1390469 3 3 6.1419531 3 10.001953 L 3 20.001953 C 3 23.860953 6.1419531 27 10.001953 27 L 20.001953 27 C 23.860953 27 27 23.858047 27 19.998047 L 27 9.9980469 C 27 6.1390469 23.858047 3 19.998047 3 L 9.9980469 3 z M 22 7 C 22.552 7 23 7.448 23 8 C 23 8.552 22.552 9 22 9 C 21.448 9 21 8.552 21 8 C 21 7.448 21.448 7 22 7 z M 15 9 C 18.309 9 21 11.691 21 15 C 21 18.309 18.309 21 15 21 C 11.691 21 9 18.309 9 15 C 9 11.691 11.691 9 15 9 z M 15 11 A 4 4 0 0 0 11 15 A 4 4 0 0 0 15 19 A 4 4 0 0 0 19 15 A 4 4 0 0 0 15 11 z">
                            </path>
                        </svg>
                    </a>
                    <a href="https://github.com/" target="_blank">
                        <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="30" height="30" viewBox="0 0 30 30">
                            <path
                                d="M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z">
                            </path>
                        </svg>
                    </a>
                    <a href="https://www.linkedin.com/in" target="_blank">
                        <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="30" height="30" viewBox="0 0 50 50">
                            <path
                                d="M41,4H9C6.24,4,4,6.24,4,9v32c0,2.76,2.24,5,5,5h32c2.76,0,5-2.24,5-5V9C46,6.24,43.76,4,41,4z M17,20v19h-6V20H17z M11,14.47c0-1.4,1.2-2.47,3-2.47s2.93,1.07,3,2.47c0,1.4-1.12,2.53-3,2.53C12.2,17,11,15.87,11,14.47z M39,39h-6c0,0,0-9.26,0-10 c0-2-1-4-3.5-4.04h-0.08C27,24.96,26,27.02,26,29c0,0.91,0,10,0,10h-6V20h6v2.56c0,0,1.93-2.56,5.81-2.56 c3.97,0,7.19,2.73,7.19,8.26V39z">
                            </path>
                        </svg>
                    </a>
                </div>
                <p>&copy; 2023 FileDrop. All rights reserved.</p>
            </div>
        </footer>
    </div>
</template>


<script setup>
import { ref, onMounted } from 'vue';

const blinkAnimation = ref(false);
const isButtonClicked = ref(false)

function scrollToTop() {
    window.scrollTo({
        top: 0,
        behavior: "smooth"
    });
    isButtonClicked.value = true
}

onMounted(() => {
    window.addEventListener('scroll', handleScroll);
});

function handleScroll() {
    const scrollTop = window.scrollY;

    if (scrollTop === 0 && isButtonClicked.value) {
        blinkAnimation.value = true;
        setTimeout(() => {
            blinkAnimation.value = false;
        }, 1000);
        isButtonClicked.value = false
    }
}

</script>

<style scoped>
@keyframes blink {

    0%,
    50%,
    100% {
        color: rgb(0, 0, 0);
    }

    25%,
    75% {
        color: white;
    }
}

.blink {
    animation: blink 1s linear;
}
</style>
