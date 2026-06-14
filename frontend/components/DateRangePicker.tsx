'use client'

import { useState } from 'react'

interface DateRangePickerProps {
  onChange: (start: Date, end: Date) => void
}

export function DateRangePicker({ onChange }: DateRangePickerProps) {
  const [start, setStart] = useState(
    new Date(Date.now() - 24 * 60 * 60 * 1000).toISOString().slice(0, 16)
  )
  const [end, setEnd] = useState(new Date().toISOString().slice(0, 16))

  const handleApply = () => {
    onChange(new Date(start), new Date(end))
  }

  const presets = [
    { label: 'Last 1h', hours: 1 },
    { label: 'Last 6h', hours: 6 },
    { label: 'Last 24h', hours: 24 },
    { label: 'Last 7d', hours: 24 * 7 },
  ]

  return (
    <div className="flex items-center gap-4">
      <div className="flex gap-2">
        {presets.map((preset) => (
          <button
            key={preset.label}
            onClick={() => {
              const now = new Date()
              const start = new Date(now.getTime() - preset.hours * 60 * 60 * 1000)
              setStart(start.toISOString().slice(0, 16))
              setEnd(now.toISOString().slice(0, 16))
            }}
            className="px-3 py-1 text-sm bg-gray-700 hover:bg-gray-600 rounded"
          >
            {preset.label}
          </button>
        ))}
      </div>
      <input
        type="datetime-local"
        value={start}
        onChange={(e) => setStart(e.target.value)}
        className="bg-gray-700 border border-gray-600 rounded px-3 py-1 text-sm"
      />
      <span className="text-gray-400">to</span>
      <input
        type="datetime-local"
        value={end}
        onChange={(e) => setEnd(e.target.value)}
        className="bg-gray-700 border border-gray-600 rounded px-3 py-1 text-sm"
      />
      <button
        onClick={handleApply}
        className="px-4 py-1 bg-blue-600 hover:bg-blue-700 rounded text-sm"
      >
        Apply
      </button>
    </div>
  )
}
