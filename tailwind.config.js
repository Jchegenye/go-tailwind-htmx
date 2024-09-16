const colors = require('tailwindcss/colors')
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './views/**/*.templ',
    './public/**/*.html',
    './views/**/*.templ.go',
    './src/**/*.js'
  ],
  theme: {
    fontFamily: {
      sans: ['"A2 Gothic"', 'sans-serif'],
    },
    container: {
      center: true,
      padding: {
        DEFAULT: '1rem',
        sm: '2rem',
        lg: '3rem',
        xl: '4rem',
        '2xl': '5rem',
      },
      screens: {
      //   // sm: '640px',
      //   // md: '768px',
      //   // lg: '1024px',
      //   // xl: '1280px',
      //   // '2xl': '1536px',
      },
    },
    extend: {
      colors: {
        background: 'hsl(var(--background))',
        foreground: 'hsl(var(--foreground))',
        primary: {
          DEFAULT: 'hsl(var(--primary))',
          foreground: 'hsl(var(--primary-foreground))',
        },
        border: 'hsl(var(--border))',
      },
    },
  },
  plugins: [],
}

