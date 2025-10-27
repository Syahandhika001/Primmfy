'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import Link from 'next/link';
import { useAuth } from '@/lib/contexts/AuthContext';
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import { ErrorMessage } from '@/components/shared/ErrorMessage';
import { LoadingSpinner } from '@/components/shared/LoadingSpinner';

// === 1. Validation Schema ===
const registerSchema = z.object({
  full_name: z
    .string()
    .min(1, 'Name is required')
    .min(3, 'Name must be at least 3 characters')
    .max(50, 'Name must not exceed 50 characters'),
  email: z
    .string()
    .min(1, 'Email is required')
    .email('Please enter a valid email'),
  password: z
    .string()
    .min(1, 'Password is required')
    .min(6, 'Password must be at least 6 characters'),
  role: z.enum(['teacher', 'student'], {
    required_error: 'Please select a role',
  }),
});

type RegisterFormData = z.infer<typeof registerSchema>;

// === 2. Register Page Component ===
export default function RegisterPage() {
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const { register: registerUser } = useAuth();

  // === 3. React Hook Form Setup ===
  const {
    register,
    handleSubmit,
    formState: { errors },
    watch,
  } = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      role: 'student', // Default role
    },
  });

  const selectedRole = watch('role');

  // === 4. Submit Handler ===
  const onSubmit = async (data: RegisterFormData) => {
    try {
      setIsLoading(true);
      setError('');
      await registerUser(data);
      // Redirect handled by AuthContext
    } catch (err: any) {
      setError(err?.message || 'Registration failed. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  // === 5. JSX Render ===
  return (
    <Card>
      <CardHeader className="space-y-1">
        <CardTitle className="text-2xl text-center">Create an account</CardTitle>
        <CardDescription className="text-center">
          Enter your information to get started with PRIMMFY
        </CardDescription>
      </CardHeader>

      <CardContent>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          {/* Global Error Message */}
          {error && <ErrorMessage message={error} />}

          {/* Full Name Field */}
          <div className="space-y-2">
            <Label htmlFor="full_name">Full Name</Label>
            <Input
              id="full_name"
              type="text"
              placeholder="John Doe"
              {...register('full_name')}
              disabled={isLoading}
            />
            {errors.full_name && (
              <p className="text-sm text-red-600">{errors.full_name.message}</p>
            )}
          </div>

          {/* Email Field */}
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              type="email"
              placeholder="john@example.com"
              {...register('email')}
              disabled={isLoading}
            />
            {errors.email && (
              <p className="text-sm text-red-600">{errors.email.message}</p>
            )}
          </div>

          {/* Password Field */}
          <div className="space-y-2">
            <Label htmlFor="password">Password</Label>
            <Input
              id="password"
              type="password"
              placeholder="••••••••"
              {...register('password')}
              disabled={isLoading}
            />
            {errors.password && (
              <p className="text-sm text-red-600">{errors.password.message}</p>
            )}
          </div>

          {/* Role Selection */}
          <div className="space-y-2">
            <Label>I am a</Label>
            <div className="grid grid-cols-2 gap-3">
              {/* Student Option */}
              <label
                className={`
                  relative flex flex-col items-center justify-center p-4 cursor-pointer
                  border-2 rounded-lg transition-all
                  ${
                    selectedRole === 'student'
                      ? 'border-[rgb(194,226,250)] bg-[rgb(229,243,253)]'
                      : 'border-gray-200 hover:border-gray-300'
                  }
                  ${isLoading ? 'cursor-not-allowed opacity-50' : ''}
                `}
              >
                <input
                  type="radio"
                  value="student"
                  {...register('role')}
                  disabled={isLoading}
                  className="sr-only"
                />
                <div className="text-3xl mb-2">🎓</div>
                <div className="font-semibold">Student</div>
                <div className="text-xs text-gray-500 text-center mt-1">
                  Learn programming
                </div>
              </label>

              {/* Teacher Option */}
              <label
                className={`
                  relative flex flex-col items-center justify-center p-4 cursor-pointer
                  border-2 rounded-lg transition-all
                  ${
                    selectedRole === 'teacher'
                      ? 'border-[rgb(183,163,227)] bg-[rgb(224,215,244)]'
                      : 'border-gray-200 hover:border-gray-300'
                  }
                  ${isLoading ? 'cursor-not-allowed opacity-50' : ''}
                `}
              >
                <input
                  type="radio"
                  value="teacher"
                  {...register('role')}
                  disabled={isLoading}
                  className="sr-only"
                />
                <div className="text-3xl mb-2">👨‍🏫</div>
                <div className="font-semibold">Teacher</div>
                <div className="text-xs text-gray-500 text-center mt-1">
                  Create lessons
                </div>
              </label>
            </div>
            {errors.role && (
              <p className="text-sm text-red-600">{errors.role.message}</p>
            )}
          </div>

          {/* Submit Button */}
          <Button type="submit" className="w-full" disabled={isLoading}>
            {isLoading ? (
              <>
                <LoadingSpinner size="sm" />
                <span className="ml-2">Creating account...</span>
              </>
            ) : (
              'Create Account'
            )}
          </Button>
        </form>
      </CardContent>

      <CardFooter className="flex flex-col space-y-4">
        <div className="text-sm text-center text-gray-600">
          Already have an account?{' '}
          <Link
            href="/login"
            className="text-[rgb(183,163,227)] font-semibold hover:underline"
          >
            Sign in
          </Link>
        </div>
      </CardFooter>
    </Card>
  );
}