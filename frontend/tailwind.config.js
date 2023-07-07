/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontSize: {
        xxl: "9rem",
      },
      colors: {
        lightPink: "#FFC9B9",
        Moderateorange: "#D68C45",
      },
      fontFamily: {
        Dela: ["Dela Gothic One"],
        Lime: ["Limelight"],
        Libre: ["LibreBaskerville"],
      },
      borderRadius: {
        xxl: "125px",
        100: "100%",
      },
      rotate: {
        134: "134 deg",
      },
    },
  },
  plugins: [],
};
