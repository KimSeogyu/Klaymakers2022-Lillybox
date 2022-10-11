/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  variants: {},
  theme: {
    extend: {},
  },
  plugins: [require("tailwind-scrollbar-hide")],
};
