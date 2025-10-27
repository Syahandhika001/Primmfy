import type { Config } from 'tailwindcss';

const config: Config = {
  darkMode: ['class'],
  content: [
    './pages/**/*.{ts,tsx}',
    './components/**/*.{ts,tsx}',
    './app/**/*.{ts,tsx}',
    './src/**/*.{ts,tsx}',
  ],
  theme: {
    container: {
      center: true,
      padding: '2rem',
      screens: {
        '2xl': '1400px',
      },
    },
    extend: {
      colors: {
        // Brand Colors (from palette)
        brand: {
          blue: '#C2E2FA',
          'blue-dark': '#9BCEF5',
          'blue-light': '#E5F3FD',
          lavender: '#B7A3E3',
          'lavender-dark': '#9B81D9',
          'lavender-light': '#E0D7F4',
          pink: '#FF8F8F',
          'pink-dark': '#FF6B6B',
          'pink-light': '#FFB8B8',
          cream: '#FFF1CB',
          'cream-dark': '#FFE8A3',
          'cream-light': '#FFF8E7',
        },

        // PRIMM Stage Colors
        primm: {
          predict: '#C2E2FA',
          run: '#B7A3E3',
          investigate: '#FF8F8F',
          modify: '#FFF1CB',
          make: '#10B981',
        },

        // Semantic Colors
        border: 'hsl(var(--border))',
        input: 'hsl(var(--input))',
        ring: 'hsl(var(--ring))',
        background: 'hsl(var(--background))',
        foreground: 'hsl(var(--foreground))',
        
        primary: {
          DEFAULT: '#C2E2FA',
          foreground: '#374151',
        },
        secondary: {
          DEFAULT: '#B7A3E3',
          foreground: '#FFFFFF',
        },
        destructive: {
          DEFAULT: '#FF8F8F',
          foreground: '#FFFFFF',
        },
        muted: {
          DEFAULT: '#F3F4F6',
          foreground: '#6B7280',
        },
        accent: {
          DEFAULT: '#FFF1CB',
          foreground: '#374151',
        },
        popover: {
          DEFAULT: 'hsl(var(--popover))',
          foreground: 'hsl(var(--popover-foreground))',
        },
        card: {
          DEFAULT: 'hsl(var(--card))',
          foreground: 'hsl(var(--card-foreground))',
        },
      },
      borderRadius: {
        lg: 'var(--radius)',
        md: 'calc(var(--radius) - 2px)',
        sm: 'calc(var(--radius) - 4px)',
      },
      backgroundImage: {
        'gradient-primary': 'linear-gradient(135deg, #C2E2FA 0%, #E5F3FD 100%)',
        'gradient-secondary': 'linear-gradient(135deg, #B7A3E3 0%, #E0D7F4 100%)',
        'gradient-accent': 'linear-gradient(135deg, #FF8F8F 0%, #FFB8B8 100%)',
        'gradient-progress': 'linear-gradient(90deg, #C2E2FA 0%, #B7A3E3 100%)',
      },
      keyframes: {
        'accordion-down': {
          from: { height: '0' },
          to: { height: 'var(--radix-accordion-content-height)' },
        },
        'accordion-up': {
          from: { height: 'var(--radix-accordion-content-height)' },
          to: { height: '0' },
        },
      },
      animation: {
        'accordion-down': 'accordion-down 0.2s ease-out',
        'accordion-up': 'accordion-up 0.2s ease-out',
      },
    },
  },
  plugins: [require('tailwindcss-animate')],
};

export default config;