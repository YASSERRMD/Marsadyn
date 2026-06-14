'use client'

import { useState } from 'react'

interface FilterBarProps {
  onFilter: (filters: Record<string, string>) => void
}

export function FilterBar({ onFilter }: FilterBarProps) {
  const [filters, setFilters] = useState<Record<string, string>>({
    service: '',
    environment: '',
    level: '',
    search: '',
  })

  const handleApply = () => {
    const activeFilters: Record<string, string> = {}
    Object.entries(filters).forEach(([key, value]) => {
      if (value) {
        activeFilters[key] = value
      }
    })
    onFilter(activeFilters)
  }

  const handleReset = () => {
    setFilters({
      service: '',
      environment: '',
      level: '',
      search: '',
    })
    onFilter({})
  }

  return (
    <div className="flex items-center gap-4 flex-wrap">
      <input
        type="text"
        placeholder="Search..."
        value={filters.search}
        onChange={(e) => setFilters({ ...filters, search: e.target.value })}
        className="bg-gray-700 border border-gray-600 rounded px-3 py-1 text-sm"
      />
      <input
        type="text"
        placeholder="Service"
        value={filters.service}
        onChange={(e) => setFilters({ ...filters, service: e.target.value })}
        className="bg-gray-700 border border-gray-600 rounded px-3 py-1 text-sm"
      />
      <select
        value={filters.environment}
        onChange={(e) => setFilters({ ...filters, environment: e.target.value })}
        className="bg-gray-700 border border-gray-600 rounded px-3 py-1 text-sm"
      >
        <option value="">All Environments</option>
        <option value="production">Production</option>
        <option value="staging">Staging</option>
        <option value="development">Development</option>
      </select>
      <select
        value={filters.level}
        onChange={(e) => setFilters({ ...filters, level: e.target.value })}
        className="bg-gray-700 border border-gray-600 rounded px-3 py-1 text-sm"
      >
        <option value="">All Levels</option>
        <option value="debug">Debug</option>
        <option value="info">Info</option>
        <option value="warn">Warning</option>
        <option value="error">Error</option>
      </select>
      <button
        onClick={handleApply}
        className="px-4 py-1 bg-blue-600 hover:bg-blue-700 rounded text-sm"
      >
        Apply
      </button>
      <button
        onClick={handleReset}
        className="px-4 py-1 bg-gray-600 hover:bg-gray-500 rounded text-sm"
      >
        Reset
      </button>
    </div>
  )
}
