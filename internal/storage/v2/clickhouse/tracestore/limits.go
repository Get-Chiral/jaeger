// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

// maxTraceIDsPerQuery caps how many trace IDs one FindTraceIDs query may
// return. Larger result sets are cut, not paginated.
const maxTraceIDsPerQuery = 1000
