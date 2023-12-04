/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.gohtml"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: [
      "light",
      "dark",
      "cupcake",
      "garden",
      "retro",
      "business",
      "aqua",
      "nord",
      {
        discolight: {
          primary: "#208de2",
          secondary: "#74B35E",
          accent: "#0085FB",
          neutral: "#0e100c",
          "base-100": "#E6EFFE",
          info: "#0086ff",
          success: "#74B35E",
          warning: "#FDBA5C",
          error: "#D35A68",
        },
      },
    ],
  },
};
