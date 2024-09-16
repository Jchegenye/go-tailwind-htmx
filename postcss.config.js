// Process Tailwind with PostCSS
module.exports = {
  plugins: [
    require('tailwindcss'), // Tailwind CSS plugin
    require('autoprefixer'), // Adds vendor prefixes to CSS rules
  ],
}
