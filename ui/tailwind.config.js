/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'ui-sans-serif', 'system-ui'],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        light: {
          "primary": "#2563eb",
          "primary-focus": "#1d4ed8",
          "primary-content": "#ffffff",
          "secondary": "#64748b",
          "secondary-focus": "#475569",
          "secondary-content": "#ffffff",
          "accent": "#3b82f6",
          "accent-focus": "#2563eb",
          "accent-content": "#ffffff",
          "neutral": "#3d4451",
          "base-100": "#ffffff",
          "base-200": "#f8fafc",
          "base-300": "#f1f5f9",
          "base-content": "#1f2937",
          "--rounded-btn": "0.5rem",
          "--rounded-box": "0.5rem",
        },
      },
    ],
  },
} 