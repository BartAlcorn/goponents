/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["web/**/*.gohtml"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "cupcake", "garden", "retro", "business", "aqua"],
  }
}
