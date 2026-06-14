'use client'

import { useEffect, useState } from 'react'

interface HealthStatus {
  status: string
  service: string
  timestamp: string
  version: string
  goVersion: string
  uptime: string
}

export default function Dashboard() {
  const [health, setHealth] = useState<HealthStatus | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetch('http://localhost:8080/health')
      .then((res) => res.json())
      .then((data) => {
        setHealth(data)
        setLoading(false)
      })
      .catch((err) => {
        console.error('Failed to fetch health:', err)
        setLoading(false)
      })
  }, [])

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-6">Dashboard</h1>
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div className="bg-gray-800 p-6 rounded-lg">
          <h2 className="text-lg font-semibold mb-2">Metrics</h2>
          <p className="text-3xl font-bold text-blue-400">--</p>
          <p className="text-sm text-gray-400">Active series</p>
        </div>
        <div className="bg-gray-800 p-6 rounded-lg">
          <h2 className="text-lg font-semibold mb-2">Logs</h2>
          <p className="text-3xl font-bold text-green-400">--</p>
          <p className="text-sm text-gray-400">Last 24h</p>
        </div>
        <div className="bg-gray-800 p-6 rounded-lg">
          <h2 className="text-lg font-semibold mb-2">Traces</h2>
          <p className="text-3xl font-bold text-purple-400">--</p>
          <p className="text-sm text-gray-400">Last 24h</p>
        </div>
        <div className="bg-gray-800 p-6 rounded-lg">
          <h2 className="text-lg font-semibold mb-2">Alerts</h2>
          <p className="text-3xl font-bold text-red-400">--</p>
          <p className="text-sm text-gray-400">Active alerts</p>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="bg-gray-800 p-6 rounded-lg">
          <h2 className="text-lg font-semibold mb-4">System Status</h2>
          {loading ? (
            <p className="text-gray-400">Loading...</p>
          ) : health ? (
            <div className="space-y-2">
              <p><span className="text-gray-400">Status:</span> <span className="text-green-400">{health.status}</span></p>
              <p><span className="text-gray-400">Service:</span> {health.service}</p>
              <p><span className="text-gray-400">Version:</span> {health.version}</p>
              <p><span className="text-gray-400">Go Version:</span> {health.goVersion}</p>
              <p><span className="text-gray-400">Uptime:</span> {health.uptime}</p>
            </div>
          ) : (
            <p className="text-red-400">Failed to connect to API</p>
          )}
        </div>
        
        <div className="bg-gray-800 p-6 rounded-lg">
          <h2 className="text-lg font-semibold mb-4">Recent Activity</h2>
          <p className="text-gray-400">No recent activity</p>
        </div>
      </div>
    </div>
  )
}
