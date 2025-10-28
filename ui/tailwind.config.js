/** @type {import('tailwindcss').Config} */
export default {
      content: [
        './assets/**/*.html',
        './src/**/*.{js,jsx,ts,tsx}'
      ],
      theme: {
        extend: {},
      },
      plugins: [require('daisyui')],
    }