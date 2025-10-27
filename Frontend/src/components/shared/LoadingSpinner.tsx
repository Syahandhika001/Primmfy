import React from 'react';

interface LoadingSpinnerProps {
  size?: 'sm' | 'md' | 'lg';
  className?: string;
}

export function LoadingSpinner({ size = 'md', className = '' }: LoadingSpinnerProps) {
  // Tentukan ukuran spinner berdasarkan prop 'size'
  const sizeClasses = {
    sm: 'h-4 w-4 border-2',  // Small: 16px
    md: 'h-8 w-8 border-3',  // Medium: 32px
    lg: 'h-12 w-12 border-4', // Large: 48px
  };

  return (
    <div className="flex items-center justify-center">
      <div
        className={`
          ${sizeClasses[size]}
          animate-spin 
          rounded-full 
          border-brand-blue 
          border-t-transparent
          ${className}
        `}
        role="status"
        aria-label="Loading"
      >
        <span className="sr-only">Loading...</span>
      </div>
    </div>
  );
}

// Variant: Full page loading
export function LoadingPage() {
  return (
    <div className="flex h-screen items-center justify-center bg-gradient-primary">
      <div className="text-center">
        <LoadingSpinner size="lg" />
        <p className="mt-4 text-gray-600 font-medium">Loading...</p>
      </div>
    </div>
  );
}