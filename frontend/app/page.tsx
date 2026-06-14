export default function Home() {
  return (
    <main className="min-h-screen bg-gray-900 text-white">
      <div className="container mx-auto px-4 py-16">
        <div className="text-center">
          <h1 className="text-6xl font-bold mb-4">Marsadyn</h1>
          <p className="text-xl text-gray-400 mb-8">
            Distributed Observability Platform
          </p>
          <div className="flex justify-center gap-4">
            <a
              href="/dashboard"
              className="bg-blue-600 hover:bg-blue-700 px-6 py-3 rounded-lg font-medium transition-colors"
            >
              Open Dashboard
            </a>
            <a
              href="/api/health"
              className="bg-gray-700 hover:bg-gray-600 px-6 py-3 rounded-lg font-medium transition-colors"
            >
              API Health
            </a>
          </div>
        </div>
        <div className="mt-16 grid grid-cols-1 md:grid-cols-3 gap-8">
          <div className="bg-gray-800 p-6 rounded-lg">
            <h2 className="text-xl font-semibold mb-2">Metrics</h2>
            <p className="text-gray-400">
              Collect and visualize time-series metrics from your services.
            </p>
          </div>
          <div className="bg-gray-800 p-6 rounded-lg">
            <h2 className="text-xl font-semibold mb-2">Logs</h2>
            <p className="text-gray-400">
              Centralized log aggregation with powerful search capabilities.
            </p>
          </div>
          <div className="bg-gray-800 p-6 rounded-lg">
            <h2 className="text-xl font-semibold mb-2">Traces</h2>
            <p className="text-gray-400">
              Distributed tracing for request flow visualization.
            </p>
          </div>
        </div>
      </div>
    </main>
  )
}
