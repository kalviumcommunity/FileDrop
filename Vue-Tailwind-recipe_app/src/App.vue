<template>
  <div class="w-full h-full bg-slate-300 flex flex-col items-center gap-5 overflow-x-hidden">
    <div class="w-screen h-fit bg-indigo-900 text-sm">
      <div class="w-screen h-20 px-8 bg-indigo-950 flex flex-wrap items-center justify-between">
        <img src="./assets/logo.png" alt="Logo">
        <button @click="reloadPage()"
          class="w-24 h-12 text-white bg-black font-bold text-sm rounded cursor-pointer shadow">Home</button>
      </div>
      <br>
      <div class="text-white text-center">
        <h1 class="pb-8 font-semibold text-2xl">
          Welcome to FoodHub
        </h1>
        <p class="mx-24">
          Our web application is designed to help eliminate the daily struggle of deciding what to make for dinner.
          With FoodHub, you'll receive new and delicious meal ideas every day, taking the guesswork out of meal planning.
          Enjoy stress-free cooking and delicious meals with FoodHub.
        </p>
        <br>
      </div>
    </div>
    <h2 class="font-bold text-2xl">Random Meal</h2>
    <div class="w-full h-fit flex justify-center items-center" ref="randomMealCard">
      <div class="w-56 h-80 px-2 py-2 bg-white rounded-b-xl text-center flex flex-col" @click="showIngredients()">
        <img class="rounded-2xl" :src="mealData.strMealThumb" alt="Meal Image">
        <h3 class="font-semibold px-2">{{ mealData.strMeal }}</h3>
      </div>
    </div>
    <br>
    <div v-if="showPopup" class="fixed top-0 left-0 w-full h-full bg-black flex justify-center items-center">
      <div class="bg-white p-5 rounded-lg">
        <h2 class="font-semibold text-2xl">Ingredients</h2>
        <ul>
          <li v-for="ingredient in mealData.ingredients" :key="ingredient">{{ ingredient }}</li>
        </ul>
        <button class="mt-3 px-4 py-3 bg-neutral-800 border-none text-white rounded cursor-pointer" @click="closePopup()">Close</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';

const mealData = ref({
  strMeal: '',
  strMealThumb: '',
  strInstructions: '',
  ingredients: []
});

const showPopup = ref(false);

let reloadPage = () => {
  window.location.reload();
};

async function randomMeal() {
  let arr;
  await fetch("https://www.themealdb.com/api/json/v1/1/random.php")
    .then((data) => data.json())
    .then((e) => {
      arr = e;
    })
    .catch((error) => console.log(error));

  mealData.value.strMeal = arr.meals[0].strMeal;
  mealData.value.strMealThumb = arr.meals[0].strMealThumb;
  mealData.value.strInstructions = arr.meals[0].strInstructions;
  mealData.value.ingredients = extractIngredients(arr.meals[0]);
}

function extractIngredients(meal) {
  const ingredients = [];
  for (let i = 1; i <= 20; i++) {
    const ingredient = meal[`strIngredient${i}`];
    const measurement = meal[`strMeasure${i}`];
    if (ingredient && ingredient.trim() !== '') {
      const ingredientString = `${measurement ? measurement.trim() + ' ' : ''}${ingredient.trim()}`;
      ingredients.push(ingredientString);
    }
  }
  return ingredients;
}

function showIngredients() {
  showPopup.value = true;
}

function closePopup() {
  showPopup.value = false;
}

onMounted(() => {
  randomMeal();
});
</script>

