'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

const navigation = [
  { name: 'Dashboard', href: '/dashboard', icon: '📊' },
  { name: 'Metrics', href: '/metrics', icon: '📈' },
  { name: 'Logs', href: '/logs', icon: '📝' },
  { name: 'Traces', href: '/traces', icon: '🔍' },
  { name: 'Alerts', href: '/alerts', icon: '🚨' },
  { name: 'Incidents', href: '/incidents', icon: '⚠️' },
  { name: 'Applications', href: '/applications', icon: '📦' },
  { name: 'Services', href: '/services', icon: '🔧' },
  { name: 'Retention', href: '/retention', icon: '🗑️' },
  { name: 'Settings', href: '/settings', icon: '⚙️' },
]

export function Sidebar() {
  const pathname = usePathname()

  return (
    <div className="w-64 bg-gray-800 border-r border-gray-700">
      <div className="p-4">
        <h1 className="text-2xl font-bold">Marsadyn</h1>
        <p className="text-sm text-gray-400">Observability Platform</p>
      </div>
      <nav className="mt-4">
        {navigation.map((item) => (
          <Link
            key={item.name}
            href={item.href}
            className={`flex items-center px-4 py-3 text-sm ${
              pathname === item.href
                ? 'bg-gray-700 text-white'
                : 'text-gray-400 hover:bg-gray-700 hover:text-white'
            }`}
          >
            <span className="mr-3">{item.icon}</span>
            {item.name}
          </Link>
        ))}
      </nav>
    </div>
  )
}
