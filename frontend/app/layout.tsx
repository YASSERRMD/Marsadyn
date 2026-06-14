import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Marsadyn - Observability Platform',
  description: 'Production-grade distributed observability platform',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
