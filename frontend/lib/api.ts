const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

interface RequestOptions {
  method?: string
  headers?: Record<string, string>
  body?: unknown
}

async function request<T>(endpoint: string, options: RequestOptions = {}): Promise<T> {
  const { method = 'GET', headers = {}, body } = options

  const config: RequestInit = {
    method,
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
  }

  if (body) {
    config.body = JSON.stringify(body)
  }

  const response = await fetch(`${API_BASE_URL}${endpoint}`, config)

  if (!response.ok) {
    throw new Error(`API error: ${response.status} ${response.statusText}`)
  }

  return response.json()
}

export const api = {
  health: {
    get: () => request<{ status: string; service: string; version: string }>('/health'),
  },
  
  metrics: {
    query: (params: Record<string, string>) => {
      const query = new URLSearchParams(params).toString()
      return request<unknown[]>(`/api/v1/query/metrics?${query}`)
    },
    summary: (params: Record<string, string>) => {
      const query = new URLSearchParams(params).toString()
      return request<unknown>(`/api/v1/summary/metrics?${query}`)
    },
  },
  
  logs: {
    query: (params: Record<string, string>) => {
      const query = new URLSearchParams(params).toString()
      return request<unknown[]>(`/api/v1/query/logs?${query}`)
    },
    summary: (params: Record<string, string>) => {
      const query = new URLSearchParams(params).toString()
      return request<unknown>(`/api/v1/summary/logs?${query}`)
    },
  },
  
  traces: {
    query: (params: Record<string, string>) => {
      const query = new URLSearchParams(params).toString()
      return request<unknown[]>(`/api/v1/query/traces?${query}`)
    },
    summary: (params: Record<string, string>) => {
      const query = new URLSearchParams(params).toString()
      return request<unknown>(`/api/v1/summary/traces?${query}`)
    },
  },
  
  alerts: {
    getRules: (tenantId: string) => 
      request<unknown[]>(`/api/v1/alerts/rules?tenantId=${tenantId}`),
    createRule: (rule: unknown) => 
      request<unknown>('/api/v1/alerts/rules', { method: 'POST', body: rule }),
    updateRule: (id: string, updates: unknown) => 
      request<unknown>(`/api/v1/alerts/rules/${id}`, { method: 'PATCH', body: updates }),
    getIncidents: (tenantId: string) => 
      request<unknown[]>(`/api/v1/alerts/incidents?tenantId=${tenantId}`),
    resolveIncident: (id: string) => 
      request<unknown>(`/api/v1/alerts/incidents/${id}/resolve`, { method: 'POST' }),
  },
  
  retention: {
    getPolicies: (tenantId: string) => 
      request<unknown[]>(`/api/v1/retention/policies?tenantId=${tenantId}`),
    createPolicy: (policy: unknown) => 
      request<unknown>('/api/v1/retention/policies', { method: 'POST', body: policy }),
    simulate: (tenantId: string, policyId: string) => 
      request<unknown>(`/api/v1/retention/simulate?tenantId=${tenantId}&policyId=${policyId}`, { method: 'POST' }),
  },
}
