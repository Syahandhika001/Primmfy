'use client';

import { useState } from 'react';
import { LoadingSpinner, LoadingPage } from '@/components/shared/LoadingSpinner';
import { ErrorMessage, SuccessMessage } from '@/components/shared/ErrorMessage';
import { Button } from '@/components/ui/button';
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

export default function Home() {
  const [showLoading, setShowLoading] = useState(false);
  const [showError, setShowError] = useState(false);
  const [showSuccess, setShowSuccess] = useState(false);

  return (
    <div className="min-h-screen bg-gradient-primary p-8">
      <div className="max-w-4xl mx-auto space-y-8">
        <h1 className="text-4xl font-bold text-gray-800 text-center">
          PRIMMFY Component Testing
        </h1>
        
        {/* Loading Spinner Test */}
        <Card>
          <CardHeader>
            <CardTitle>Loading Spinner</CardTitle>
            <CardDescription>Test different spinner sizes</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="flex gap-4 items-center">
              <div className="text-center">
                <LoadingSpinner size="sm" />
                <p className="text-xs mt-2 text-gray-500">Small</p>
              </div>
              <div className="text-center">
                <LoadingSpinner size="md" />
                <p className="text-xs mt-2 text-gray-500">Medium</p>
              </div>
              <div className="text-center">
                <LoadingSpinner size="lg" />
                <p className="text-xs mt-2 text-gray-500">Large</p>
              </div>
            </div>
          </CardContent>
          <CardFooter>
            <Button onClick={() => setShowLoading(!showLoading)}>
              Toggle Full Page Loading
            </Button>
          </CardFooter>
        </Card>
        {showLoading && <LoadingPage />}

        {/* Error Message Test */}
        <Card>
          <CardHeader>
            <CardTitle>Error & Success Messages</CardTitle>
            <CardDescription>Test notification components</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <p className="text-sm font-medium mb-2">Inline Error:</p>
              <ErrorMessage message="This is an inline error message" />
            </div>
            <div>
              <p className="text-sm font-medium mb-2">Success Message:</p>
              <SuccessMessage message="Operation completed successfully!" />
            </div>
          </CardContent>
          <CardFooter className="gap-2">
            <Button onClick={() => setShowError(!showError)}>
              Toggle Toast Error
            </Button>
            <Button 
              variant="secondary" 
              onClick={() => setShowSuccess(!showSuccess)}
            >
              Toggle Success Toast
            </Button>
          </CardFooter>
        </Card>
        {showError && (
          <ErrorMessage 
            message="This is a toast error notification!"
            variant="toast"
            onClose={() => setShowError(false)}
          />
        )}

        {/* Form Components Test */}
        <Card>
          <CardHeader>
            <CardTitle>Form Components</CardTitle>
            <CardDescription>Test input, label, and button components</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="email">Email</Label>
              <Input 
                id="email" 
                type="email" 
                placeholder="Enter your email" 
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="password">Password</Label>
              <Input 
                id="password" 
                type="password" 
                placeholder="Enter your password" 
              />
            </div>
          </CardContent>
          <CardFooter className="gap-2">
            <Button>Primary Button</Button>
            <Button variant="secondary">Secondary</Button>
            <Button variant="outline">Outline</Button>
            <Button variant="destructive">Destructive</Button>
          </CardFooter>
        </Card>
      </div>
    </div>
  );
}