/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
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
    extend: {
      colors: {
        background: 'hsl(var(--background))',
        foreground: 'hsl(var(--foreground))',
        primary: {
          DEFAULT: 'hsl(var(--primary))',
          foreground: 'hsl(var(--primary-foreground))',
        },
        secondary: {
          DEFAULT: 'hsl(var(--secondary))',
          foreground: 'hsl(var(--secondary-foreground))',
        },
        border: 'hsl(var(--border))',
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
      },
      screens: {
        'xs': '376px', // Custom breakpoint for extra-small screens
        // sm: '640px',
        // md: '768px',
        // lg: '1024px',
        // xl: '1280px',
        // '2xl': '1536px',
      },
    },
  },
  plugins: [],
}

