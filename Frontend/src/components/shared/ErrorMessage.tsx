import React from 'react';
import { AlertCircle, XCircle } from 'lucide-react';

interface ErrorMessageProps {
  message: string;
  variant?: 'inline' | 'toast';
  onClose?: () => void;
}

export function ErrorMessage({ 
  message, 
  variant = 'inline',
  onClose 
}: ErrorMessageProps) {
  
  if (!message) return null; // Jangan render jika tidak ada message
  
  if (variant === 'toast') {
    return (
      <div className="fixed top-4 right-4 z-50 animate-in slide-in-from-top-5">
        <div className="bg-red-50 border-l-4 border-red-500 rounded-lg shadow-lg p-4 max-w-md">
          <div className="flex items-start">
            <AlertCircle className="h-5 w-5 text-red-500 mt-0.5" />
            <div className="ml-3 flex-1">
              <p className="text-sm font-medium text-red-800">{message}</p>
            </div>
            {onClose && (
              <button
                onClick={onClose}
                className="ml-3 text-red-500 hover:text-red-700 transition-colors"
              >
                <XCircle className="h-5 w-5" />
              </button>
            )}
          </div>
        </div>
      </div>
    );
  }
  
  // Inline variant (default)
  return (
    <div className="bg-red-50 border border-red-200 rounded-lg p-3 flex items-start gap-2">
      <AlertCircle className="h-5 w-5 text-red-500 mt-0.5 flex-shrink-0" />
      <p className="text-sm text-red-700 flex-1">{message}</p>
    </div>
  );
}

// Success variant (bonus)
export function SuccessMessage({ message }: { message: string }) {
  if (!message) return null;
  
  return (
    <div className="bg-green-50 border border-green-200 rounded-lg p-3 flex items-start gap-2">
      <AlertCircle className="h-5 w-5 text-green-500 mt-0.5 flex-shrink-0" />
      <p className="text-sm text-green-700 flex-1">{message}</p>
    </div>
  );
}