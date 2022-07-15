module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        creo: "#1bb1dc",
      },
    },
  },
  plugins: [require("tailwind-scrollbar")],
};
