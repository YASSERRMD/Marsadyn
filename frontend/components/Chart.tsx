'use client'

interface ChartProps {
  title: string
  data: number[]
  labels: string[]
  color?: string
}

export function Chart({ title, data, labels, color = '#3B82F6' }: ChartProps) {
  const max = Math.max(...data)
  const min = Math.min(...data)
  const range = max - min || 1

  return (
    <div className="bg-gray-800 p-4 rounded-lg">
      <h3 className="text-lg font-semibold mb-4">{title}</h3>
      <div className="h-40 flex items-end">
        {data.map((value, index) => (
          <div
            key={index}
            className="flex-1 mx-0.5 rounded-t"
            style={{
              height: `${((value - min) / range) * 100}%`,
              backgroundColor: color,
              minHeight: '4px',
            }}
            title={`${labels[index]}: ${value}`}
          />
        ))}
      </div>
      <div className="flex justify-between mt-2 text-xs text-gray-400">
        <span>{labels[0]}</span>
        <span>{labels[labels.length - 1]}</span>
      </div>
    </div>
  )
}
